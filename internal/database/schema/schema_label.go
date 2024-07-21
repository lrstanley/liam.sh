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
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

var reLabel = regexp.MustCompile(`^[a-z\d][a-z\d-]*$`)

type Label struct {
	ent.Schema
}

func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Match(reLabel).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("golang"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("Label name."),
	}
}

func (Label) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinTime{},
	}
}

func (Label) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			AllowRoles([]string{"admin"}, true),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Label) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(
			entrest.WithFilter(entrest.FilterEdge),
		),
		edge.To("github_repositories", GithubRepository.Type).Annotations(
			entrest.WithFilter(entrest.FilterEdge),
		),
	}
}

func (Label) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
