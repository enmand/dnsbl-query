// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// DnsblQueriesColumns holds the columns for the "dnsbl_queries" table.
	DnsblQueriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ip_queries", Type: field.TypeInt, Nullable: true},
	}
	// DnsblQueriesTable holds the schema information for the "dnsbl_queries" table.
	DnsblQueriesTable = &schema.Table{
		Name:       "dnsbl_queries",
		Columns:    DnsblQueriesColumns,
		PrimaryKey: []*schema.Column{DnsblQueriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dnsbl_queries_ips_queries",
				Columns: []*schema.Column{DnsblQueriesColumns[1]},

				RefColumns: []*schema.Column{IpsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DnsblResponsesColumns holds the columns for the "dnsbl_responses" table.
	DnsblResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "dnsbl_query_responses", Type: field.TypeInt, Nullable: true},
	}
	// DnsblResponsesTable holds the schema information for the "dnsbl_responses" table.
	DnsblResponsesTable = &schema.Table{
		Name:       "dnsbl_responses",
		Columns:    DnsblResponsesColumns,
		PrimaryKey: []*schema.Column{DnsblResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dnsbl_responses_dnsbl_queries_responses",
				Columns: []*schema.Column{DnsblResponsesColumns[3]},

				RefColumns: []*schema.Column{DnsblQueriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// IpsColumns holds the columns for the "ips" table.
	IpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ip_address", Type: field.TypeString},
	}
	// IpsTable holds the schema information for the "ips" table.
	IpsTable = &schema.Table{
		Name:        "ips",
		Columns:     IpsColumns,
		PrimaryKey:  []*schema.Column{IpsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DnsblQueriesTable,
		DnsblResponsesTable,
		IpsTable,
	}
)

func init() {
	DnsblQueriesTable.ForeignKeys[0].RefTable = IpsTable
	DnsblResponsesTable.ForeignKeys[0].RefTable = DnsblQueriesTable
}
