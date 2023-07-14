package handlers

import (
	"go-echo-modular/app/user/actions"
	"go-echo-modular/app/user/requests"
	"log"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	p := new(requests.UserCreateForm)
	if err := c.Bind(p); err != nil {
		log.Println(err)
		return c.JSON(500, echo.Map{
			"message": "Empty payloads",
		})
	}

	v := validate.Struct(p)
	if !v.Validate() {
		return c.JSON(400, echo.Map{
			"message": v.Errors.One(),
		})
	}

	log.Println(p.Name)

	res := actions.Create(p)
	return c.JSON(res.HttpStatus, res)

}
