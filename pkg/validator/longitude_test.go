package validator

import "testing"

func TestValidateLongitudeVal(t *testing.T) {
	tables := []struct {
		val float64
		exp bool
	}{
		{0, true},
		{-180.1, false},
		{180.1, false},
		{180, true},
		{-180, true},
		{-180.1, false},
	}

	for _, r := range tables {
		if act := validateLongitudeVal(r.val); r.exp != act {
			t.Errorf("validateLongitudeVal(%f) function is incorrect, actual: %t, expected: %t", r.val, act, r.exp)
		}
	}
}
