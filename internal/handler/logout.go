package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	security "github.com/rpinedafocus/u-library/internal/middleware"
)

func LogoutController(c *gin.Context) {

	metadata, err := security.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if metadata != nil {
		delErr := security.DestroySession(metadata.AccessUuid)
		if delErr != nil {
			c.JSON(http.StatusUnauthorized, delErr.Error())
			return
		}
		c.JSON(http.StatusOK, "Successfully logged out")
	}
}
