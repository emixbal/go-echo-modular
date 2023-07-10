package actions

import (
	role "go-echo-modular/app/role/models"
	"go-echo-modular/app/role/requests"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"log"
	"net/http"
)

func Create(p *requests.RoleCreateForm) helpers.Response {
	var role role.Role
	var res helpers.Response

	db := config.GetDBInstance()

	if result := db.Create(&role); result.Error != nil {
		log.Println("error Create ")
		log.Println(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = role

	return res
}
