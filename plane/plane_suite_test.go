package plane_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlane(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plane Suite")
}
