// Package graphql provides GraphQL schema, code-generation for gqlgen,
// resolver definitions and any necessary model definitions where ent cannot
// be autobinded.
//
// entgql from ent-contrib is used for pagination, and Relay Node Support.
// Transactional mutations are available, but not currently used, as enqueue only
// preforms a single database update.
package graphql
