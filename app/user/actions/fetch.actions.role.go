package actions

import (
	user "go-echo-modular/app/user/models"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Fetch(*echo.Map) helpers.Response {
	var users []user.User
	var res helpers.Response

	db := config.GetDBInstance()

	if result := db.Find(&users); result.Error != nil {
		log.Print("error Fetch")
		log.Print(result.Error)

		res.HttpStatus = http.StatusInternalServerError
		res.Status = "nok"
		res.Message = "error fetchin records"
		return res
	}

	res.HttpStatus = http.StatusOK
	res.Status = "ok"
	res.Message = "ok"
	res.Data = users

	return res
}
