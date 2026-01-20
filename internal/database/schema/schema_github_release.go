// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v63/github"
	"github.com/lrstanley/entrest"
	"github.com/ogen-go/ogen"
)

type GithubRelease struct {
	ent.Schema
}

func (GithubRelease) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("release_id").
			Unique().
			Positive().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The ID of the release."),
		field.String("html_url").
			NotEmpty().
			Annotations(
				entrest.WithExample("https://github.com/lrstanley/entrest/releases/tag/v0.1.0"),
			).
			Comment("The URL of the release."),
		field.String("tag_name").
			NotEmpty().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("v0.1.0"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The tag name of the release."),
		field.String("target_commitish").
			NotEmpty().
			Annotations(
				entrest.WithExample("master"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch."),
		field.String("name").
			Optional().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The name of the release."),
		field.Bool("draft").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Indicates whether the release is a draft."),
		field.Bool("prerelease").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Indicates whether the release is a prerelease."),
		field.Time("created_at").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the release was created."),
		field.Time("published_at").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the release was published."),
		field.JSON("author", &github.User{}).
			Annotations(
				entrest.WithSchema(&ogen.Schema{Ref: "#/components/schemas/GithubUser"}),
			),
	}
}

func (GithubRelease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", GithubRepository.Type).
			Ref("releases").
			Required().
			Unique().
			Annotations(
				entrest.WithEagerLoad(true),
				entrest.WithFilter(entrest.FilterEdge),
			),
		edge.To("assets", GithubAsset.Type).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
				entrest.WithEagerLoad(true),
				entrest.WithFilter(entrest.FilterEdge),
			),
	}
}

func (GithubRelease) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.WithIncludeOperations(entrest.OperationList, entrest.OperationRead),
	}
}
