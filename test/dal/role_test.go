package dal

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/stretchr/testify/assert"
)

var r = &model.RoleEntity{
	Entity: model.Entity{
		ID:        uuid.New().String(),
		CreatedBy: "user",
		CreatedAt: time.Now().Format("01-02-2006"),
	},
	Role: model.Role{
		Name: "Test",
	},
}

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

func TestCreateRole(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}

	defer func() {
		repo.db.Close()
	}()

	query := "INSERT INTO university.roles \\(id, name, created_by, created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(r.ID, r.Name, r.CreatedBy, r.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	myQuery := "INSERT INTO university.roles (id, name, created_by, created_at) VALUES ($1, $2, $3, $4)"
	stmt, err := repo.db.PrepareContext(ctx, myQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, r.ID, r.Name, r.CreatedBy, r.CreatedAt)

	assert.NoError(t, err)
}
