# gofes
Golang bindings for CNES AVISO's FES / Tides library

This implements the interface of fes v2.9.3

# Requirements

You need the libfes.so installed on your system:

https://bitbucket.org/cnes_aviso/fes/src/master/

You will also need fes-compatible data.
For more information, check the link above.

# Test - Requirements

You need a `fes.ini` and suitable nc-files inside the directory `./third-party/fes2014-data`

Run complete tests with `make`.

# Use example

```go
import (
    "time"
)

func doSomethingWithTides() {
    fes, errInit := NewFes(fes.OceanTide, fes.ModeMem, "path/to/my/fes.ini")
    // handle initiation errors
    defer fes.Close()
    lat := 0.0
    lon := 0.0
    t := time.Now()
    height, heightLongPeriod, err := fes.Tide(lat, lon, t)
    // handle tide getting errors

    // Do the stuff...
}
```
