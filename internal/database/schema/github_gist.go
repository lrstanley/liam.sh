// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/go-github/v50/github"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

type GithubGist struct {
	ent.Schema
}

func (GithubGist) Fields() []ent.Field {
	return []ent.Field{
		field.String("gist_id"),
		field.String("html_url").NotEmpty(),
		field.Bool("public"),
		field.Time("created_at").Annotations(
			entgql.OrderField("CREATED_AT"),
		),
		field.Time("updated_at").Annotations(
			entgql.OrderField("UPDATED_AT"),
		),
		field.String("description").Optional(),
		field.JSON("owner", &github.User{}).Annotations(
			entgql.Type("GithubUser"),
		),
		// Fields specific to each file.
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("type").Annotations(
			entgql.OrderField("TYPE"),
		),
		field.String("language").Optional().Annotations(
			entgql.OrderField("LANGUAGE"),
		),
		field.Int64("size").Annotations(
			entgql.OrderField("SIZE"),
		),
		field.String("raw_url"),
		field.String("content"),
	}
}

func (GithubGist) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("gist_id", "name").Unique(),
	}
}

func (GithubGist) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			AllowPublicOnly(),
			privacy.AlwaysAllowRule(),
		},
	}
}

func (GithubGist) Edges() []ent.Edge {
	return nil
}

func (GithubGist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
