package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/internal/dal"
	security "github.com/rpinedafocus/u-library/internal/middleware"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/rpinedafocus/u-library/internal/utils"
)

func CreateUserController(c *gin.Context) {

	mySession := security.GetSessionDetail(c)
	if mySession == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), true, utils.ErrNoSessionInformation.Error(), false)})
		return
	}

	user := &model.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	entity, err := dal.CreateUser(mySession.UserId, user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}
