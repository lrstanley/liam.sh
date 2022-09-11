// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentHTML applies equality check predicate on the "content_html" field. It's identical to ContentHTMLEQ.
func ContentHTML(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentHTML), v))
	})
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// PublishedAt applies equality check predicate on the "published_at" field. It's identical to PublishedAtEQ.
func PublishedAt(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublishedAt), v))
	})
}

// ViewCount applies equality check predicate on the "view_count" field. It's identical to ViewCountEQ.
func ViewCount(v int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldViewCount), v))
	})
}

// Public applies equality check predicate on the "public" field. It's identical to PublicEQ.
func Public(v bool) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublic), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlug), v))
	})
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSlug), v...))
	})
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSlug), v...))
	})
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlug), v))
	})
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlug), v))
	})
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlug), v))
	})
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlug), v))
	})
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlug), v))
	})
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlug), v))
	})
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlug), v))
	})
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlug), v))
	})
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlug), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// ContentHTMLEQ applies the EQ predicate on the "content_html" field.
func ContentHTMLEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLNEQ applies the NEQ predicate on the "content_html" field.
func ContentHTMLNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLIn applies the In predicate on the "content_html" field.
func ContentHTMLIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContentHTML), v...))
	})
}

// ContentHTMLNotIn applies the NotIn predicate on the "content_html" field.
func ContentHTMLNotIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContentHTML), v...))
	})
}

// ContentHTMLGT applies the GT predicate on the "content_html" field.
func ContentHTMLGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLGTE applies the GTE predicate on the "content_html" field.
func ContentHTMLGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLLT applies the LT predicate on the "content_html" field.
func ContentHTMLLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLLTE applies the LTE predicate on the "content_html" field.
func ContentHTMLLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLContains applies the Contains predicate on the "content_html" field.
func ContentHTMLContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLHasPrefix applies the HasPrefix predicate on the "content_html" field.
func ContentHTMLHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLHasSuffix applies the HasSuffix predicate on the "content_html" field.
func ContentHTMLHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLEqualFold applies the EqualFold predicate on the "content_html" field.
func ContentHTMLEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContentHTML), v))
	})
}

// ContentHTMLContainsFold applies the ContainsFold predicate on the "content_html" field.
func ContentHTMLContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContentHTML), v))
	})
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSummary), v))
	})
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSummary), v...))
	})
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSummary), v...))
	})
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSummary), v))
	})
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSummary), v))
	})
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSummary), v))
	})
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSummary), v))
	})
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSummary), v))
	})
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSummary), v))
	})
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSummary), v))
	})
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSummary), v))
	})
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSummary), v))
	})
}

// PublishedAtEQ applies the EQ predicate on the "published_at" field.
func PublishedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublishedAt), v))
	})
}

// PublishedAtNEQ applies the NEQ predicate on the "published_at" field.
func PublishedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublishedAt), v))
	})
}

// PublishedAtIn applies the In predicate on the "published_at" field.
func PublishedAtIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPublishedAt), v...))
	})
}

// PublishedAtNotIn applies the NotIn predicate on the "published_at" field.
func PublishedAtNotIn(vs ...time.Time) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPublishedAt), v...))
	})
}

// PublishedAtGT applies the GT predicate on the "published_at" field.
func PublishedAtGT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPublishedAt), v))
	})
}

// PublishedAtGTE applies the GTE predicate on the "published_at" field.
func PublishedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPublishedAt), v))
	})
}

// PublishedAtLT applies the LT predicate on the "published_at" field.
func PublishedAtLT(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPublishedAt), v))
	})
}

// PublishedAtLTE applies the LTE predicate on the "published_at" field.
func PublishedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPublishedAt), v))
	})
}

// ViewCountEQ applies the EQ predicate on the "view_count" field.
func ViewCountEQ(v int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldViewCount), v))
	})
}

// ViewCountNEQ applies the NEQ predicate on the "view_count" field.
func ViewCountNEQ(v int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldViewCount), v))
	})
}

// ViewCountIn applies the In predicate on the "view_count" field.
func ViewCountIn(vs ...int) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldViewCount), v...))
	})
}

// ViewCountNotIn applies the NotIn predicate on the "view_count" field.
func ViewCountNotIn(vs ...int) predicate.Post {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldViewCount), v...))
	})
}

// ViewCountGT applies the GT predicate on the "view_count" field.
func ViewCountGT(v int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldViewCount), v))
	})
}

// ViewCountGTE applies the GTE predicate on the "view_count" field.
func ViewCountGTE(v int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldViewCount), v))
	})
}

// ViewCountLT applies the LT predicate on the "view_count" field.
func ViewCountLT(v int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldViewCount), v))
	})
}

// ViewCountLTE applies the LTE predicate on the "view_count" field.
func ViewCountLTE(v int) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldViewCount), v))
	})
}

// PublicEQ applies the EQ predicate on the "public" field.
func PublicEQ(v bool) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublic), v))
	})
}

// PublicNEQ applies the NEQ predicate on the "public" field.
func PublicNEQ(v bool) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublic), v))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLabels applies the HasEdge predicate on the "labels" edge.
func HasLabels() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LabelsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LabelsTable, LabelsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLabelsWith applies the HasEdge predicate on the "labels" edge with a given conditions (other predicates).
func HasLabelsWith(preds ...predicate.Label) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LabelsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LabelsTable, LabelsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		p(s.Not())
	})
}
