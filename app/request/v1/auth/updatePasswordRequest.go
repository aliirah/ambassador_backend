package auth

import (
	"alirah/util/validation"
	"github.com/go-playground/validator"
)

type UpdatePasswordData struct {
	Password        string `json:"password" validate:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}

func UpdatePasswordValidate(data *UpdatePasswordData) map[string]string {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return validation.GetErrorBag(err)
	}

	return nil
}
