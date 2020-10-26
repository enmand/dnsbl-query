package mixin

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/google/uuid"
)

// ID is the mixin providing ID fields to entities, with a hook for auto-generating
// IDs to be used on create mutations to the entity
type ID struct {
	mixin.Schema
}

// Fields represents the fields that are added by the ID mixin. This is the
// UUID-based field for entity IDs
func (ID) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
	}
}
