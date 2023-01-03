package common

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(data interface{}) []*ErrorValidationResponse {
	var errors []*ErrorValidationResponse

	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidationResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)
		}
	}

	return errors
}
