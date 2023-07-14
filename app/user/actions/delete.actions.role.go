package actions

import (
	"errors"
	user "go-echo-modular/app/user/models"
	"go-echo-modular/config"
	"go-echo-modular/helpers"

	"gorm.io/gorm"
)

func Delete(id string) helpers.Response {
	var user user.User
	var res helpers.Response

	db := config.GetDBInstance()

	result := db.Where("id = ?", id).Take(&user)
	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.HttpStatus = 400
			res.Status = "nok"
			res.Message = "can't find record"
			return res
		}

		res.HttpStatus = 500
		res.Status = "nok"
		res.Message = "something went wrong"
		return res
	}

	db.Delete(&user)

	res.HttpStatus = 200
	res.Status = "ok"
	res.Message = "ok"

	return res
}
