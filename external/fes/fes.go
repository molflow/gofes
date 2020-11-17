/*Package fes is a go interface for the fes c-api
  Copyright (C) 2020 Möller Data Workflow Systems AB

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Lesser General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Lesser General Public License for more details.
*/
package fes

// #cgo LDFLAGS: -lfes
// #include <stdlib.h>
// #include <fes.h>
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

// Fes ...
type Fes struct {
	handle *C.FES
}

// Tide returns the height and long period height for supplied coordinates and time
func (fes *Fes) Tide(lat float64, lon float64, time time.Time) (float64, float64, error) {
	if fes.handle == nil {
		return 0, 0, fmt.Errorf("FES is not properly initialized or already closed")
	}
	h := C.double(math.NaN())
	hLongPeriod := C.double(math.NaN())
	julian := JulianDate(time)

	ok := C.fes_core(
		*fes.handle,
		C.double(lat),
		C.double(lon),
		C.double(julian),
		&h,
		&hLongPeriod,
	)
	if ok == 0 {
		return float64(h), float64(hLongPeriod), nil
	} else if ok == 1 {
		return 0, 0, fmt.Errorf("FES returned error status for tides")
	}
	return 0, 0, fmt.Errorf("FES returned unexpected error status %v for tides", ok)
}

// Close ...
func (fes *Fes) Close() {
	C.fes_delete(*fes.handle)
	fes.handle = nil
}

// NewFes creates a new Fes interface
func NewFes(tideType TideType, mode Mode, path string) (*Fes, error) {
	var handle C.FES
	var cpath *C.char = C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	ok := C.fes_new(&handle, C.fes_enum_tide_type(tideType), C.fes_enum_access(mode), cpath)
	if ok == 0 {
		return &Fes{&handle}, nil
	} else if ok == 1 {
		return nil, fmt.Errorf("Could not create a new FES Session")
	}
	return nil, fmt.Errorf("FES returned unexpected error status %v during init", ok)
}
