package diamond_test

import (
	. "diamond"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Diamond", func() {
	It("prints all the diamonds", func() {
		Expect(ToDiamond("H")).To(Equal([]string{
			"       A",
			"      B B",
			"     C   C",
			"    D     D",
			"   E       E",
			"  F         F",
			" G           G",
			"H             H",
			" G           G",
			"  F         F",
			"   E       E",
			"    D     D",
			"     C   C",
			"      B B",
			"       A",
		}))

	})
})
