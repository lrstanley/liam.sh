package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var (
	reUserLogin = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{0,38}$`)
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Unique().Positive().Immutable(),
		field.String("login").Unique().Match(reUserLogin).Annotations(),
		field.String("name").Optional().MaxLen(400),
		field.String("avatar_url").Optional().MaxLen(2048),
		field.String("email").Optional().MaxLen(320),
		field.String("location").Optional().MaxLen(400),
		field.String("bio").Optional().MaxLen(400),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		PrivacyMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
