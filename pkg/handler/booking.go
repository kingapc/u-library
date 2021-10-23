package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

// create handles the user create request
func CreateBookingController(c *gin.Context) {

	booking := &model.BookingRent{}

	if err := c.BindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	if _, resul := utils.IsValidUUID(booking.BookId); !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	if _, resul := utils.IsValidUUID(booking.BookingRentBy); !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	if totalAvailable := IsBookAvailableController(booking.BookId); totalAvailable == 0 {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningNoBooksAvailables.Error(), true)})
		return
	}

	exist, err := ExistValidBooking(booking.BookId, booking.BookingRentBy)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), true)})
		return
	} else if exist {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningBookingExist.Error(), true)})
		return
	}

	entity, err := dal.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}

func ReleaseBookingController(c *gin.Context) {

	req, resul := utils.IsValidUUID(c.Params.ByName("id"))
	if !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	result, err := dal.IsValidBookin(req)
	if err != nil || !result {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningNoBookingAvailable.Error(), false)})
		return
	}

	entity, err := dal.ReleaseBookedBook(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": entity})
	}
}
