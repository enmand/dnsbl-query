package schema

import (
	"github.com/enmand/dnsbl-query/internal/ent/schema/mixin"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// User represents a user that can authenticate to the service, and has an
// identity.
type User struct {
	ent.Schema
}

// Fields are the ent fields on the User entity
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.Bytes("password"),
	}
}

// Indexes defines database indexes for the User entity
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
	}
}

// Mixin defines ent Mixins to use with User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Timed{},
	}
}
