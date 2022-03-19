package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TinyURL holds the schema definition for the TinyURL entity.
type TinyURL struct {
	ent.Schema
}

// Fields of the TinyURL.
func (TinyURL) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.String("url"),
	}
}

// Edges of the TinyURL.
func (TinyURL) Edges() []ent.Edge {
	return nil
}

func (TinyURL) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tiny_urls"},
	}
}
