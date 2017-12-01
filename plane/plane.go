package plane

import . "github.com/cgulli/airport_go/airport"

type Plane struct {

	LandingAirport Airport
}

func (p Plane) Land() bool {
	return p.LandingAirport.CanLand()
}