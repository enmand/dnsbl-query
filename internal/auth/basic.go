package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/user"
)

func BasicAuth(cl *ent.Client, log *zap.SugaredLogger, next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if err := basicAuth(ctx, cl, r); err != nil {
			log.With(zap.Error(err)).Error("basic auth failed")
			http.Error(rw, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, r)
	}
}

func basicAuth(ctx context.Context, cl *ent.Client, r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		return fmt.Errorf("unable to parse basic auth")
	}

	u, err := cl.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return fmt.Errorf("querying user: %w", err)
	}

	return ComparePassword(u.Password, []byte(strings.TrimSpace(password)))
}
