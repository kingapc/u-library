package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	security "github.com/rpinedafocus/u-library/pkg/middleware"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

func CreateBookController(c *gin.Context) {

	mySession := security.GetSessionDetail(c)
	if mySession == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), true, utils.ErrNoSessionInformation.Error(), false)})
		return
	}

	book := &model.Book{}

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, err.Error(), false)})
		return
	}

	entity, err := dal.CreateBook(mySession.UserId, book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusCreated, gin.H{"operation": utils.Success(http.StatusText(http.StatusCreated)), "response": entity})
	}
}

func FetchAllBooksController(c *gin.Context) {

	a := []string{"title", "author", "genre"}
	p, v := utils.ValidateParams(c, a)

	if p == "" || v == "" {
		if !(c.Request.RequestURI == c.Request.URL.Path) {
			c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), false, "", false)})
			return
		}
	}

	entity, err := dal.FetchAll(p, v)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": entity})
	}
}

func FetchBookByIdController(c *gin.Context) {

	req, resul := utils.IsValidUUID(c.Params.ByName("id"))
	if !resul {
		c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.InvalidId.Error(), false)})
		return
	}

	entity, err := dal.FetchBookById(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), true, err.Error(), false)})
	} else {
		c.JSON(http.StatusOK, gin.H{"operation": utils.Success(http.StatusText(http.StatusOK)), "response": entity})
	}
}

func IsBookAvailableController(id string) int {

	entity, err := dal.FetchBookById(id)
	if err != nil {
		return 0
	} else {
		return entity.TotalAvailable
	}
}
