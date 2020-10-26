package mixin

import (
	"testing"

	gm "github.com/onsi/gomega"
)

func TestMixin_ID(t *testing.T) {
	gm.RegisterTestingT(t)

	id := ID{}
	f := id.Fields()
	gm.Expect(f).To(gm.HaveLen(1))
	gm.Expect(f[0].Descriptor().Name).To(gm.Equal("id"))
}
