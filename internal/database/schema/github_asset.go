// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v50/github"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

type GithubAsset struct {
	ent.Schema
}

func (GithubAsset) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("asset_id").Unique().Positive(),
		field.String("browser_download_url").NotEmpty(),
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("label").Optional(),
		field.String("state").Optional(),
		field.String("content_type"),
		field.Int64("size"),
		field.Int64("download_count").Annotations(
			entgql.OrderField("DOWNLOAD_COUNT"),
		),
		field.Time("created_at").Annotations(
			entgql.OrderField("CREATED_AT"),
		),
		field.Time("updated_at").Optional().Annotations(
			entgql.OrderField("UPDATED_AT"),
		),
		field.JSON("uploader", &github.User{}).Annotations(
			entgql.Type("GithubUser"),
		),
	}
}

func (GithubAsset) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (GithubAsset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("release", GithubRelease.Type).Ref("assets").Required().Unique(),
	}
}

func (GithubAsset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
