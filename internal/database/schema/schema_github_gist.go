// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/go-github/v82/github"
	"github.com/lrstanley/entrest"
	"github.com/ogen-go/ogen"
)

type GithubGist struct {
	ent.Schema
}

func (GithubGist) Fields() []ent.Field {
	return []ent.Field{
		field.String("gist_id").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The ID of the gist."),
		field.String("html_url").
			NotEmpty().
			Annotations(
				entrest.WithExample("https://gist.github.com/lrstanley/c4f0a3f2b8a2f3a4c2c0"),
			).
			Comment("The URL of the gist."),
		field.Bool("public").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the gist is public or not."),
		field.Time("created_at").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the gist was created."),
		field.Time("updated_at").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the gist was last updated."),
		field.String("description").
			Optional().
			Annotations(
				entrest.WithFilter(entrest.FilterGroupContains),
			).
			Comment("The description of the gist."),
		field.JSON("owner", &github.User{}).
			Annotations(
				entrest.WithSchema(&ogen.Schema{Ref: "#/components/schemas/GithubUser"}),
			).
			Comment("The owner data of the gist."),
		// Fields specific to each file.
		field.String("name").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The name of the file."),
		field.String("type").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The type of the file."),
		field.String("language").
			Optional().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The programming language of the file."),
		field.Int64("size").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The size of the file in bytes."),
		field.String("raw_url").
			Annotations().
			Comment("The raw URL of the file."),
		field.String("content").
			Annotations(
				entrest.WithFilter(entrest.FilterGroupContains),
			).
			Comment("The content of the file."),
	}
}

func (GithubGist) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("gist_id", "name").Unique(),
	}
}

func (GithubGist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.WithIncludeOperations(entrest.OperationList, entrest.OperationRead),
	}
}
