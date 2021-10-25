package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateBookingRentController(t *testing.T) {
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
			CreateBookingRentController(tt.args.c)
		})
	}
}

func TestReleaseBookingRentController(t *testing.T) {
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
			ReleaseBookingRentController(tt.args.c)
		})
	}
}

func TestFetchBookingRentBooks(t *testing.T) {
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
			FetchBookingRentBooks(tt.args.c)
		})
	}
}
