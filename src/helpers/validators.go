package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ErrorMessages is a map to hold error messages
type ErrorMessages map[string]string

// ValidateStruct validates the struct fields based on the provided tags
func ValidateStruct(data interface{}, customErrMsg ErrorMessages) ErrorMessages {
	valid := validator.New()
	errorMessages := make(ErrorMessages)

	if err := valid.Struct(data); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			key := fmt.Sprintf("%s.%s", e.StructNamespace(), e.Tag())

			if msg, exists := customErrMsg[key]; exists {
				errorMessages[key] = msg
			} else {
				errorMessages[key] = "campo invalido"
			}
		}
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	return nil
}