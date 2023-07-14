package handlers

import (
	"go-echo-modular/app/role/actions"
	"go-echo-modular/app/role/requests"
	"go-echo-modular/helpers"
	"log"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	var res helpers.Response

	p := new(requests.RoleCreateForm)
	if err := c.Bind(p); err != nil {
		log.Println(err)

		res.HttpStatus = 500
		res.Status = "nok"
		res.Message = "Empty payloads"

		return c.JSON(res.HttpStatus, res)
	}

	v := validate.Struct(p)
	if !v.Validate() {

		res.HttpStatus = 400
		res.Status = "nok"
		res.Message = v.Errors.One()

		return c.JSON(res.HttpStatus, res)
	}

	result := actions.Create(p)
	return c.JSON(result.HttpStatus, result)

}
