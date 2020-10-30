package dnsbl

import (
	"context"
	"fmt"
	"net"
	"testing"

	gm "github.com/onsi/gomega"
)

func Test_NewSpamhaus(t *testing.T) {
	tcs := map[string]struct {
		Options []Option
	}{
		"no options": {},
		"custom dns server": {
			Options: []Option{
				WithDNSServer("1.1.1.1"),
			},
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			s := NewSpamhaus(tc.Options...)
			gm.Expect(s).To(gm.BeAssignableToTypeOf(&spamhaus{}))
		})
	}
	gm.RegisterTestingT(t)
}

type mockResolver struct {
	resp []net.IP
	err  error
}

func (m *mockResolver) LookupIP(_ context.Context, _, _ string) ([]net.IP, error) {
	return m.resp, m.err
}

func Test_spamhaus_Query(t *testing.T) {
	tcs := map[string]struct {
		resolver resolver
		expected *Response
		hasErr   bool
	}{
		"ok": {
			resolver: &mockResolver{
				resp: []net.IP{net.ParseIP("127.0.0.1")},
			},
			expected: &Response{Codes: []Code{"127.0.0.1"}},
		},
		"query err": {
			resolver: &mockResolver{
				err: fmt.Errorf("error"),
			},
			hasErr: true,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			ctx := context.Background()
			sh := &spamhaus{
				network:  "test",
				resolver: tc.resolver,
			}

			r, err := sh.Query(ctx, "127.0.0.1")
			if tc.hasErr {
				gm.Expect(err).To(gm.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gm.HaveOccurred())
				gm.Expect(r).To(gm.Equal(tc.expected))
			}
		})
	}
}

func Test_reverseIP(t *testing.T) {
	tcs := map[string]struct {
		ip      net.IP
		reverse net.IP
	}{
		"ok": {
			ip:      net.ParseIP("127.0.0.1"),
			reverse: net.ParseIP("1.0.0.127"),
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)
			out := reverseIP(tc.ip.String())
			gm.Expect(out).To(gm.Equal(tc.reverse.String()))
		})
	}
}
