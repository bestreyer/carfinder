package validation

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
)

func validateLatitude(fl validator.FieldLevel) bool {
	val := fl.Field().Float()

	if val > 90 || val < -90 {
		return false
	}

	return true
}

func registerLatitudeValidation(v *validator.Validate, t ut.Translator) {
	v.RegisterValidation("latitude", validateLatitude)

	v.RegisterTranslation("latitude", t, func(ut ut.Translator) error {
		return ut.Add("latitude", "{0} should be between +/- 90", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("latitude", fe.Field())
		return t
	})
}
