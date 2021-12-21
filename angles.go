package geopoint

import "math"

const degToRad = math.Pi / 180.0 // degrees -> radians conversion.

// Degrees are units of angular measure.
type Degrees float64

// Radians converts Degrees to Radians.
func (degrees Degrees) Radians() Radians {
	return Radians(degrees * degToRad)
}

// Radians are units of angular measure.
type Radians float64

// Degrees converts from Radians to Degrees.
func (radians Radians) Degrees() Degrees {
	return Degrees(radians / degToRad)
}
