package actions

import (
	role "go-echo-modular/app/role/models"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Fetch(*echo.Map) helpers.Response {
	var roles []role.Role
	var res helpers.Response

	db := config.GetDBInstance()

	if result := db.Find(&roles); result.Error != nil {
		log.Print("error Fetch")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = roles

	return res
}
