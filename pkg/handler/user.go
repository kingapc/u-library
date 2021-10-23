package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

// create handles the user create request
func CreateUserController(c *gin.Context) {

	user := &model.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	temppw, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusMethodNotAllowed), true, err.Error(), false)})
		return
	}
	user.Password = temppw

	entity, err := dal.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}
