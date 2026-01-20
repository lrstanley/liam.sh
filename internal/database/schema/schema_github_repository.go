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

type GithubRepository struct {
	ent.Schema
}

func (GithubRepository) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").
			Unique().
			Positive().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The ID of the repository."),
		field.String("name").
			NotEmpty().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("entrest"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The name of the repository."),
		field.String("full_name").
			NotEmpty().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("lrstanley/entrest"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The full name of the repository, which includes the owner."),
		field.String("owner_login").
			NotEmpty().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("lrstanley"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The username which owns the repository (user or organization)."),
		field.JSON("owner", &github.User{}).
			Annotations(
				entrest.WithSchema(&ogen.Schema{Ref: "#/components/schemas/GithubUser"}),
			).
			Comment("The owner data of the repository."),
		field.Bool("public").
			Default(false).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the repository is public or not."),
		field.String("html_url").
			NotEmpty().
			Annotations(
				entrest.WithExample("https://github.com/lrstanley/entrest"),
			).
			Comment("The URL of the repository."),
		field.String("description").
			Optional().
			Annotations(
				entrest.WithFilter(entrest.FilterGroupContains),
			).
			Comment("The description of the repository."),
		field.Bool("fork").
			Default(false).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the repository is a fork or not."),
		field.String("homepage").
			Optional().
			Annotations(
				entrest.WithExample("https://example.com"),
			).
			Comment("The homepage of the repository."),
		field.Int("star_count").
			NonNegative().
			Default(0).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupLength),
			).
			Comment("The number of stars the repository has."),
		field.String("default_branch").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("master"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The default branch of the repository."),
		field.Bool("is_template").
			Default(false).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the repository is a template repo or not."),
		field.Bool("has_issues").
			Default(true).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the repository has issues enabled or not."),
		field.Bool("archived").
			Default(false).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the repository is archived or not."),
		field.Time("pushed_at").
			Optional().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupArray),
			).
			Comment("The date the repository was last pushed to."),
		field.Time("created_at").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the repository was created."),
		field.Time("updated_at").
			Optional().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the repository was last updated."),
		field.JSON("license", &github.License{}).
			Optional().
			Annotations(
				entrest.WithSchema(entrest.SchemaObjectAny),
			),
	}
}

func (GithubRepository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("labels", Label.Type).
			Ref("github_repositories").
			Annotations(
				entrest.WithEagerLoad(true),
				entrest.WithFilter(entrest.FilterEdge),
			),
		edge.To("releases", GithubRelease.Type).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
				entrest.WithFilter(entrest.FilterEdge),
			),
	}
}

func (GithubRepository) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.WithIncludeOperations(entrest.OperationList, entrest.OperationRead),
	}
}
