package errors

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"project-management/models"
)

type ValidatorError struct {
	Err validator.ValidationErrors
}

func NewValidatorError(err validator.ValidationErrors) *ValidatorError {
	return &ValidatorError{
		Err: err,
	}
}

func (e ValidatorError) Error() string {
	return e.Err.Error()
}

func (e ValidatorError) ToZapFields() []zap.Field {
	fields := make([]zap.Field, len(e.Err))

	for i, err := range e.Err {
		fields[i] = zap.Any(err.Field(), err.Error())
	}

	return fields
}

func (e ValidatorError) ToErrorResponse() *models.ErrorResponse {
	details := make([]models.Validation, len(e.Err))

	for i, err := range e.Err {
		field := toSnakeCase(err.Field())
		message := toReadableMessage(field, err.Tag(), err.Kind())

		details[i] = models.Validation{
			Field:   field,
			Message: message,
			Value:   err.Value(),
		}
	}

	return &models.ErrorResponse{
		Message: e.Err.Error(),
		Details: details,
	}
}

func toSnakeCase(s string) string {
	// Replace spaces and punctuation with underscores
	input := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return '_'
		}
		return r
	}, s)

	// Convert to lowercase
	input = strings.ToLower(input)

	// Remove consecutive underscores
	for strings.Contains(input, "__") {
		input = strings.ReplaceAll(input, "__", "_")
	}

	// Remove leading and trailing underscores
	input = strings.Trim(input, "_")

	return input
}

func toReadableMessage(field, tag string, kind reflect.Kind) string {
	switch tag {
	case "required":
		return "The " + field + " field is required."
	case "email":
		return "The " + field + " field must be a valid email address."
	case "min":
		return "The " + field + " field must be at least " + tag + " characters."
	case "max":
		return "The " + field + " field must not be greater than " + tag + " characters."
	case "eqfield":
		return "The " + field + " field must be equal to " + tag + " field."
	case "nefield":
		return "The " + field + " field must not be equal to " + tag + " field."
	case "ltfield":
		return "The " + field + " field must be less than " + tag + " field."
	case "ltefield":
		return "The " + field + " field must be less than or equal to " + tag + " field."
	case "gtfield":
		return "The " + field + " field must be greater than " + tag + " field."
	case "gtefield":
		return "The " + field + " field must be greater than or equal to " + tag + " field."
	case "oneof":
		return "The " + field + " field must be one of " + tag + "."
	case "unique":
		return "The " + field + " field must be unique."
	case "uuid", "uuid3", "uuid4", "uuid5":
		return "The " + field + " field must be a valid UUID."
	case "ip":
		return "The " + field + " field must be a valid IP address."
	case "url":
		return "The " + field + " field must be a valid URL."
	case "uri":
		return "The " + field + " field must be a valid URI."
	case "boolean":
		return "The " + field + " field must be a valid boolean."
	case "numeric":
		return "The " + field + " field must be a valid numeric."
	default:
		return "The " + field + " field is invalid."
	}
}
