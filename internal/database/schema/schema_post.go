// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/entrest"
)

const summaryLen = 40

var rePostSlug = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{4,50}$`)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			Match(rePostSlug).
			MinLen(5).
			MaxLen(50).
			Unique().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("hello-world"),
				entrest.WithFilter(entrest.FilterGroupEqual),
				entrest.WithFilterGroup("search"),
			).
			Comment("Post slug."),
		field.String("title").
			MinLen(5).
			MaxLen(100).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithExample("Hello World"),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
				entrest.WithFilterGroup("search"),
			).
			Comment("Post title."),
		field.String("content").
			NotEmpty().
			Annotations(
				entrest.WithExample("## Title\n\nHello World"),
				entrest.WithFilter(entrest.FilterGroupContains),
				entrest.WithFilterGroup("search"),
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

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("posts").
			Unique().
			Required().
			Annotations(
				entrest.WithEagerLoad(true),
				entrest.WithFilter(entrest.FilterEdge),
				entrest.WithExcludeOperations(entrest.OperationCreate, entrest.OperationUpdate),
			),
		edge.From("labels", Label.Type).
			Ref("posts").
			Annotations(
				entrest.WithEagerLoad(true),
				entrest.WithFilter(entrest.FilterEdge),
				entrest.WithEdgeUpdateBulk(true),
			),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
