package actions

import (
	user "go-echo-modular/app/user/models"
	"go-echo-modular/app/user/requests"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"log"
	"net/http"
)

func Create(p *requests.UserCreateForm) helpers.Response {
	var user user.User
	var res helpers.Response

	db := config.GetDBInstance()

	user.Name = p.Name

	if result := db.Create(&user); result.Error != nil {
		log.Println("error Create ")
		log.Println(result.Error)

		res.HttpStatus = http.StatusInternalServerError
		res.Status = "nok"
		res.Message = "error save new record"
		return res
	}

	res.HttpStatus = http.StatusOK
	res.Status = "ok"
	res.Message = "success"
	res.Data = user

	return res
}
