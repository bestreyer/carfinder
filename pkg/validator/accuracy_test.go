package validator

import (
	"testing"
)

func TestValidateAccuracyVal(t *testing.T) {
	tables := []struct {
		val float64
		exp bool
	}{
		{0, true},
		{-0.1, false},
		{1.1, false},
		{1, true},
		{0.5, true},
		{0.1, true},
	}

	for _, r := range tables {
		if act := validateAccuracyVal(r.val); r.exp != act {
			t.Errorf("validateAccuracyVal(%f) function is incorrect, actual: %t, expected: %t", r.val, act, r.exp)
		}
	}
}
