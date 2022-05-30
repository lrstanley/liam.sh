package schema

import (
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

var (
	reLabel = regexp.MustCompile(`^[a-z\d][a-z\d-]*$`)
)

type Label struct {
	ent.Schema
}

func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(reLabel).Annotations(
			entgql.OrderField("NAME"),
		),
	}
}

func (Label) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
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
			entgql.RelayConnection(),
		),
		edge.To("github_repositories", GithubRepository.Type).Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Label) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
