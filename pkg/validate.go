package pkg

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field string
	Tag   string
	Value interface{}
	Message string
}

// Error returns the error message
func (e ValidationError) Error() string {
	return e.Message
}

// ValidationErrors is a slice of ValidationError
type ValidationErrors []ValidationError

// Error returns all error messages concatenated
func (es ValidationErrors) Error() string {
	var errStrings []string
	for _, err := range es {
		errStrings = append(errStrings, err.Error())
	}
	return strings.Join(errStrings, "; ")
}

// Validate validates a struct and returns custom error messages
func Validate(s interface{}) error {
	validate := validator.New()
	
	// Register function to get tag name from json tag
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return fld.Name
		}
		return name
	})

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors := ValidationErrors{}
	
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		value := err.Value()
		
		// Create custom error message based on the validation tag
		message := fmt.Sprintf("Validation failed on field '%s'", field)
		
		switch tag {
		case "required":
			message = fmt.Sprintf("The %s field is required", field)
		case "email":
			message = fmt.Sprintf("The %s must be a valid email address", field)
		case "min":
			message = fmt.Sprintf("The %s must be at least %s characters long", field, err.Param())
		case "max":
			message = fmt.Sprintf("The %s must not be longer than %s characters", field, err.Param())
		case "len":
			message = fmt.Sprintf("The %s must be exactly %s characters long", field, err.Param())
		case "alphanum":
			message = fmt.Sprintf("The %s must only contain alphanumeric characters", field)
		case "numeric":
			message = fmt.Sprintf("The %s must be a numeric value", field)
		case "uuid":
			message = fmt.Sprintf("The %s must be a valid UUID", field)
		case "url":
			message = fmt.Sprintf("The %s must be a valid URL", field)
		case "oneof":
			options := strings.Replace(err.Param(), " ", ", ", -1)
			message = fmt.Sprintf("The %s must be one of: %s", field, options)
		}
		
		validationErrors = append(validationErrors, ValidationError{
			Field:   field,
			Tag:     tag,
			Value:   value,
			Message: message,
		})
	}

	return validationErrors
}