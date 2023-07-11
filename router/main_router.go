package router

import (
	"go-echo-modular/app/role"
	"go-echo-modular/app/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "/home")
	})

	role.Router(e)
	user.Router(e)
}
