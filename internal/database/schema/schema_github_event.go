// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v52/github"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

type GithubEvent struct {
	ent.Schema
}

func (GithubEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").Unique().Annotations(
			entgql.OrderField("EVENT_ID"),
		),
		field.String("event_type").NotEmpty().Annotations(
			entgql.OrderField("EVENT_TYPE"),
		),
		field.Time("created_at").Annotations(
			entgql.OrderField("CREATED_AT"),
		),
		field.Bool("public").Default(false),
		field.Int64("actor_id").Annotations(
			entgql.OrderField("ACTOR_ID"),
		),
		field.JSON("actor", &github.User{}).Annotations(entgql.Type("GithubUser")),
		field.Int64("repo_id").Positive().Annotations(
			entgql.OrderField("REPO_ID"),
		),
		field.JSON("repo", &github.Repository{}).Annotations(entgql.Type("GithubEventRepo")),
		field.JSON("payload", map[string]any{}).Annotations(entgql.Type("Map")),
	}
}

func (GithubEvent) Policy() ent.Policy {
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

func (GithubEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
