package plane_test

import (
	. "github.com/cgulli/airport_go/plane"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/golang/mock/gomock"
	"github.com/cgulli/airport_go/mocks"
)

var _ = Describe("Plane", func() {

	var(
		mockCtrl *gomock.Controller
		mockAirport *mocks.MockAirport
	)

	Describe("Landing", func() {
		It("Should ask airport if it can land", func() {
			mockCtrl = gomock.NewController(GinkgoT())
			defer mockCtrl.Finish()
			mockAirport = mocks.NewMockAirport(mockCtrl)

			mockAirport.EXPECT().CanLand().Return(true).Times(1)

			plane := Plane{mockAirport}
			Expect(plane.Land()).To(Equal(true))
		})
	})


})


