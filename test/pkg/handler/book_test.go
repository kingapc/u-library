package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateBookController(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBookController(tt.args.c)
		})
	}
}

func TestFetchAllBooksController(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FetchAllBooksController(tt.args.c)
		})
	}
}

func TestFetchBookByIdController(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FetchBookByIdController(tt.args.c)
		})
	}
}

func TestIsBookAvailableController(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBookAvailableController(tt.args.id); got != tt.want {
				t.Errorf("IsBookAvailableController() = %v, want %v", got, tt.want)
			}
		})
	}
}
