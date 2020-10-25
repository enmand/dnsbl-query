package mixin

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Timed is a mixin providing time-based fields recording the creation
// and update time of an entity
type Timed struct {
	ent.Mixin
}

// Fields provides a common fieldset to entities that get mixed in
func (Timed) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Indexes provides indexes by default on time-based fields for time-based
// queries
func (Timed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("updated_at", "created_at"),
	}
}
