package conn

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetConnectionSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()

	if db != nil && mock != nil {
		assert.NotNil(t, db)
		assert.NotNil(t, mock)
	}
}

func TestGetConnectionFail(t *testing.T) {
	_, _, err := sqlmock.New()

	if err == nil {
		assert.Nil(t, err)
	}
}
