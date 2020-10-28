package graphql

import (
	"log"

	"github.com/jessevdk/go-flags"

	"github.com/enmand/dnsbl-query/internal/graphql"
)

type cmd struct{}

// Register registers the graphql server command to the binary
func Register(p *flags.Parser) {
	c, err := p.AddCommand("graphql", "run graphql", "", &cmd{})
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.AddGroup("server", "server options", &graphql.Flags)
	if err != nil {
		log.Fatalln(err)
	}
}

// Execute creates and starts a new GraphQL service
func (*cmd) Execute(args []string) error {
	s, err := graphql.New()
	if err != nil {
		log.Fatalf("unable to start server: %s\n", err)
	}

	return s.Start()
}
