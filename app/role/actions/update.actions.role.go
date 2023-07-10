package actions

import (
	"errors"
	role "go-echo-modular/app/role/models"
	"go-echo-modular/app/role/requests"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"net/http"

	"gorm.io/gorm"
)

func Update(id string, p *requests.RoleUpdateForm) helpers.Response {
	var role role.Role
	var res helpers.Response

	db := config.GetDBInstance()

	result := db.Where("id = ?", id).Take(&role)
	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusOK
			res.Message = "can't find record"
			return res
		}

		res.Status = http.StatusInternalServerError
		res.Message = "something went wrong"
		return res
	}

	role.Name = p.Name

	db.Save(&role)

	return res
}
