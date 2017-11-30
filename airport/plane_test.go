package airport

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func PlaneTakesInstructionToLand(t *testing.T){
	Convey("A plane takes the instruction to land from an airport", t, func() {
		airport := Airport.NewAirport()
		plane := Plane.NewPlane(airport)
		So(plane.Land(), ShouldEqual, true)
	})
}