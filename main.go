package main

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/rpinedafocus/u-library/pkg/handler"
	security "github.com/rpinedafocus/u-library/pkg/middleware"
)

func main() {

	var router = gin.Default()
	security.Init()
	router.POST("/login", handler.LoginController)
	router.POST("/logout", security.ValidateTokenSession(), handler.LogoutController)

	router.POST("/access/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateAccessRolesController)

	router.GET("/books", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchAllBooksController)
	router.POST("/books/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateBookController)
	router.GET("/books/:id", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchBookByIdController)

	router.POST("/bookings/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateBookingController)
	router.POST("/bookings/release/:id", security.ValidateTokenSession(), security.ValidateAccess(), handler.ReleaseBookingController)

	router.POST("/users/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateUserController)
	router.POST("/roles/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateRoleController)
	router.POST("/genres/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateGenreController)
	router.POST("/authors/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateAuthorController)

	router.POST("/rent/create", security.ValidateTokenSession(), security.ValidateAccess(), handler.CreateRentController)
	router.POST("/rent/return/:id", security.ValidateTokenSession(), security.ValidateAccess(), handler.ReturnRentedBookController)

	router.GET("/common/mybooks", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchBookingRentBooks)     //User inquiry his books
	router.GET("/common/mybooks/:id", security.ValidateTokenSession(), security.ValidateAccess(), handler.FetchBookingRentBooks) //admin inquiry the user's books

	log.Fatal(router.Run("localhost:8080"))
}
