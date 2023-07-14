package handlers

import (
	"go-echo-modular/app/user/actions"
	"go-echo-modular/app/user/requests"
	"log"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid id",
		})
	}

	p := new(requests.UserUpdateForm)
	if err := c.Bind(p); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Empty payloads",
		})
	}

	v := validate.Struct(p)
	if !v.Validate() {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": v.Errors.One(),
		})
	}

	res := actions.Update(id, p)
	return c.JSON(res.HttpStatus, res)
}