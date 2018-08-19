package validator

import (
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func New(t ut.Translator) *validator.Validate {
	v := validator.New()

	registerLatitudeValidation(v, t)
	registerLogitudeValidation(v, t)
	registerRequiredValidation(v, t)
	registerAccuracyValidation(v, t)

	return v
}
