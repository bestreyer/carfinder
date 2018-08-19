package validation

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
)

func New(t ut.Translator) (*validator.Validate) {
	v := validator.New()

	registerLatitudeValidation(v, t)
	registerLogitudeValidation(v, t)
	registerRequiredValidation(v, t)

	return v
}
