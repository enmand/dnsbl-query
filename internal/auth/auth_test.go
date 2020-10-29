package auth

import (
	"context"
	"testing"

	// TODO: since all tests would use sqlite3 and wrapper enttest.WithDatabase
	// would help getting rid of this dep in the file
	scrypt "github.com/elithrar/simple-scrypt"
	_ "github.com/mattn/go-sqlite3"
	gm "github.com/onsi/gomega"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/enttest"
)

func Test_CreateUser(t *testing.T) {
	tcs := map[string]struct {
		client   *ent.Client
		setup    func(context.Context, *ent.Client)
		hasError bool
	}{
		"already exists": {
			client: enttest.Open(t, "sqlite3", "file:exists?mode=memory&cache=shared&_fk=1"),
			setup: func(ctx context.Context, cl *ent.Client) {
				cl.User.Create().SetUsername("username").SetPassword([]byte("pw")).SaveX(ctx)
			},
			hasError: true,
		},
		"success": {
			client: enttest.Open(t, "sqlite3", "file:success?mode=memory&cache=shared&_fk=1"),
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			ctx := context.Background()
			defer tc.client.Close()
			tx, err := tc.client.Tx(ctx)
			gm.Expect(err).NotTo(gm.HaveOccurred())
			cl := tx.Client()

			if tc.setup != nil {
				tc.setup(ctx, cl)
			}

			err = CreateUser(ctx, cl, "username", "pw")
			if tc.hasError {
				gm.Expect(err).To(gm.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gm.HaveOccurred())
			}
		})
	}
}

func Test_ComparePasswords(t *testing.T) {
	tcs := map[string]struct {
		pass     string
		hash     func() []byte
		hasError bool
	}{
		"bad": {
			pass:     "pass",
			hash:     func() []byte { return []byte("not a hash") },
			hasError: true,
		},
		"ok": {
			pass: "pass",
			hash: func() []byte {
				hash, _ := scrypt.GenerateFromPassword([]byte("pass"), scrypt.DefaultParams)
				return hash
			},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			err := ComparePassword(tc.hash(), []byte(tc.pass))
			if tc.hasError {
				gm.Expect(err).To(gm.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gm.HaveOccurred())
			}
		})
	}
}
