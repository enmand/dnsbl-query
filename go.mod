module github.com/enmand/dnsbl-query

go 1.15

require (
	// gqlgen is for resolver generation, and gql service scaffolding
	github.com/99designs/gqlgen v0.13.0
	// ent is an entity relationship library, which gqlgen support via ent-contrib
	github.com/facebook/ent v0.4.4-0.20201018111720-17ee19e23a5c
	// ent-contrib provides ent and gqlgen templates for codegen
	github.com/facebookincubator/ent-contrib v0.0.0-20201022175118-63295bc19d1b
	// golangci-lint preforms linting in CI/CD or development phases
	github.com/golangci/golangci-lint v1.32.0
	// uuid is used to generate new UUID-based IDs for entities
	github.com/google/uuid v1.1.2
	// go-multierror is used in entc generated code for Relay Node Interface support
	github.com/hashicorp/go-multierror v1.1.0
	// Mage is a pure Go-based replacement for make
	github.com/magefile/mage v1.10.0
	// gqlparser/v2 is used in entc generated code for pagination
	github.com/vektah/gqlparser/v2 v2.1.0
	// msgpack is used in entc generated code for pagination
	github.com/vmihailenco/msgpack/v5 v5.0.0-beta.1
)
