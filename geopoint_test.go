package geopoint

import (
	"reflect"
	"testing"
)

func TestDegrees_Radians(t *testing.T) {
	//goland:noinspection GoPreferNilSlice
	tests := []struct {
		name string
		self Degrees
		want Radians
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Radians(); got != tt.want {
				t.Errorf("Radians() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeoPoint_DistanceTo(t *testing.T) {
	type fields struct {
		Latitude  Degrees
		Longitude Degrees
	}
	type args struct {
		another GeoPoint
		f       Formula
	}
	//goland:noinspection GoPreferNilSlice
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Kilometres
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeoPoint{
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			if got := g.DistanceTo(tt.args.another, tt.args.f); got != tt.want {
				t.Errorf("DistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHaversine(t *testing.T) {
	type args struct {
		one GeoPoint
		two GeoPoint
	}
	//goland:noinspection GoPreferNilSlice
	tests := []struct {
		name string
		args args
		want Kilometres
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Haversine(tt.args.one, tt.args.two); got != tt.want {
				t.Errorf("Haversine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKilometres_Miles(t *testing.T) {
	//goland:noinspection GoPreferNilSlice
	tests := []struct {
		name string
		self Kilometres
		want Miles
	}{
		// TODO: Add test cases.
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
	//goland:noinspection GoPreferNilSlice
	tests := []struct {
		name string
		self Miles
		want Kilometres
	}{
		// TODO: Add test cases.
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
	//goland:noinspection GoPreferNilSlice
	tests := []struct {
		name string
		args args
		want *GeoPoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGeoPoint(tt.args.latitude, tt.args.longitude); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGeoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadians_Degrees(t *testing.T) {
	//goland:noinspection GoPreferNilSlice
	tests := []struct {
		name string
		self Radians
		want Degrees
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Degrees(); got != tt.want {
				t.Errorf("Degrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
