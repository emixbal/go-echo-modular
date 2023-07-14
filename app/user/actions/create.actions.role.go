package actions

import (
	user "go-echo-modular/app/user/models"
	"go-echo-modular/app/user/requests"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"log"
)

func Create(p *requests.UserCreateForm) helpers.Response {
	var user user.User
	var res helpers.Response

	db := config.GetDBInstance()

	user.Name = p.Name

	if result := db.Create(&user); result.Error != nil {
		log.Println("error Create ")
		log.Println(result.Error)

		res.HttpStatus = 500
		res.Status = "nok"
		res.Message = "error save new record"
		return res
	}

	res.HttpStatus = 200
	res.Status = "ok"
	res.Message = "success"
	res.Data = user

	return res
}
