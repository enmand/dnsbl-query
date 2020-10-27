package mixin

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/facebook/ent/schema/mixin"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// Timed is a mixin providing time-based fields recording the creation
// and update time of an entity
type Timed struct {
	mixin.Schema
}

// Fields provides a common fieldset to entities that get mixed in
func (Timed) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Indexes provides indexes by default on time-based fields for time-based
// queries
func (Timed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("updated_at", "created_at"),
	}
}
