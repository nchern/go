package dsp

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	in := []float64{2, 3, 4}
	expected := []float64{2. / 9, 1. / 3, 4. / 9}
	r := Normalize(in)
	for i, v := range r {
		if v != expected[i] {
			t.Error("expected: ", expected[i], "actual: ", v)
		}
	}

}
