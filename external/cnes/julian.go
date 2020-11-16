package cnes

import "time"

var epochCNES int64 = time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
var secondsPerDay float64 = 24 * 60 * 60

// CNESJulian converts a time to days since 1950-01-01
func CNESJulian(t time.Time) float64 {
	return float64(t.UTC().Unix()-epochCNES) / secondsPerDay
}
