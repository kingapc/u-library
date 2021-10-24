package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

// create handles the user create request
func CreateAccessRolesController(c *gin.Context) {

	accrol := &model.AccessRole{}

	if err := c.BindJSON(&accrol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	entity, err := dal.CreateAccessRole(accrol)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}

// func FetchExistAccessRole(roleId int, endpoint string) (bool, error) {

// 	result, err := dal.FetchAccessRolById(roleId, endpoint)
// 	if err != nil {
// 		return false, err
// 	} else {
// 		return result, nil
// 	}
// }
