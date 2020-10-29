package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/enttest"
	gm "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func Test_BasicAuth(t *testing.T) {
	authedReq := func(c *http.Client, url string) (*http.Response, error) {
		req, _ := http.NewRequest(http.MethodGet, url+"/path", nil)
		req.SetBasicAuth("username", "pw")
		return c.Do(req)
	}

	tcs := map[string]struct {
		client   *ent.Client
		setup    func(context.Context, *ent.Client)
		req      func(*http.Client, string) (*http.Response, error)
		hasError bool
	}{
		"ok": {
			client: enttest.Open(t, "sqlite3", "file:ok?mode=memory&cache=shared&_fk=1"),
			setup: func(ctx context.Context, cl *ent.Client) {
				_ = CreateUser(ctx, cl, "username", "pw")
			},
			req: authedReq,
		},
		"bad auth": {
			client:   enttest.Open(t, "sqlite3", "file:bad_auth?mode=memory&cache=shared&_fk=1"),
			req:      authedReq,
			hasError: true,
		},
		"no auth": {
			req: func(c *http.Client, url string) (*http.Response, error) {
				return c.Get(url + "/path")
			},
			hasError: true,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)
			ctx := context.Background()

			if tc.setup != nil {
				tc.setup(ctx, tc.client)
			}

			log := zap.NewNop().Sugar()
			s := httptest.NewServer(
				BasicAuth(tc.client, log, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					_, _ = rw.Write([]byte("ok"))
				})),
			)
			defer s.Close()

			cl := s.Client()

			if tc.req != nil {
				r, _ := tc.req(cl, s.URL)
				if tc.hasError {
					gm.Expect(r.StatusCode).To(gm.Equal(http.StatusUnauthorized))
				} else {
					gm.Expect(r.StatusCode).To(gm.Equal(http.StatusOK))
				}
			} else {
				t.Fatal("no request defined")
			}
		})
	}
}
