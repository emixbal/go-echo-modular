package actions

import (
	"errors"
	role "go-echo-modular/app/role/models"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"net/http"

	"gorm.io/gorm"
)

func Detail(id string) helpers.Response {
	var role role.Role
	var res helpers.Response

	db := config.GetDBInstance()

	result := db.First(&role, id)
	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.HttpStatus = http.StatusOK
			res.Status = "nok"
			res.Message = "can't find record"
			return res
		}

		res.HttpStatus = http.StatusInternalServerError
		res.Status = "nok"
		res.Message = "Something went wrong!"
		return res
	}
	res.HttpStatus = http.StatusOK
	res.Status = "ok"
	res.Message = "ok"
	res.Data = role

	return res
}
