package validator

import (
	"testing"
)

func TestValidateLatitudeVal(t *testing.T) {
	tables := []struct {
		val float64
		exp bool
	}{
		{0, true},
		{-90.1, false},
		{90.1, false},
		{90, true},
		{-90, true},
		{-90.1, false},
	}

	for _, r := range tables {
		if act := validateLatitudeVal(r.val); r.exp != act {
			t.Errorf("validateLatitudeVal(%f) function is incorrect, actual: %t, expected: %t", r.val, act, r.exp)
		}
	}
}
