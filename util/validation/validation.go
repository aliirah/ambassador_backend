package validation

import (
	stringHelper "alirah/util/string"
	"fmt"
	"github.com/go-playground/validator"
)

func GetErrorBag(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {

		///////// CUSTOM MESSAGE /////////
		//if err.Tag() == "email" {
		//	errors[strings.ToLower(err.Field())] = "Invalid E-mail format."
		//	continue
		//}

		field := stringHelper.ToSnakeCase(err.Field())
		param := stringHelper.ToSnakeCase(err.Param())
		errors[field] = fmt.Sprintf("%s should be %s %s", field, err.Tag(), param)
	}
	return errors
}
