package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v44/github"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

type GithubRepository struct {
	ent.Schema
}

func (GithubRepository) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("repo_id").Unique().Positive(),
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("full_name").NotEmpty().Annotations(
			entgql.OrderField("FULL_NAME"),
		),
		field.String("owner_login").NotEmpty().Annotations(
			entgql.OrderField("OWNER_LOGIN"),
		),
		field.JSON("owner", &github.User{}).Annotations(
			entgql.Type("GithubUser"),
		),
		field.Bool("public").Default(false),
		field.String("html_url").NotEmpty(),
		field.String("description").Optional(),
		field.Bool("fork").Default(false),
		field.String("homepage").Optional(),
		field.Int("star_count").NonNegative().Default(0).Annotations(
			entgql.OrderField("STAR_COUNT"),
		),
		field.String("default_branch"),
		field.Bool("is_template").Default(false),
		field.Bool("has_issues").Default(true),
		field.Bool("archived").Default(false),
		field.Time("pushed_at").Optional().Annotations(
			entgql.OrderField("PUSHED_AT"),
		),
		field.Time("created_at").Annotations(
			entgql.OrderField("CREATED_AT"),
		),
		field.Time("updated_at").Optional().Annotations(
			entgql.OrderField("UPDATED_AT"),
		),
		field.JSON("license", &github.License{}).Optional().Annotations(
			entgql.Type("GithubLicense"),
		),
	}
}

func (GithubRepository) Policy() ent.Policy {
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

func (GithubRepository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("labels", Label.Type).Ref("github_repositories").Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (GithubRepository) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
