package validator

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()

	err := validate.Struct(obj)

	// Handle successful validation
	if err == nil {
		return nil
	}

	// Type assert safely
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err // Unexpected error type
	}

	if len(validationErrors) == 0 {
		return errors.New("validation failed")
	}

	validationError := validationErrors[0]
	field := strings.ToLower(validationError.StructField())

	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " must be at most " + validationError.Param())
	case "min":
		return errors.New(field + " must be at least " + validationError.Param())
	case "email":
		return errors.New(field + " is invalid")
	}

	return nil
}
