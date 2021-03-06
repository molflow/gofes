package fes

import (
	"math"
	"testing"
	"time"
)

const inifile = "../../third-party/fes2014-data/fes.ini"

func approx(a float64, b float64, tol float64) bool {
	return math.Abs(a-b) < tol
}

func TestNewFes(t *testing.T) {
	type args struct {
		tideType TideType
		mode     Mode
		path     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Missing ini should error",
			args{
				OceanTide,
				ModeIO,
				"not-a-fes.ini",
			},
			true,
		},
		{
			"Should return a fes",
			args{
				OceanTide,
				ModeIO,
				inifile,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFes(tt.args.tideType, tt.args.mode, tt.args.path)
			if got != nil {
				defer got.Close()
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewFes() should have returned a Fes")
			}
		})
	}
}

func TestFes_Tide(t *testing.T) {
	type fesSettings struct {
		tide TideType
		mode Mode
	}
	type args struct {
		lat  float64
		lon  float64
		time time.Time
	}
	tests := []struct {
		name            string
		fesSettings     fesSettings
		args            args
		wantH           float64
		wantHLongPeriod float64
		wantErr         bool
	}{
		{
			"Get ocean tide mem-mode",
			fesSettings{OceanTide, ModeMem},
			args{
				0,
				0,
				time.Date(2020, 11, 16, 13, 0, 0, 0, time.UTC),
			},
			-27.9,
			-0.71,
			false,
		},
		{
			"Get ocean tide changes with lat",
			fesSettings{OceanTide, ModeMem},
			args{
				1,
				0,
				time.Date(2020, 11, 16, 13, 0, 0, 0, time.UTC),
			},
			-30.16,
			-0.72,
			false,
		},
		{
			"Get ocean tide changes with lon",
			fesSettings{OceanTide, ModeMem},
			args{
				0,
				2,
				time.Date(2020, 11, 16, 13, 0, 0, 0, time.UTC),
			},
			-24.11,
			-0.71,
			false,
		},
		{
			"Changes with time",
			fesSettings{OceanTide, ModeMem},
			args{
				0,
				0,
				time.Date(2020, 11, 16, 18, 0, 0, 0, time.UTC),
			},
			51.97,
			-0.96,
			false,
		},
		{
			"Get ocean tide io-mode",
			fesSettings{OceanTide, ModeIO},
			args{
				0,
				0,
				time.Date(2020, 11, 16, 13, 0, 0, 0, time.UTC),
			},
			-27.9,
			-0.71,
			false,
		},
		{
			"Get radial tide mem-mode",
			fesSettings{RadialTide, ModeMem},
			args{
				0,
				0,
				time.Date(2020, 11, 16, 13, 0, 0, 0, time.UTC),
			},
			1.8,
			0,
			false,
		},
		{
			"Get radial tide io-mode",
			fesSettings{RadialTide, ModeIO},
			args{
				0,
				0,
				time.Date(2020, 11, 16, 13, 0, 0, 0, time.UTC),
			},
			1.8,
			0,
			false,
		},
		{
			"Continental Point of Inaccessibility errors because very much on land",
			fesSettings{OceanTide, ModeMem},
			args{
				46.283333,
				86.666667,
				time.Date(2020, 11, 16, 13, 0, 0, 0, time.UTC),
			},
			0,
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fes, initerr := NewFes(tt.fesSettings.tide, tt.fesSettings.mode, inifile)
			if initerr != nil {
				t.Errorf("Could not init the FES")
			} else {
				defer fes.Close()
				height, heightLongPeriod, err := fes.Tide(tt.args.lat, tt.args.lon, tt.args.time)
				if (err != nil) != tt.wantErr {
					t.Errorf("Fes.Tide() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !approx(height, tt.wantH, 0.01) || !approx(heightLongPeriod, tt.wantHLongPeriod, 0.01) {
					t.Errorf("Fes.Tide() = %v, %v, want %v, %v", height, heightLongPeriod, tt.wantH, tt.wantHLongPeriod)
				}
			}
		})
	}
}
