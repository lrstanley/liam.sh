// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v50/github"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

type GithubRelease struct {
	ent.Schema
}

func (GithubRelease) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("release_id").Unique().Positive(),
		field.String("html_url").NotEmpty(),
		field.String("tag_name").NotEmpty().Annotations(
			entgql.OrderField("TAG_NAME"),
		),
		field.String("target_commitish").NotEmpty(),
		field.String("name").Optional().Annotations(
			entgql.OrderField("NAME"),
		),
		field.Bool("draft"),
		field.Bool("prerelease"),
		field.Time("created_at").Annotations(
			entgql.OrderField("CREATED_AT"),
		),
		field.Time("published_at").Annotations(
			entgql.OrderField("PUBLISHED_AT"),
		),
		field.JSON("author", &github.User{}).Annotations(
			entgql.Type("GithubUser"),
		),
	}
}

func (GithubRelease) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (GithubRelease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", GithubRepository.Type).Ref("releases").Required().Unique(),
		edge.To("assets", GithubAsset.Type).Annotations(
			entgql.RelayConnection(),
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
		),
	}
}

func (GithubRelease) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
