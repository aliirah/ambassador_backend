package product

import (
	"alirah/util/validation"
	"github.com/go-playground/validator"
)

type UpdateData struct {
	Title       string  `json:"title" validate:"required,min=3"`
	Description string  `json:"description" validate:"required,min=8,max=512"`
	Image       string  `json:"image" validate:"required"`
	Price       float64 `json:"price" validate:"required,numeric"`
}

func UpdateValidate(data *UpdateData) map[string]string {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return validation.GetErrorBag(err)
	}

	return nil
}
