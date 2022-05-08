package schema

import (
	"regexp"
	"time"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
	"github.com/lrstanley/liam.sh/internal/policies"
)

var (
	rePostSlug = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{4,50}$`)
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Match(rePostSlug).Unique(),
		field.String("title").MaxLen(100),
		field.String("content").NotEmpty(),
		field.Time("published_at").Default(time.Now),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Post) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			policies.AllowRoles([]string{"admin"}, true),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("posts").Unique().Required().Annotations(
			entoas.Groups("post:read"),
		),
		edge.From("labels", Label.Type).Ref("posts").Annotations(
			entoas.Groups("post:read"),
		),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ReadOperation(
			entoas.OperationGroups("post:read"),
		),
	}
}
