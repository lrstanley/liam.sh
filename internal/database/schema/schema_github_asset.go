// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v63/github"
	"github.com/lrstanley/entrest"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
	"github.com/ogen-go/ogen"
)

type GithubAsset struct {
	ent.Schema
}

func (GithubAsset) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("asset_id").
			Unique().
			Positive().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The ID of the asset."),
		field.String("browser_download_url").
			NotEmpty().
			Annotations(
				entrest.WithExample("https://github.com/lrstanley/entrest/releases/download/v0.1.0/entrest-v0.1.0-linux-amd64.tar.gz"),
			).
			Comment("The URL of the asset."),
		field.String("name").
			NotEmpty().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The name of the asset."),
		field.String("label").
			Optional().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The label of the asset."),
		field.String("state").
			Optional().
			Comment("The state of the asset."),
		field.String("content_type").
			Comment("The content type of the asset."),
		field.Int64("size").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupLength),
			).
			Comment("The size of the asset in bytes."),
		field.Int64("download_count").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupLength),
			),
		field.Time("created_at").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the asset was created."),
		field.Time("updated_at").
			Optional().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the asset was last updated."),
		field.JSON("uploader", &github.User{}).
			Annotations(
				entrest.WithSchema(&ogen.Schema{Ref: "#/components/schemas/GithubUser"}),
			).
			Comment("The data of the user that uploaded the asset."),
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
		edge.From("release", GithubRelease.Type).
			Ref("assets").
			Required().
			Unique().
			Annotations(
				entrest.WithFilter(entrest.FilterEdge),
			),
	}
}

func (GithubAsset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.WithIncludeOperations(entrest.OperationList, entrest.OperationRead),
	}
}
