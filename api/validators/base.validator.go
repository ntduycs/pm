package validators

import (
	"github.com/go-playground/validator/v10"

	pmerror "project-management/errors"
)

type RequestValidator struct {
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{}
}

func (v *RequestValidator) Validate(req any) *pmerror.ValidatorError {
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	return pmerror.NewValidatorError(errs.(validator.ValidationErrors))
}
