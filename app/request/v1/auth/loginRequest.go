package auth

import (
	domain "alirah/app/domain/user"
	"alirah/database"
	"alirah/util/validation"
	"github.com/go-playground/validator"
)

type LoginData struct {
	Password string `json:"password" validate:"required,min=8,max=32"`
	Email    string `json:"email" validate:"required,email,max=255"`
}

func LoginValidate(data *LoginData) map[string]string {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return validation.GetErrorBag(err)
	}

	var user domain.User
	res := database.DB.
		Where("email = ?", data.Email).
		Find(&user)

	if res.RowsAffected == 0 {
		return map[string]string{
			"credential": "email or password is wrong!",
		}
	}

	err := user.ComparePassword(data.Password)
	if err != nil {
		return map[string]string{
			"credential": "email or password is wrong!",
		}
	}

	return nil
}
