package handler

import "github.com/go-playground/validator/v10"

type ErrorValidate struct {
	Message string
}

func ValidateStruct(data interface{}) []*ErrorValidate {
	var validate = validator.New()
	var errors []*ErrorValidate
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidate
			element.Message = err.Field() + " is " + err.Tag() + err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
