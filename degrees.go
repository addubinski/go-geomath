package geomath

import (
	"math"
)

func DecimalToMinutesSeconds(dd float64) (int, int, float64) {
	i, d := math.Modf(dd)
	m, d := math.Modf(d * 60.0)
	s := d * 60.0
	return int(i), int(m), s
}

func MinutesSecondsToDecimal(degrees int, minutes int, seconds float64) float64 {
	return float64(degrees) + float64(minutes)/60.0 + seconds/3600.0
}
