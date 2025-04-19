// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/entrest"
)

var reUserLogin = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{0,38}$`)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Unique().
			Positive().
			Immutable().
			Annotations(
				entrest.WithExample(12345),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupArray),
			).Comment("Users GitHub ID."),
		field.String("login").
			Unique().
			Match(reUserLogin).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("lrstanley"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("Users GitHub login ID (username)."),
		field.String("name").
			Optional().
			MaxLen(400).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("Liam Stanley"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("Users GitHub display name."),
		field.String("avatar_url").
			Optional().
			MaxLen(2048).
			Annotations(
				entrest.WithExample("https://avatars.githubusercontent.com/u/1847365?v=4"),
			).
			Comment("GitHub avatar of the user, provided by GitHub."),
		field.String("html_url").
			Optional().
			MaxLen(2048).
			Annotations(
				entrest.WithExample("https://github.com/lrstanley"),
			).
			Comment("Users GitHub profile URL."),
		field.String("email").
			Optional().
			MaxLen(320).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("Users GitHub email address."),
		field.String("location").
			Optional().
			MaxLen(400).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("Users GitHub location."),
		field.String("bio").
			Optional().
			MaxLen(400).
			Comment("Users GitHub bio."),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinTime{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.WithIncludeOperations(entrest.OperationList, entrest.OperationRead),
	}
}
