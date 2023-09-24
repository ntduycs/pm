package validators

import (
	"github.com/go-playground/validator/v10"

	pmerror "project-management/errors"
)

type MemberValidator struct {
	pagingValidator *PagingValidator
}

func NewMemberValidator(pagingValidator *PagingValidator) *MemberValidator {
	return &MemberValidator{
		pagingValidator: pagingValidator,
	}
}

func (v *MemberValidator) Validate(req any) *pmerror.ValidatorError {
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	return pmerror.NewValidatorError(errs.(validator.ValidationErrors))
}
