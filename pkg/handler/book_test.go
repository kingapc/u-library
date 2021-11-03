package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookController(t *testing.T) {

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)

	e := &model.AccessDetails{}
	e.AccessUuid = "0b766bed69764a6a81c393c702aad3ed"
	e.MyId = "27f7f309d76e4ff3adf286cdf6712c1a"
	e.UserId = "admin"
	e.RoleId = "2"

	myData := map[string]interface{}{
		"uuid":   e.AccessUuid,
		"myId":   e.MyId,
		"userId": e.UserId,
		"roleId": e.RoleId,
	}
	var option int = 2

	c.Set("myData", myData)

	b := &model.Book{}
	b.Title = "Cien años de soledad III"
	b.AuthorId = 8
	b.GenreId = 15
	b.PublishDate = "01-01-1967"
	b.TotalAvailable = 1

	body, _ := json.Marshal(b)
	c.Request = httptest.NewRequest("POST", "http://localhost:8080/books/create", bytes.NewReader(body))

	CreateBookController(c)

	switch option {
	case 1:
		t.Run("Success", func(t *testing.T) { assert.Equal(t, 201, w.Code) })
	case 2:
		t.Run("Fail", func(t *testing.T) { assert.NotEqual(t, 201, w.Code) })
	}
}

func TestFetchAllBooksController(t *testing.T) {

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)

	var option int = 1

	c.Request = httptest.NewRequest("GET", "http://localhost:8080/books", nil)
	c.Request.RequestURI = "http://localhost:8080/books"
	c.Request.URL.Path = "http://localhost:8080/books"

	FetchAllBooksController(c)

	switch option {
	case 1:
		t.Run("Success", func(t *testing.T) { assert.Equal(t, 200, w.Code) })
	case 2:
		t.Run("Fail", func(t *testing.T) { assert.NotEqual(t, 200, w.Code) })
	}
}

func TestFetchBookByIdController(t *testing.T) {

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)

	var params = map[string]string{"id": "0b766bed69764a6a81c393c702aad3ed"}
	var paramsSlice []gin.Param
	var option int = 1

	for key, value := range params {
		paramsSlice = append(paramsSlice, gin.Param{
			Key:   key,
			Value: value,
		})
	}
	c.Params = paramsSlice

	FetchBookByIdController(c)

	switch option {
	case 1:
		t.Run("Success", func(t *testing.T) { assert.Equal(t, 200, w.Code) })
	case 2:
		t.Run("Fail", func(t *testing.T) { assert.NotEqual(t, 200, w.Code) })
	}
}
