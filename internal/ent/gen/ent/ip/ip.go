// Code generated by entc, DO NOT EDIT.

package ip

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ip type in the database.
	Label = "ip"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIPAddress holds the string denoting the ip_address field in the database.
	FieldIPAddress = "ip_address"

	// EdgeQueries holds the string denoting the queries edge name in mutations.
	EdgeQueries = "queries"

	// Table holds the table name of the ip in the database.
	Table = "ips"
	// QueriesTable is the table the holds the queries relation/edge.
	QueriesTable = "dnsbl_queries"
	// QueriesInverseTable is the table name for the DNSBLQuery entity.
	// It exists in this package in order to avoid circular dependency with the "dnsblquery" package.
	QueriesInverseTable = "dnsbl_queries"
	// QueriesColumn is the table column denoting the queries relation/edge.
	QueriesColumn = "ip_queries"
)

// Columns holds all SQL columns for ip fields.
var Columns = []string{
	FieldID,
	FieldIPAddress,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
