package validator

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
)

func validateAccuracyVal(val float64) bool {
	if val < 0 || val > 1 {
		return false
	}

	return true
}

func validateAccuracyField(fl validator.FieldLevel) bool {
	return validateAccuracyVal(fl.Field().Float())
}

func registerAccuracyValidation(v *validator.Validate, t ut.Translator) {
	v.RegisterValidation("accuracy", validateAccuracyField)

	v.RegisterTranslation("accuracy", t, func(ut ut.Translator) error {
		return ut.Add("accuracy", "{0} should be between [0..1]", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("accuracy", fe.Field())
		return t
	})
}
