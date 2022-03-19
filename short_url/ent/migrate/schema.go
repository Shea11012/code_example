// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TinyUrlsColumns holds the columns for the "tiny_urls" table.
	TinyUrlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "url", Type: field.TypeString},
	}
	// TinyUrlsTable holds the schema information for the "tiny_urls" table.
	TinyUrlsTable = &schema.Table{
		Name:       "tiny_urls",
		Columns:    TinyUrlsColumns,
		PrimaryKey: []*schema.Column{TinyUrlsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TinyUrlsTable,
	}
)

func init() {
	TinyUrlsTable.Annotation = &entsql.Annotation{
		Table: "tiny_urls",
	}
}
