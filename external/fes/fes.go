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

	"github.com/molflow/gofes/external/cnes"
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
	if fes.handle == nil {
		return Tide{}, fmt.Errorf("")
	}
	h := C.double(math.NaN())
	hLongPeriod := C.double(math.NaN())
	julian := cnes.JulianDate(time)

	ok := C.fes_core(
		*fes.handle,
		C.double(lat),
		C.double(lon),
		C.double(julian),
		&h,
		&hLongPeriod,
	)
	if ok == 0 {
		return Tide{float64(h), float64(hLongPeriod)}, nil
	} else if ok == 1 {
		return Tide{}, fmt.Errorf("FES returned error status for tides")
	}
	return Tide{}, fmt.Errorf("FES returned unexpected error status %v for tides", ok)
}

// Close ...
func (fes *Fes) Close() {
	C.fes_delete(*fes.handle)
	C.free(unsafe.Pointer(fes.handle))
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
