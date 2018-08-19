package validation

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
)

func validateLongitudeVal(val float64) bool {
	if val > 180 || val < -180 {
		return false
	}

	return true
}

func validateLongitude(fl validator.FieldLevel) bool {
	return validateLongitudeVal(fl.Field().Float())
}

func registerLogitudeValidation(v *validator.Validate, t ut.Translator) {
	v.RegisterValidation("longitude", validateLongitude)

	v.RegisterTranslation("longitude", t, func(ut ut.Translator) error {
		return ut.Add("longitude", "{0} should be between +/- 180", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("longitude", fe.Field())
		return t
	})
}
