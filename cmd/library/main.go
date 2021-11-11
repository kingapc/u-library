package main

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/rpinedafocus/u-library/internal/handler"
	security "github.com/rpinedafocus/u-library/internal/middleware"
)

func main() {

	var router = gin.Default()
	security.Init()
	router.POST("/login", handler.LoginController)
	router.POST("/logout", security.ValidateTokenSession(), handler.LogoutController)

	router.POST("/access/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateAccessRolesController)
	router.POST("/roles/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateRoleController)

	router.GET("/books", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchAllBooksController)
	router.POST("/books/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateBookController)
	router.GET("/books/:id", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchBookByIdController)

	router.POST("/booking-rent/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateBookingRentController)
	router.PUT("/booking-rent/release/:id", security.ValidateTokenSession(), security.ValidateAccess(), handler.ReleaseBookingRentController)

	router.POST("/users/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateUserController)
	router.POST("/genres/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateGenreController)
	router.POST("/authors/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateAuthorController)

	router.GET("/common/mybooks", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchBookingRentBooks)     //User inquiry his books
	router.GET("/common/mybooks/:id", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchBookingRentBooks) //admin inquiry the user's books

	log.Fatal(router.Run("localhost:8080"))
}
