package requests

import "github.com/gookit/validate"

type UserCreateForm struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required"`
}

// Messages you can custom validator error messages.
func (f UserCreateForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f UserCreateForm) Translates() map[string]string {
	return validate.MS{
		"Name": "name",
	}
}
