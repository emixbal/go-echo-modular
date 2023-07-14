package handlers

import (
	"go-echo-modular/app/user/actions"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return c.JSON(400, echo.Map{
			"message": "invalid id",
		})
	}

	res := actions.Delete(id)
	return c.JSON(res.HttpStatus, res)
}
