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

	result := db.Where("is_active = ?", true).First(&role, id)
	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusOK
			res.Message = "can't find record"
			return res
		}

		res.Status = http.StatusInternalServerError
		res.Message = "Something went wrong!"
		return res
	}
	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = role

	return res
}
