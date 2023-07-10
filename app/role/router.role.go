package role

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	router := e.Group("roles")

	router.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "roles ep")
	})
}
