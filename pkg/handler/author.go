package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
)

// create handles the user create request
func CreateAuthorController(c *gin.Context) (*model.AuthorEntity, string) {

	author := &model.Author{}

	if err := c.BindJSON(&author); err != nil {
		return nil, err.Error() // utils.ErrorX(400)
	}

	entity, err := dal.CreateAuthor(author)
	if err != nil {
		return nil, err.Error() //utils.ErrorX(400)
	}

	return entity, ""
}
