package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		fmt.Println("/ hitted")
		return c.String(http.StatusOK, "hi, im from echo")
	})

}
