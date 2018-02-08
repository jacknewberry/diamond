package diamond_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDiamond(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Diamond Suite")
}
