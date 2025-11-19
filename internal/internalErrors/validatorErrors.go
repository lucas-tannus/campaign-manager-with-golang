package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateDomain(obj interface{}) error {
	validate := validator.New()

	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	firstError := err.(validator.ValidationErrors)[0]
	field := strings.ToLower(firstError.Field())

	switch firstError.Tag() {
	case "min":
		return errors.New(field + " should have at least " + firstError.Param() + " characters")
	case "max":
		return errors.New(field + " should have less than " + firstError.Param() + " characters")
	case "required":
		return errors.New(field + " is required")
	case "email":
		return errors.New(field + " is invalid")
	}
	return nil
}
