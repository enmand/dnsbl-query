package admin

import (
	"context"
	"fmt"
	"log"

	gflags "github.com/jessevdk/go-flags"

	"github.com/enmand/dnsbl-query/internal/auth"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/migrate"
	"github.com/enmand/dnsbl-query/internal/flags"
)

type cmd struct{}
type addUserCmd struct {
	flags.Database
}

// Register registers the administrative commands for CLI admin tooling
func Register(p *gflags.Parser) {
	c, err := p.AddCommand("admin", "administrative commands", "", &cmd{})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.AddCommand("add-user", "", "add a new user", &addUserCmd{})
	if err != nil {
		log.Fatal(err)
	}
}

// Execute executes the admin base command
func (c *cmd) Execute(args []string) error {
	return fmt.Errorf("no subcommand provided")
}

func (c *addUserCmd) Execute(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("%d invalid number of arguments: required 2", len(args))
	}
	ctx := context.Background()

	// TODO: ent open logic with options+migration should be in common pkg
	cl, err := ent.Open(c.DatabaseDriver, c.DatabaseURI)
	if err != nil {
		return fmt.Errorf("")
	}
	err = cl.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true))
	if err != nil {
		return fmt.Errorf("migration: %w", err)
	}

	return auth.CreateUser(ctx, cl, args[0], args[1])
}
