// Package schema provides entity definitions, which includes model
// definitions for schemas, entity hooks and policies and graph-like
// edges for relationships
//
// The entc compiler is used to generate type-safe code for accessing
// the database, and quering entity nodes and edges
//
// Code-generation can be done using `go generate`, and changes to
// generated code must be checked in.
//
// The IP schema represents a root node, that has edges to DNSBLQuery. The entity
// DNSBLQuery has edges to DNSBLResponse.
//
// IP represents a given (IPv4) address. DNSBLQuery represents a query to a DNBL
// service. DNSBLResponse represents the response(s) from a given query
package schema
