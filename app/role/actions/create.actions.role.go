package actions

import (
	role "go-echo-modular/app/role/models"
	"go-echo-modular/app/role/requests"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"log"
)

func Create(p *requests.RoleCreateForm) helpers.Response {
	var role role.Role
	var res helpers.Response

	db := config.GetDBInstance()

	role.Name = p.Name

	if result := db.Create(&role); result.Error != nil {
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
	res.Data = role

	return res
}
