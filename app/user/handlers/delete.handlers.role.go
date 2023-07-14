package handlers

import (
	"go-echo-modular/app/user/actions"
	"go-echo-modular/helpers"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	var res helpers.Response

	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)

		res.HttpStatus = 400
		res.Status = "nok"
		res.Message = "invalid id"

		return c.JSON(res.HttpStatus, res)
	}

	result := actions.Delete(id)
	return c.JSON(result.HttpStatus, result)
}
