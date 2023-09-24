package validators

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

var Module = fx.Provide(
	NewPagingValidator,
	NewMemberValidator,
)
