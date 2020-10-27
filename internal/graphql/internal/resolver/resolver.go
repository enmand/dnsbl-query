package resolver

import (
	"github.com/99designs/gqlgen/graphql"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/graphql/internal/gen"
)

// Resolver is the root resolver
type Resolver struct {
	client *ent.Client
}

// NewSchema generates a new executable schema for graphql
func NewSchema(c *ent.Client) graphql.ExecutableSchema {
	return gen.NewExecutableSchema(gen.Config{
		Resolvers: &Resolver{client: c},
	})
}
