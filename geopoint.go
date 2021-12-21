package geopoint

// Formula to calculate distance between two GeoPoints.
type Formula func(one, two GeoPoint) Kilometres

// GeoPoint represents GPS coordinates. Has no filtered or unexported fields.
type GeoPoint struct {

	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude Degrees `json:"latitude,omitempty"`

	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude Degrees `json:"longitude,omitempty"`
}

// NewGeoPoint returns a pointer to a new GeoPoint.
func NewGeoPoint(latitude, longitude Degrees) *GeoPoint {
	return &GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// DistanceTo calculates distance to another GeoPoint.
func (g GeoPoint) DistanceTo(another GeoPoint, f Formula) Kilometres {
	return f(g, another)
}
