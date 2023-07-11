package handlers

import (
	"go-echo-modular/app/user/actions"

	"github.com/labstack/echo/v4"
)

func Fetch(c echo.Context) error {

	res := actions.Fetch(&echo.Map{
		"per_page": 10,
		"page":     1,
	})

	return c.JSON(res.HttpStatus, res)

}
