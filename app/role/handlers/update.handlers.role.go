package handlers

import (
	"go-echo-modular/app/role/actions"
	"go-echo-modular/app/role/requests"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"message": "invalid id",
		})
	}
	p := new(requests.RoleUpdateForm)

	res := actions.Update(id, p)

	return c.JSON(res.Status, res)
}
