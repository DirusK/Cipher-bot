package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

func (Request) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("active", "done", "expired").
			Default("active"),
		field.Enum("type").
			Values("encryption", "decryption"),
		field.Enum("algorithm").
			Values("AES", "RC4").
			Optional().
			Nillable(),
		field.Enum("key_mode").
			Values("auto", "manual").
			Optional().
			Nillable(),
		field.Bool("manual_key_validation").
			Optional().
			Nillable(),
		field.Int("user_id"),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("requests").
			Field("user_id").
			Unique().
			Required(),
	}
}
