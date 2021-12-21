package geopoint

import "math"

const earthRadiusInKilometres = 6371

// Haversine calculates distance between two GeoPoints using Haversine formula.
// https://en.wikipedia.org/wiki/Haversine_formula
func Haversine(one, two GeoPoint) Kilometres {
	lat1 := one.Latitude.Radians()
	lng1 := one.Longitude.Radians()
	lat2 := two.Latitude.Radians()
	lng2 := two.Longitude.Radians()

	deltaLng := float64(lng2 - lng1)
	deltaLat := float64(lat2 - lat1)
	a := math.Pow(math.Sin(deltaLat/2), 2.0) + math.Cos(float64(lat1))*
		math.Cos(float64(lat2))*math.Pow(math.Sin(deltaLng/2), 2.0)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	return Kilometres(earthRadiusInKilometres * c)
}
