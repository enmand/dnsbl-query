package schema

import (
	"github.com/enmand/dnsbl-query/internal/ent/schema/mixin"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// DNSBLQuery is the edge containing query results for a given DNSBL
// check for an IP address
type DNSBLQuery struct {
	ent.Schema
}

// Mixins are the field mixins for the DNSBLQuery entity
func (DNSBLQuery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Timed{},
	}
}

// Edges are the entity edges for the DNSBLQuery in the graph.
func (DNSBLQuery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("responses", DNSBLResponse.Type),
		edge.From("ip_address", IP.Type).
			Ref("queries").
			Unique(),
	}
}
