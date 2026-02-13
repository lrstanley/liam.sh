// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v82/github"
	"github.com/lrstanley/entrest"
	"github.com/ogen-go/ogen"
)

type GithubEvent struct {
	ent.Schema
}

func (GithubEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").
			Unique().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The ID of the event."),
		field.String("event_type").
			NotEmpty().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("The type of the event."),
		field.Time("created_at").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("The date the event was created."),
		field.Bool("public").
			Default(false).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the event is public or not."),
		field.Int64("actor_id").
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The ID of the actor."),
		field.JSON("actor", &github.User{}).
			Annotations(
				entrest.WithSchema(&ogen.Schema{Ref: "#/components/schemas/GithubUser"}),
			).
			Comment("The actor data of the event."),
		field.Int64("repo_id").
			Positive().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("The ID of the repository."),
		field.JSON("repo", &github.Repository{}).
			Annotations(
				entrest.WithSchema(entrest.SchemaObjectAny),
			).
			Comment("The repository of the event."),
		field.JSON("payload", map[string]any{}).
			Annotations(
				entrest.WithSchema(entrest.SchemaObjectAny),
			).
			Comment("The payload of the event."),
	}
}

func (GithubEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.WithIncludeOperations(entrest.OperationList, entrest.OperationRead),
	}
}
