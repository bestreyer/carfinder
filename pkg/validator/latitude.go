package validator

import (
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func validateLatitudeVal(val float64) bool {
	if val > 90 || val < -90 {
		return false
	}

	return true
}

func validateLatitudeField(fl validator.FieldLevel) bool {
	return validateLatitudeVal(fl.Field().Float())
}

func registerLatitudeValidation(v *validator.Validate, t ut.Translator) {
	v.RegisterValidation("latitude", validateLatitudeField)

	v.RegisterTranslation("latitude", t, func(ut ut.Translator) error {
		return ut.Add("latitude", "{0} should be between +/- 90", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("latitude", fe.Field())
		return t
	})
}
