// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"context"
	"regexp"
	"strings"
	"sync"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/apex/log"
	"github.com/lrstanley/chix"
	gen "github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/hook"
	"github.com/lrstanley/liam.sh/internal/database/ent/intercept"
	"github.com/lrstanley/liam.sh/internal/database/ent/post"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
	"github.com/lrstanley/liam.sh/internal/markdown"
)

const (
	summaryLen             = 40
	postViewCountThreshold = 5 * time.Minute
)

var (
	rePostSlug = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{4,50}$`)
	postViewMu = &sync.Mutex{}
	postView   = map[int]map[string]time.Time{}
)

func postViewCounter(ctx context.Context, p *gen.Post) {
	ip := chix.GetContextIP(ctx)
	if ip == nil {
		return
	}

	nctx, cancel := context.WithTimeout(
		privacy.DecisionContext(context.Background(), privacy.Allow),
		time.Second*5,
	)
	defer cancel()

	postViewMu.Lock()
	if _, ok := postView[p.ID]; !ok {
		postView[p.ID] = map[string]time.Time{}
	}

	if t, ok := postView[p.ID][ip.String()]; ok {
		if time.Since(t) < postViewCountThreshold {
			postViewMu.Unlock()
			return
		}
	}
	postView[p.ID][ip.String()] = time.Now()
	postViewMu.Unlock()

	if _, err := p.Update().AddViewCount(1).Save(nctx); err != nil {
		log.FromContext(ctx).WithError(err).WithField("post_id", p.ID).
			Error("failed to update post view count")
	}
}

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Match(rePostSlug).Unique().Annotations(
			entgql.OrderField("SLUG"),
		),
		field.String("title").MaxLen(100).Annotations(
			entgql.OrderField("TITLE"),
		),
		field.String("content").NotEmpty(),
		field.String("content_html").NotEmpty().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		field.String("summary").NotEmpty().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		field.Time("published_at").Default(time.Now).Annotations(
			entgql.OrderField("DATE"),
		),
		field.Int("view_count").Default(0).NonNegative().Annotations(
			entgql.OrderField("VIEW_COUNT"),
			entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
		),
		field.Bool("public").Default(false),
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
			AllowRoles([]string{"admin"}, true),
		},
		Query: privacy.QueryPolicy{
			AllowPublicUnlessRole([]string{"admin"}),
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("posts").Unique().Required().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		edge.From("labels", Label.Type).Ref("posts").Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Post) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return intercept.PostFunc(func(ctx context.Context, q *gen.PostQuery) (ent.Value, error) {
				v, err := next.Query(ctx, q)
				if err != nil {
					return nil, err
				}

				switch v := v.(type) {
				case []*gen.Post:
					if len(v) == 1 {
						go postViewCounter(ctx, v[0])
					}
				case *gen.Post:
					go postViewCounter(ctx, v)
				}

				return v, nil
			})
		}),
	}
}

func (Post) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			hook.If(func(next ent.Mutator) ent.Mutator {
				return hook.PostFunc(func(ctx context.Context, m *gen.PostMutation) (ent.Value, error) {
					content, ok := m.Content()
					if !ok {
						return next.Mutate(ctx, m)
					}

					md, err := markdown.Generate(ctx, content)
					if err != nil {
						return m, err
					}

					m.SetContentHTML(md)

					summary := strings.Fields(markdown.Sanitize(md))
					if len(summary) > summaryLen {
						m.SetSummary(strings.Join(summary[:summaryLen], " ") + "...")
					} else {
						m.SetSummary(strings.Join(summary, " ") + "...")
					}

					return next.Mutate(ctx, m)
				})
			}, hook.HasFields(post.FieldContent)),
			ent.OpCreate|ent.OpUpdateOne|ent.OpUpdate,
		),
	}
}
