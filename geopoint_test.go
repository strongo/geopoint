package geopoint

import (
	"math"
	"reflect"
	"testing"
)

func TestDegrees_Radians(t *testing.T) {
	tests := []struct {
		name string
		self Degrees
		want Radians
	}{
		{name: "0 deg", self: 0, want: 0},
		{name: "180 deg", self: 180, want: Radians(math.Pi)},
		{name: "90 deg", self: 90, want: Radians(math.Pi / 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Radians(); math.Abs(float64(got-tt.want)) > 1e-9 {
				t.Errorf("Radians() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadians_Degrees(t *testing.T) {
	tests := []struct {
		name string
		self Radians
		want Degrees
	}{
		{name: "0 rad", self: 0, want: 0},
		{name: "pi rad", self: Radians(math.Pi), want: 180},
		{name: "pi/2 rad", self: Radians(math.Pi / 2), want: 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Degrees(); math.Abs(float64(got-tt.want)) > 1e-9 {
				t.Errorf("Degrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKilometres_Miles(t *testing.T) {
	tests := []struct {
		name string
		self Kilometres
		want Miles
	}{
		{name: "0 km", self: 0, want: 0},
		{name: "10 km", self: 10, want: Miles(10 * mileInKilometres)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Miles(); got != tt.want {
				t.Errorf("Miles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiles_Kilometres(t *testing.T) {
	tests := []struct {
		name string
		self Miles
		want Kilometres
	}{
		{name: "0 mi", self: 0, want: 0},
		{name: "10 mi", self: 10, want: Kilometres(10 * mileInKilometres)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Kilometres(); got != tt.want {
				t.Errorf("Kilometres() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGeoPoint(t *testing.T) {
	type args struct {
		latitude  Degrees
		longitude Degrees
	}
	tests := []struct {
		name string
		args args
		want *GeoPoint
	}{
		{
			name: "London",
			args: args{latitude: 51.5074, longitude: -0.1278},
			want: &GeoPoint{Latitude: 51.5074, Longitude: -0.1278},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGeoPoint(tt.args.latitude, tt.args.longitude); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGeoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeoPoint_DistanceTo(t *testing.T) {
	p1 := GeoPoint{Latitude: 51.5074, Longitude: -0.1278}
	p2 := GeoPoint{Latitude: 48.8566, Longitude: 2.3522}
	called := false
	mockFormula := func(one, two GeoPoint) Kilometres {
		called = true
		if one != p1 || two != p2 {
			t.Errorf("unexpected arguments to formula: %v, %v", one, two)
		}
		return 343.5
	}
	got := p1.DistanceTo(p2, mockFormula)
	if !called {
		t.Error("expected formula to be called")
	}
	if got != 343.5 {
		t.Errorf("DistanceTo() = %v, want 343.5", got)
	}
}

func TestHaversine(t *testing.T) {
	p1 := GeoPoint{Latitude: 51.5074, Longitude: -0.1278} // London
	p2 := GeoPoint{Latitude: 48.8566, Longitude: 2.3522}  // Paris

	// Same point distance should be 0
	if dist := Haversine(p1, p1); dist != 0 {
		t.Errorf("Haversine(p1, p1) = %v, want 0", dist)
	}

	dist := Haversine(p1, p2)
	// London to Paris is ~343-344 km
	if dist < 340 || dist > 345 {
		t.Errorf("Haversine(London, Paris) = %v km, expected ~343-344 km", dist)
	}
}
