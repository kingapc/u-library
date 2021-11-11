package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	security "github.com/rpinedafocus/u-library/internal/middleware"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/rpinedafocus/u-library/internal/utils"
)

func LoginController(c *gin.Context) {

	credentials := &model.Credentials{}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	e, err := security.PrepareLogin(credentials.User, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
		return
	}

	tok, err := security.CreateToken(e.Entity.ID, e.User.User, e.User.Role)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusNonAuthoritativeInfo), true, err.Error(), false)})
		return
	}

	saveErr := security.CreateSession(e.Entity.ID, tok)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	tokens := map[string]string{
		"access_token": tok.AccessToken,
	}

	c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": tokens})
}
