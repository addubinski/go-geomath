package geomath

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDegreeConversion(t *testing.T) {

	opt := cmp.Comparer(func(x, y float64) bool {
		a := math.Round(x*math.Pow(10, 5)) / math.Pow(10, 5)
		b := math.Round(y*math.Pow(10, 5)) / math.Pow(10, 5)
		return a == b
	})

	dd := MinutesSecondsToDecimal(32, 45, 28.0)
	if !cmp.Equal(dd, 32.7577778, opt) {
		t.Errorf("Incorrect conversion from minute second to decimal degrees result: %f", dd)
	}

	d, m, s := DecimalToMinutesSeconds(dd)

	if d != 32 {
		t.Errorf("Incorrect conversion from decimal to degrees result: %d", d)
	}

	if !cmp.Equal(float64(m), 45.0, opt) {
		t.Errorf("Incorrect conversion from decimal to minutes result: %d", m)
	}

	if !cmp.Equal(s, 28.0, opt) {
		t.Errorf("Incorrect conversion from decimal to seconds result: %f", s)
	}
}
