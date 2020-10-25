// +build ignore

// Package main provides a custom entc compiler, which includes the
// ent-contrib templates for gqlgen
package main

import (
	"log"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entgql"
)

// TODO: pkg should be configurable, and/or set using go.mod name
const pkg = "github.com/enmand/dnsbl-query/internal/ent/gen/ent"

func main() {
	err := entc.Generate("./schema", &gen.Config{
		IDType:    &field.TypeInfo{Type: field.TypeString},
		Target:    "./gen/ent/",
		Package:   pkg,
		Templates: entgql.AllTemplates,
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
