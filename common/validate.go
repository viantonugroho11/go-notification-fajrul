package common

import (

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate = validator.New()

func ValidateStruct(user interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Message = err.Error()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateType(Type string)[]*ErrorResponse{
	var errors []*ErrorResponse
	if Type != "firebase" && Type != "email" {
		var element ErrorResponse
		element.Field = "type"
		element.Message = "type must be firebase or email"
		errors = append(errors, &element)
	}
	return errors
}
