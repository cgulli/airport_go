package airport_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAirport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Airport Suite")
}
