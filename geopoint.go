package geopoint

// Formula to calculate distance between two GeoPoints.
type Formula func(one, two *GeoPoint) Kilometres

// GeoPoint represents GPS coordinates.
type GeoPoint struct {
	Latitude, Longitude Degrees
}

// NewGeoPoint returns a pointer to a new GeoPoint.
func NewGeoPoint(latitude, longitude Degrees) *GeoPoint {
	return &GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// DistanceTo calculates distance to another GeoPoint.
func (g *GeoPoint) DistanceTo(another *GeoPoint, f Formula) Kilometres {
	return f(g, another)
}
