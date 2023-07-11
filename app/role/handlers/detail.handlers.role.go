package handlers

import (
	"go-echo-modular/app/role/actions"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Detail(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"message": "invalid user id",
		})
	}

	res := actions.Detail(id)

	return c.JSON(res.HttpStatus, res)

}
