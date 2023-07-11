package role

import (
	"go-echo-modular/app/role/handlers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	router := e.Group("roles")

	router.GET("", handlers.Fetch)
	router.GET("/:id", handlers.Detail)
	router.PUT("/:id", handlers.Update)
	router.DELETE("/:id", handlers.Delete)
	router.POST("", handlers.Create)
}
