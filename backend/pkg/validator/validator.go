package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	// Convert validation errors to map
	errorMap := make(map[string]string)
	for _, fieldErr := range validationErrors {
		fieldName := toSnakeCase(fieldErr.Field())
		errorMap[fieldName] = fmt.Sprintf("%s is required", fieldName)
	}
	return &ValidationError{Errors: errorMap}
}

// ValidationError struct to wrap the validation errors
type ValidationError struct {
	Errors map[string]string `json:"errors"`
}

func (v *ValidationError) Error() string {
	return "validation failed"
}

// optional: convert CamelCase field name to snake_case
func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
