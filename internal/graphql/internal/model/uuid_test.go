package model

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/google/uuid"
	gm "github.com/onsi/gomega"
)

func TestScalar_MarshalUUID(t *testing.T) {
	tcs := map[string]uuid.UUID{
		"empty": uuid.UUID{},
		"new":   uuid.New(),
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			m := MarshalUUID(tc)

			buf := &bytes.Buffer{}
			m.MarshalGQL(buf)

			gm.Expect(buf.String()).To(gm.Equal(strconv.Quote(tc.String())))
		})
	}
}

func TestScalar_UnmarshalUUID(t *testing.T) {
	tcs := map[string]struct {
		id       string
		hasError bool
	}{
		"empty": {
			id:       "",
			hasError: true,
		},
		"new": {
			id:       uuid.New().String(),
			hasError: false,
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			id, err := UnmarshalUUID(tc.id)
			if tc.hasError {
				gm.Expect(err).To(gm.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gm.HaveOccurred())
				gm.Expect(id.String()).To(gm.Equal(tc.id))

			}
		})
	}
}
