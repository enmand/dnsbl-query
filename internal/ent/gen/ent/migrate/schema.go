// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// DnsblQueriesColumns holds the columns for the "dnsbl_queries" table.
	DnsblQueriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "ip_queries", Type: field.TypeUUID, Nullable: true},
	}
	// DnsblQueriesTable holds the schema information for the "dnsbl_queries" table.
	DnsblQueriesTable = &schema.Table{
		Name:       "dnsbl_queries",
		Columns:    DnsblQueriesColumns,
		PrimaryKey: []*schema.Column{DnsblQueriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dnsbl_queries_ips_queries",
				Columns: []*schema.Column{DnsblQueriesColumns[3]},

				RefColumns: []*schema.Column{IpsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dnsblquery_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{DnsblQueriesColumns[2], DnsblQueriesColumns[1]},
			},
		},
	}
	// DnsblResponsesColumns holds the columns for the "dnsbl_responses" table.
	DnsblResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "dnsbl_query_responses", Type: field.TypeUUID, Nullable: true},
	}
	// DnsblResponsesTable holds the schema information for the "dnsbl_responses" table.
	DnsblResponsesTable = &schema.Table{
		Name:       "dnsbl_responses",
		Columns:    DnsblResponsesColumns,
		PrimaryKey: []*schema.Column{DnsblResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dnsbl_responses_dnsbl_queries_responses",
				Columns: []*schema.Column{DnsblResponsesColumns[5]},

				RefColumns: []*schema.Column{DnsblQueriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dnsblresponse_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{DnsblResponsesColumns[2], DnsblResponsesColumns[1]},
			},
		},
	}
	// IpsColumns holds the columns for the "ips" table.
	IpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "ip_address", Type: field.TypeString},
	}
	// IpsTable holds the schema information for the "ips" table.
	IpsTable = &schema.Table{
		Name:        "ips",
		Columns:     IpsColumns,
		PrimaryKey:  []*schema.Column{IpsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "ip_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{IpsColumns[2], IpsColumns[1]},
			},
		},
	}
	// OperationsColumns holds the columns for the "operations" table.
	OperationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"IPDNSBL"}},
		{Name: "ip_address", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"WAITING", "IN_PROGRESS", "DONE", "ERROR"}},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "done_at", Type: field.TypeTime, Nullable: true},
	}
	// OperationsTable holds the schema information for the "operations" table.
	OperationsTable = &schema.Table{
		Name:        "operations",
		Columns:     OperationsColumns,
		PrimaryKey:  []*schema.Column{OperationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "operation_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{OperationsColumns[2], OperationsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeBytes},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_updated_at_created_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[2], UsersColumns[1]},
			},
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DnsblQueriesTable,
		DnsblResponsesTable,
		IpsTable,
		OperationsTable,
		UsersTable,
	}
)

func init() {
	DnsblQueriesTable.ForeignKeys[0].RefTable = IpsTable
	DnsblResponsesTable.ForeignKeys[0].RefTable = DnsblQueriesTable
}
