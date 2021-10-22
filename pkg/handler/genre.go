package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
)

// create handles the user create request
func CreateGenreController(c *gin.Context) (*model.GenreEntity, string) {

	genre := &model.Genre{}

	if err := c.BindJSON(&genre); err != nil {
		return nil, err.Error() // utils.ErrorX(400)
	}

	entity, err := dal.CreateGenre(genre)
	if err != nil {
		return nil, err.Error() //utils.ErrorX(400)
	}

	return entity, ""
}
