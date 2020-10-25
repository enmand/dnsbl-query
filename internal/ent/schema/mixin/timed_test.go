package mixin

import (
	"testing"

	gm "github.com/onsi/gomega"
)

func TestMixin_Timed(t *testing.T) {
	gm.RegisterTestingT(t)

	m := Timed{}
	fs := m.Fields()
	gm.Expect(fs).To(gm.HaveLen(2))

	is := m.Indexes()
	gm.Expect(is).To(gm.HaveLen(1))
}
