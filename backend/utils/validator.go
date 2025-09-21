package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Single field validation error
type FieldError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

// ValidationError wraps multiple field errors
type ValidationError struct {
	Errors []FieldError `json:"validation_errors"`
}

func (v ValidationError) Error() string {
	return "validation failed"
}

// FormatValidationError converts validator errors into structured JSON
func FormatValidationError(err error, model interface{}) error {
	if verrs, ok := err.(validator.ValidationErrors); ok {
		var fieldErrors []FieldError

		typ := reflect.TypeOf(model)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}

		for _, v := range verrs {
			fieldName := v.Field() // fallback if no json tag

			if field, ok := typ.FieldByName(v.StructField()); ok {
				if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
					fieldName = jsonTag
				}
			}

			fieldErrors = append(fieldErrors, FieldError{
				Field: fieldName,
				Rule:  v.Tag(),
			})
		}

		return ValidationError{Errors: fieldErrors}
	}
	return err
}
