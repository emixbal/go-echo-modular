package user

import (
	"go-echo-modular/app/user/handlers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	router := e.Group("users")

	router.GET("", handlers.Fetch)
	router.GET("/:id", handlers.Detail)
	router.PUT("/:id", handlers.Update)
	router.DELETE("/:id", handlers.Delete)
	router.POST("", handlers.Create)
}
