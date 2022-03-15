package auth

import (
	domain "alirah/app/domain/user"
	"alirah/database"
	"alirah/util/validation"
	"github.com/go-playground/validator"
)

type RegisterData struct {
	Password        string `json:"password" validate:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	FirstName       string `json:"first_name" validate:"required,max=255"`
	LastName        string `json:"last_name" validate:"required,max=255"`
	Email           string `json:"email" validate:"required,email,max=255"`
}

func RegisterValidate(data *RegisterData) map[string]string {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return validation.GetErrorBag(err)
	}

	var user domain.User
	res := database.DB.Find(&user, "email = ?", data.Email)
	if res.RowsAffected > 0 {
		return map[string]string{
			"email": "email already exists.",
		}
	}

	return nil
}
