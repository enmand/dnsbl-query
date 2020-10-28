package resolver

import (
	"context"
	"testing"
	"time"

	// TODO: since all tests would use sqlite3 and wrapper enttest.WithDatabase
	// would help getting rid of this dep in the file
	_ "github.com/mattn/go-sqlite3"
	gm "github.com/onsi/gomega"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/enttest"
)

func Test_dNSBLQueryResolver_ResponseCode(t *testing.T) {
	tcs := map[string]struct {
		setup  func(context.Context, *ent.Client, *ent.IP)
		hasErr bool
	}{
		"single edge": {
			setup: func(ctx context.Context, cl *ent.Client, ip *ent.IP) {
				q, _ := cl.DNSBLQuery.Create().
					SetIPAddress(ip).
					Save(ctx)
				_ = cl.DNSBLResponse.Create().
					SetCode("code").
					SetDescription("tests").
					SetQuery(q).
					SaveX(ctx)
			},
		},
		"multiple queries - find latest": {
			setup: func(ctx context.Context, cl *ent.Client, ip *ent.IP) {
				_, _ = cl.DNSBLQuery.Create().
					SetIPAddress(ip).
					SetCreatedAt(time.Now().Add(-5 * time.Minute)).
					Save(ctx)
				q, _ := cl.DNSBLQuery.Create().
					SetIPAddress(ip).
					Save(ctx)
				_ = cl.DNSBLResponse.Create().
					SetCode("code").
					SetDescription("tests").
					SetQuery(q).
					SaveX(ctx)
			},
		},
		"no edges": {
			setup:  func(ctx context.Context, cl *ent.Client, ip *ent.IP) {},
			hasErr: true,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)
			ctx := context.Background()

			cl := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cl.Close()
			ip := cl.IP.Create().SetIPAddress("127.0.0.1").SaveX(ctx)
			tc.setup(ctx, cl, ip)

			r := &iPResolver{
				Resolver: &Resolver{client: cl},
			}

			code, err := r.ResponseCode(ctx, ip)
			if tc.hasErr {
				gm.Expect(err).To(gm.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gm.HaveOccurred())
				gm.Expect(code).To(gm.Equal("code"))
			}
		})
	}
}
