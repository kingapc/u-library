package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	handler "github.com/rpinedafocus/u-library/internal/handler"
	"github.com/rpinedafocus/u-library/internal/model"
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
	var option int = 1 //change the switch option to see a fail test

	c.Set("myData", myData)

	b := &model.Book{}
	b.Title = "Cien años de soledad III"
	b.AuthorId = 8 //Set a fake AuthorId if you want to see an error
	b.GenreId = 1  //Set a fake GenreId if you want to see an error
	b.PublishDate = "01-01-1967"
	b.TotalAvailable = 1

	body, _ := json.Marshal(b)
	c.Request = httptest.NewRequest("POST", "http://localhost:8080/books/create", bytes.NewReader(body))

	handler.CreateBookController(c)

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

	var option int = 1 //change the switch option to see a fail test

	c.Request = httptest.NewRequest("GET", "http://localhost:8080/books", nil)
	c.Request.RequestURI = "http://localhost:8080/books"
	c.Request.URL.Path = "http://localhost:8080/books"

	handler.FetchAllBooksController(c)

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
	var option int = 1 //change the switch option to see a fail test

	for key, value := range params {
		paramsSlice = append(paramsSlice, gin.Param{
			Key:   key,
			Value: value,
		})
	}
	c.Params = paramsSlice

	handler.FetchBookByIdController(c)

	switch option {
	case 1:
		t.Run("Success", func(t *testing.T) { assert.Equal(t, 200, w.Code) })
	case 2:
		t.Run("Fail", func(t *testing.T) { assert.NotEqual(t, 200, w.Code) })
	}
}
