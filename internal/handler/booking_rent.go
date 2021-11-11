package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/internal/dal"
	security "github.com/rpinedafocus/u-library/internal/middleware"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/rpinedafocus/u-library/internal/utils"
)

func CreateBookingRentController(c *gin.Context) {

	mySession := security.GetSessionDetail(c)
	if mySession == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), true, utils.ErrNoSessionInformation.Error(), false)})
		return
	}

	br := &model.BookingRent{}

	if err := c.BindJSON(&br); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	if br.IsBooking && br.ProcessDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.ErrRequiredDateProcess.Error(), false)})
		return
	}

	if _, resul := utils.IsValidUUID(br.BookId); !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	if totalAvailable := IsBookAvailableController(br.BookId); totalAvailable == 0 {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningNoBooksAvailables.Error(), true)})
		return
	}

	br.ProcessBy = mySession.MyId
	exist, err := dal.ExistValidBookingRent(br.BookId, mySession.MyId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), true)})
		return
	} else if exist {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningBookingRentExist.Error(), true)})
		return
	}

	entity, err := dal.CreateBookingRent(mySession.UserId, br)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), true)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}

func ReleaseBookingRentController(c *gin.Context) {

	mySession := security.GetSessionDetail(c)
	if mySession == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), true, utils.ErrNoSessionInformation.Error(), false)})
		return
	}

	req, resul := utils.IsValidUUID(c.Params.ByName("id"))
	if !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	result, err := dal.IsValidBookingRent(req)
	if err != nil || result {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningNoBookingRentAvailable.Error(), true)})
		return
	}

	entity, err := dal.ReleaseBookingRent(mySession.UserId, req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": entity})
	}
}

func FetchBookingRentBooks(c *gin.Context) {

	mySession := security.GetSessionDetail(c)
	if mySession == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), true, utils.ErrNoSessionInformation.Error(), false)})
		return
	}

	var req string
	var resul bool
	if c.Request.URL.Path == "/common/mybooks" {
		req = mySession.MyId
	} else {
		req, resul = utils.IsValidUUID(c.Params.ByName("id"))
		if !resul {
			c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
			return
		}
	}

	entity, err := dal.FetchBookingRent(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": entity})
	}
}
