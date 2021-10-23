package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

func FetchBookingRentBooks(c *gin.Context) {

	req, resul := utils.IsValidUUID(c.Params.ByName("id"))

	if !resul {
		//validar el token
		req = "fd4fd48bb0fa405d8049d4041b58cee4"
	}

	entity, err := dal.FetchBookingRent(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": entity})
	}
}

func ExistValidRent(bookId string, rentBy string) (bool, error) {

	result, err := dal.ExistValidRent(bookId, rentBy)
	if err != nil {
		return false, err
	} else {
		return result, nil
	}
}

func ExistValidBooking(bookId string, rentBy string) (bool, error) {

	result, err := dal.ExistValidBooking(bookId, rentBy)
	if err != nil {
		return false, err
	} else {
		return result, nil
	}
}
