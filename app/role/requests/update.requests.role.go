package requests

import "github.com/gookit/validate"

type RoleUpdateForm struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required"`
}

// Messages you can custom validator error messages.
func (f RoleUpdateForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f RoleUpdateForm) Translates() map[string]string {
	return validate.MS{
		"Name": "name",
	}
}
