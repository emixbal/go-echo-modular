package handlers

import (
	"go-echo-modular/app/role/actions"
	"go-echo-modular/app/role/requests"
	"go-echo-modular/helpers"
	"log"
	"strconv"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
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

	p := new(requests.RoleUpdateForm)
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

	result := actions.Update(id, p)
	return c.JSON(result.HttpStatus, result)
}
