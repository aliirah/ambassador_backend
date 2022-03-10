package auth

import (
	"alirah/util/validation"
	"github.com/go-playground/validator"
)

type UpdateUserData struct {
	FirstName string `json:"first_name" validate:"required,max=255"`
	LastName  string `json:"last_name" validate:"required,max=255"`
	Email     string `json:"email" validate:"required,email,max=255"`
}

func UpdateUserValidate(data *UpdateUserData) map[string]string {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return validation.GetErrorBag(err)
	}

	return nil
}
