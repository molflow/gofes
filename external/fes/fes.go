package fes

// #include <stdlib.h>
// #include "fes.h"
import "C"
import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

// Mode ...
type Mode int8

const (
	// ModeIO reads grids from disk
	ModeIO Mode = iota
	// ModeMem loads grids into memory
	ModeMem
)

// TideType ...
type TideType int8

const (
	// OceanTide ...
	OceanTide TideType = iota
	// RadialTide ...
	RadialTide
)

// Tide ...
type Tide struct {
	Height           float64
	HeightLongPeriod float64
}

// Fes ...
type Fes struct {
	handle *C.FES
}

// Tide ...
func (fes *Fes) Tide(lat float64, lon float64, time time.Time) (Tide, error) {
	var ok int
	h := C.double(math.NaN)
	hLongPeriod := C.double(math.NaN)
	//TODO: Convert  to julian
	julian := float64(time)
	ok = C.fes_core(
		fes.handle,
		C.double(lat),
		C.double(lon),
		C.double(time),
		&h,
		&hLongPeriod,
	)
	if ok == 0 {
		return Tide{float64(h), float64(hLongPeriod)}, nil
	} else if ok == 1 {
		return nil, error.Error("FES returned error status for tides")
	}
	return nil, fmt.Errorf("FES returned unexpected error status %v for tides", ok)
}

// Close ...
func (fes *Fes) Close() {
	C.fes_delete(fes.handle)
	C.free(unsafe.Pointer(fes.handle))
	fes.handle = null
}

// NewFes creates a new Fes interface
func NewFes(tideType TideType, mode Mode, path string) *Fes {
	// TODO: Is this it??
	handle := C.FES()
	// TODO: Convert path to char *const
	cpath := C.C
	// TODO: Can cast int to enum?
	C.fes_new(handle, C.fes_enum_tide_type(tideType), C.fes_enum_access(mode))
	return &Fes{handle}
}
