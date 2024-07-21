// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/entrest"
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
		field.String("slug").
			Match(rePostSlug).
			Unique().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("hello-world"),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Post slug."),
		field.String("title").
			MaxLen(100).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("Hello World"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			).
			Comment("Post title."),
		field.String("content").
			NotEmpty().
			Annotations(
				entrest.WithExample("## Title\n\nHello World"),
				entrest.WithFilter(entrest.FilterGroupContains),
			).
			Comment("Post content in Markdown."),
		field.String("content_html").
			NotEmpty().
			Annotations(
				entrest.WithReadOnly(true),
				entrest.WithExample("<h1>Title</h1>\n\n<p>Hello World</p>"),
				entrest.WithFilter(entrest.FilterGroupContains),
			).
			Comment("Generated HTML content (produced from 'content' field)."),
		field.String("summary").
			NotEmpty().
			Annotations(
				entrest.WithReadOnly(true),
				entrest.WithExample("Some example content here..."),
			).
			Comment("Post summary, which is produced from the first sentence or two of the post content."),
		field.Time("published_at").
			Default(time.Now).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			),
		field.Int("view_count").
			Default(0).
			NonNegative().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithReadOnly(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("Number of times the post has been viewed."),
		field.Bool("public").
			Default(false).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact),
			).
			Comment("Whether the post is public or not."),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinTime{},
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
		edge.From("author", User.Type).
			Ref("posts").
			Unique().
			Required().
			Annotations(
				entrest.WithEagerLoad(true),
				entrest.WithFilter(entrest.FilterEdge),
			),
		edge.From("labels", Label.Type).
			Ref("posts").
			Annotations(
				entrest.WithEagerLoad(true),
				entrest.WithFilter(entrest.FilterEdge),
			),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{}
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
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.PostFunc(func(ctx context.Context, m *gen.PostMutation) (ent.Value, error) {
				ident := chix.IdentFromContext[gen.User](ctx)
				if ident == nil {
					return nil, errors.New("unauthenticated")
				}
				m.SetAuthorID(ident.ID)
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
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
