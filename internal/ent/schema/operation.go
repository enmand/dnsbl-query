package schema

import (
	"github.com/enmand/dnsbl-query/internal/ent/schema/mixin"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Operation is a handler to an async task that is performed outside of the web
// context.
//
// In reality, the background worker should be something like Cadence or Temporal
// but in the interest of time and simplicity a tracking Operation gives us the
// same end result
type Operation struct {
	ent.Schema
}

// Fields represent the fields on the Operation
func (Operation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("IPDNSBL"),
		field.String("ip_address").Optional(),
		field.Enum("status").
			Values("WAITING", "IN_PROGRESS", "DONE", "ERROR"),
		field.String("error").Optional(),
		field.Time("done_at").Optional(),
	}
}

// Mixin are the ent.Mixins that are used with the IP entity
func (Operation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Timed{},
	}
}
