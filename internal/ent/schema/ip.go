package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"

	"github.com/enmand/dnsbl-query/internal/ent/schema/mixin"
)

// IP is the database schema representing an IPv4 address on the graph.
// It containes ID, and timing information, and has an edge to response
// information
type IP struct {
	ent.Schema
}

// Fields are the ent field definitions of the IP entity
func (IP) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip_address"),
	}
}

// Mixins are the ent.Mixins that are used with the IP entity
func (IP) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Timed{},
	}
}

// Edges define the entity edges for the IP node in the graph
func (IP) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("queries", DNSBLQuery.Type),
	}
}
