package fes

import (
	"testing"
	"time"
)

var loc, _ = time.LoadLocation("CET")

func TestJulianDate(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want float64
	}{
		{
			"Epoch is 0",
			time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC),
			0,
		},
		{
			"Counts days",
			time.Date(1950, time.January, 2, 12, 0, 0, 0, time.UTC),
			1.5,
		},
		{
			"Allows negatives",
			time.Date(1949, time.December, 31, 12, 0, 0, 0, time.UTC),
			-0.5,
		},
		{
			"Handles timezones",
			time.Date(1950, time.January, 1, 12, 0, 0, 0, loc),
			11. / 24.,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JulianDate(tt.t); got != tt.want {
				t.Errorf("JulianDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJulianDate_sufficient_resolution(t *testing.T) {
	d1 := time.Date(2030, 4, 15, 10, 15, 0, 0, time.UTC)
	d2 := d1.Add(time.Duration(1) * time.Minute)
	first := JulianDate(d1)
	second := JulianDate(d2)
	if first == second {
		t.Errorf("JulianDate() equated %v == %v as %v", d1, d2, first)
	}
}
