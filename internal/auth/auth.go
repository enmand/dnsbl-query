package auth

import (
	"context"
	"fmt"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"

	scrypt "github.com/elithrar/simple-scrypt"
)

var params scrypt.Params

func init() {
	params = scrypt.DefaultParams
}

// Create user creates a new user in the authentication system
func CreateUser(ctx context.Context, cl *ent.Client, username, password string) error {
	hash, err := scrypt.GenerateFromPassword([]byte(password), params)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	_, err = cl.User.Create().
		SetUsername(username).
		SetPassword(hash).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	return nil
}
