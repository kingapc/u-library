package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/internal/app/u-library/dal"
	security "github.com/rpinedafocus/u-library/internal/app/u-library/middleware"
	"github.com/rpinedafocus/u-library/internal/app/u-library/model"
	"github.com/rpinedafocus/u-library/internal/app/u-library/utils"
)

func CreateRoleController(c *gin.Context) {

	mySession := security.GetSessionDetail(c)
	if mySession == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), true, utils.ErrNoSessionInformation.Error(), false)})
		return
	}

	role := &model.Role{}

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	entity, err := dal.CreateRole(mySession.UserId, role)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}
