package validators

type PagingValidator struct {
}

func NewPagingValidator() *PagingValidator {
	return &PagingValidator{}
}

func (v *PagingValidator) Validate(req any) error {
	return validate.Struct(req)
}
