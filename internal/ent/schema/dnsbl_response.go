package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"

	"github.com/enmand/dnsbl-query/internal/ent/schema/mixin"
)

// DNSBLResponse is the response from the DNSBL service
type DNSBLResponse struct {
	ent.Schema
}

// Fields are the ent field definitions for the DNSBLResponse entity
func (DNSBLResponse) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("description"),
	}
}

// Mixins are the field mixins for the DNSBLResponse entity
func (DNSBLResponse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Timed{},
	}
}

// Edges are the entity edges for the DNSBLResponse entity in the graph
func (DNSBLResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("query", DNSBLQuery.Type).
			Ref("responses").
			Annotations(entgql.Bind()).
			Unique().
			Required(),
	}
}
