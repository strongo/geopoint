package geopoint

const mileInKilometres = 1.60934 // kilometres <-> miles conversion.

// Kilometres are units of distance.
type Kilometres float64

// Miles converts Kilometres to Miles.
func (kilometers Kilometres) Miles() Miles {
	return Miles(kilometers * mileInKilometres)
}

// Miles are units of distance.
type Miles float64

// Kilometres converts miles to kilometres.
func (miles Miles) Kilometres() Kilometres {
	return Kilometres(miles * mileInKilometres)
}
