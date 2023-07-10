package handlers

import (
	"go-echo-modular/app/role/actions"
	"go-echo-modular/app/role/requests"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {

	p := new(requests.RoleCreateForm)

	res := actions.Create(p)
	return c.JSON(res.Status, res)

}
