// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"regexp"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

var rePostSlug = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{4,50}$`)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Match(rePostSlug).Unique().Annotations(
			entgql.OrderField("SLUG"),
		),
		field.String("title").MaxLen(100).Annotations(
			entgql.OrderField("TITLE"),
		),
		field.String("content").NotEmpty(),
		field.String("content_html").NotEmpty().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		field.String("summary").NotEmpty().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		field.Time("published_at").Default(time.Now).Annotations(
			entgql.OrderField("DATE"),
		),
		field.Int("view_count").Default(0).NonNegative().Annotations(
			entgql.OrderField("VIEW_COUNT"),
			entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
		),
		field.Bool("public").Default(false),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Post) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			AllowRoles([]string{"admin"}, true),
		},
		Query: privacy.QueryPolicy{
			AllowPublicUnlessRole([]string{"admin"}),
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("posts").Unique().Required().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		edge.From("labels", Label.Type).Ref("posts").Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
