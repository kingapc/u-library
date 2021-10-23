package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

func CreateRentController(c *gin.Context) {

	rent := &model.BookingRent{}

	if err := c.BindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	if _, resul := utils.IsValidUUID(rent.BookId); !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	if _, resul := utils.IsValidUUID(rent.BookingRentBy); !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	if totalAvailable := IsBookAvailableController(rent.BookId); totalAvailable == 0 {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningNoBooksAvailables.Error(), false)})
		return
	}

	exist, err := ExistValidRent(rent.BookId, rent.BookingRentBy)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), true)})
		return
	} else if exist {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningRentExist.Error(), true)})
		return
	}

	entity, err := dal.CreateRent(rent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}

func ReturnRentedBookController(c *gin.Context) {

	req, resul := utils.IsValidUUID(c.Params.ByName("id"))
	if !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	result, err := dal.IsValidRent(req)
	if err != nil || !result {
		c.JSON(http.StatusOK, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusOK), true, utils.WarningNoRentAvailable.Error(), true)})
		return
	}

	entity, err := dal.ReturnRentedBook(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": entity})
	}
}
