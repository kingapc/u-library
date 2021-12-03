package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/stretchr/testify/assert"
)

type repository struct {
	db *sql.DB
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

//Struct in BookController
type MockBookStruct struct {
	SessionDetailFn func(*gin.Context) *model.AccessDetails
}

//Session detail
var e = &model.AccessDetails{
	AccessUuid: "0b766bed69764a6a81c393c702aad3ed",
	MyId:       "27f7f309d76e4ff3adf286cdf6712c1a",
	UserId:     "admin",
	RoleId:     "2",
}

//Session detail context
var myData = map[string]interface{}{
	"uuid":   e.AccessUuid,
	"myId":   e.MyId,
	"userId": e.UserId,
	"roleId": e.RoleId,
}

var b = &model.FetchBook{
	ID:             uuid.New().String(),
	Title:          "Test Book",
	Author:         "Author Test",
	Genre:          "Genre Test",
	PublishDate:    "01-01-1990",
	TotalAvailable: 12,
}

func TestCreateBookControllerSession(t *testing.T) {

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)

	c.Set("myData", myData)

	b := &model.Book{}
	b.Title = "Cien años de soledad III"
	b.AuthorId = 8
	b.GenreId = 1
	b.PublishDate = "01-01-1967"
	b.TotalAvailable = 1

	body, _ := json.Marshal(b)
	c.Request = httptest.NewRequest("POST", "/books/create", bytes.NewReader(body))

	myMock := MockBookStruct{}
	myMock.SessionDetailFn = func(c *gin.Context) *model.AccessDetails {
		data, _ := c.Get("myData")
		mapData, _ := data.(map[string]interface{})

		e := &model.AccessDetails{}
		e.AccessUuid = mapData["uuid"].(string)
		e.MyId = mapData["myId"].(string)
		e.UserId = mapData["userId"].(string)
		e.RoleId = mapData["roleId"].(string)

		return e
	}

	assert.Equal(t, myMock.SessionDetailFn(c), e)
}

func TestCreateBookControllerSessionFail(t *testing.T) {

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)

	c.Set("myData", myData)

	b := &model.Book{}
	b.Title = "Cien años de soledad III"
	b.AuthorId = 8
	b.GenreId = 1
	b.PublishDate = "01-01-1967"
	b.TotalAvailable = 1

	body, _ := json.Marshal(b)
	c.Request = httptest.NewRequest("POST", "/books/create", bytes.NewReader(body))

	myMock := MockBookStruct{}
	myMock.SessionDetailFn = func(c *gin.Context) *model.AccessDetails {
		e := &model.AccessDetails{}
		return e
	}

	assert.NotEqual(t, myMock.SessionDetailFn(c), e)
}

func TestFetchAllBooksController(t *testing.T) {

	db, mock := NewMock()
	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	query := `SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available 
				FROM university.books b
				INNER JOIN university.authors a on a.author_id = b.author_id
				INNER JOIN university.genres g on g.genre_id = b.genre_id
				WHERE b.active = true AND b.is_deleted = 'N'`

	rows := sqlmock.NewRows([]string{"id", "title", "author", "genre", "publish_date", "total_avilable"}).
		AddRow(b.ID, b.Title, b.Author, b.Genre, b.PublishDate, b.TotalAvailable)

	mock.ExpectQuery(query).WillReturnRows(rows)
	books := make([]*model.FetchBook, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	myQuery := `SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available 
				FROM university.books b
				INNER JOIN university.authors a on a.author_id = b.author_id
				INNER JOIN university.genres g on g.genre_id = b.genre_id
				WHERE b.active = true AND b.is_deleted = 'N'`

	rowsR, err := repo.db.QueryContext(ctx, myQuery)

	for rowsR.Next() {
		book := new(model.FetchBook)
		err = rowsR.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Genre,
			&book.PublishDate,
			&book.TotalAvailable,
		)

		books = append(books, book)
	}

	assert.NotEmpty(t, books)
	assert.NoError(t, err)
	assert.Len(t, books, 1)
}

func TestFetchAllBooksControllerError(t *testing.T) {

	db, mock := NewMock()
	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	query := `SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available 
				FROM university.books b
				INNER JOIN university.authors a on a.author_id = b.author_id
				INNER JOIN university.genres g on g.genre_id = b.genre_id
				WHERE b.active = true AND b.is_deleted = 'N'`

	rows := sqlmock.NewRows([]string{"id", "title", "author", "genre", "publish_date", "total_avilable"}).
		AddRow(b.ID, b.Title, b.Author, b.Genre, b.PublishDate, b.TotalAvailable)

	mock.ExpectQuery(query).WillReturnRows(rows)
	books := make([]*model.FetchBook, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	myQuery := `SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available 
				FROM university.books b
				INNER JOIN university.authors a on a.author_id = b.author_id
				INNER JOIN university.genres g on g.genre_id = b.genre_id
				WHERE b.active = true AND b.is_deleted = 'S'`

	rowsR, err := repo.db.QueryContext(ctx, myQuery)

	if rowsR == nil {
		assert.Empty(t, books)
		assert.Error(t, err)
		assert.Len(t, books, 0)
	}
}

func TestFetchBookByIdController(t *testing.T) {

	db, mock := NewMock()
	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	query := "SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available " +
		"FROM university.books b " +
		"INNER JOIN university.authors a on a.author_id = b.author_id " +
		"INNER JOIN university.genres g on g.genre_id = b.genre_id " +
		"WHERE b.active = true AND b.is_deleted = 'N' AND b.id = \\$1"

	rows := sqlmock.NewRows([]string{"id", "title", "author", "genre", "publish_date", "total_avilable"}).
		AddRow(b.ID, b.Title, b.Author, b.Genre, b.PublishDate, b.TotalAvailable)

	mock.ExpectQuery(query).WithArgs(b.ID).WillReturnRows(rows)

	book := new(model.FetchBook)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	myQuery := "SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available " +
		"FROM university.books b " +
		"INNER JOIN university.authors a on a.author_id = b.author_id " +
		"INNER JOIN university.genres g on g.genre_id = b.genre_id " +
		"WHERE b.active = true AND b.is_deleted = 'N' AND b.id = $1 "

	err := repo.db.QueryRowContext(ctx, myQuery, b.ID).Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.PublishDate, &book.TotalAvailable)

	assert.NotNil(t, book)
	assert.NoError(t, err)
}
