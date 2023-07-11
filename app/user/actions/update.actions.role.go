package actions

import (
	"errors"
	user "go-echo-modular/app/user/models"
	"go-echo-modular/app/user/requests"
	"go-echo-modular/config"
	"go-echo-modular/helpers"
	"net/http"

	"gorm.io/gorm"
)

func Update(id string, p *requests.UserUpdateForm) helpers.Response {
	var user user.User
	var res helpers.Response

	db := config.GetDBInstance()

	result := db.Where("id = ?", id).Take(&user)
	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.HttpStatus = http.StatusOK
			res.Status = "nok"
			res.Message = "can't find record"
			return res
		}

		res.HttpStatus = http.StatusInternalServerError
		res.Message = "something went wrong"
		return res
	}

	user.Name = p.Name
	db.Save(&user)

	res.HttpStatus = http.StatusOK
	res.Message = "ok"
	res.Data = user

	return res
}
