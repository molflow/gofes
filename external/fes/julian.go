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

import "time"

var epochCNES int64 = time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
var secondsPerDay float64 = 24 * 60 * 60

// JulianDate converts a time to days since 1950-01-01
func JulianDate(t time.Time) float64 {
	return float64(t.UTC().Unix()-epochCNES) / secondsPerDay
}
