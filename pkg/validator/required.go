package validator

import (
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func registerRequiredValidation(v *validator.Validate, t ut.Translator) {
	v.RegisterTranslation("required", t, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}
