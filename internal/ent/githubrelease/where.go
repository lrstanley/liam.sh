// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package githubrelease

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldID, id))
}

// ReleaseID applies equality check predicate on the "release_id" field. It's identical to ReleaseIDEQ.
func ReleaseID(v int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldReleaseID, v))
}

// HTMLURL applies equality check predicate on the "html_url" field. It's identical to HTMLURLEQ.
func HTMLURL(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldHTMLURL, v))
}

// TagName applies equality check predicate on the "tag_name" field. It's identical to TagNameEQ.
func TagName(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldTagName, v))
}

// TargetCommitish applies equality check predicate on the "target_commitish" field. It's identical to TargetCommitishEQ.
func TargetCommitish(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldTargetCommitish, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldName, v))
}

// Draft applies equality check predicate on the "draft" field. It's identical to DraftEQ.
func Draft(v bool) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldDraft, v))
}

// Prerelease applies equality check predicate on the "prerelease" field. It's identical to PrereleaseEQ.
func Prerelease(v bool) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldPrerelease, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldCreatedAt, v))
}

// PublishedAt applies equality check predicate on the "published_at" field. It's identical to PublishedAtEQ.
func PublishedAt(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldPublishedAt, v))
}

// ReleaseIDEQ applies the EQ predicate on the "release_id" field.
func ReleaseIDEQ(v int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldReleaseID, v))
}

// ReleaseIDNEQ applies the NEQ predicate on the "release_id" field.
func ReleaseIDNEQ(v int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldReleaseID, v))
}

// ReleaseIDIn applies the In predicate on the "release_id" field.
func ReleaseIDIn(vs ...int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldReleaseID, vs...))
}

// ReleaseIDNotIn applies the NotIn predicate on the "release_id" field.
func ReleaseIDNotIn(vs ...int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldReleaseID, vs...))
}

// ReleaseIDGT applies the GT predicate on the "release_id" field.
func ReleaseIDGT(v int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldReleaseID, v))
}

// ReleaseIDGTE applies the GTE predicate on the "release_id" field.
func ReleaseIDGTE(v int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldReleaseID, v))
}

// ReleaseIDLT applies the LT predicate on the "release_id" field.
func ReleaseIDLT(v int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldReleaseID, v))
}

// ReleaseIDLTE applies the LTE predicate on the "release_id" field.
func ReleaseIDLTE(v int64) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldReleaseID, v))
}

// HTMLURLEQ applies the EQ predicate on the "html_url" field.
func HTMLURLEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldHTMLURL, v))
}

// HTMLURLNEQ applies the NEQ predicate on the "html_url" field.
func HTMLURLNEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldHTMLURL, v))
}

// HTMLURLIn applies the In predicate on the "html_url" field.
func HTMLURLIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldHTMLURL, vs...))
}

// HTMLURLNotIn applies the NotIn predicate on the "html_url" field.
func HTMLURLNotIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldHTMLURL, vs...))
}

// HTMLURLGT applies the GT predicate on the "html_url" field.
func HTMLURLGT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldHTMLURL, v))
}

// HTMLURLGTE applies the GTE predicate on the "html_url" field.
func HTMLURLGTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldHTMLURL, v))
}

// HTMLURLLT applies the LT predicate on the "html_url" field.
func HTMLURLLT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldHTMLURL, v))
}

// HTMLURLLTE applies the LTE predicate on the "html_url" field.
func HTMLURLLTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldHTMLURL, v))
}

// HTMLURLContains applies the Contains predicate on the "html_url" field.
func HTMLURLContains(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContains(FieldHTMLURL, v))
}

// HTMLURLHasPrefix applies the HasPrefix predicate on the "html_url" field.
func HTMLURLHasPrefix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasPrefix(FieldHTMLURL, v))
}

// HTMLURLHasSuffix applies the HasSuffix predicate on the "html_url" field.
func HTMLURLHasSuffix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasSuffix(FieldHTMLURL, v))
}

// HTMLURLEqualFold applies the EqualFold predicate on the "html_url" field.
func HTMLURLEqualFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEqualFold(FieldHTMLURL, v))
}

// HTMLURLContainsFold applies the ContainsFold predicate on the "html_url" field.
func HTMLURLContainsFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContainsFold(FieldHTMLURL, v))
}

// TagNameEQ applies the EQ predicate on the "tag_name" field.
func TagNameEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldTagName, v))
}

// TagNameNEQ applies the NEQ predicate on the "tag_name" field.
func TagNameNEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldTagName, v))
}

// TagNameIn applies the In predicate on the "tag_name" field.
func TagNameIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldTagName, vs...))
}

// TagNameNotIn applies the NotIn predicate on the "tag_name" field.
func TagNameNotIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldTagName, vs...))
}

// TagNameGT applies the GT predicate on the "tag_name" field.
func TagNameGT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldTagName, v))
}

// TagNameGTE applies the GTE predicate on the "tag_name" field.
func TagNameGTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldTagName, v))
}

// TagNameLT applies the LT predicate on the "tag_name" field.
func TagNameLT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldTagName, v))
}

// TagNameLTE applies the LTE predicate on the "tag_name" field.
func TagNameLTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldTagName, v))
}

// TagNameContains applies the Contains predicate on the "tag_name" field.
func TagNameContains(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContains(FieldTagName, v))
}

// TagNameHasPrefix applies the HasPrefix predicate on the "tag_name" field.
func TagNameHasPrefix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasPrefix(FieldTagName, v))
}

// TagNameHasSuffix applies the HasSuffix predicate on the "tag_name" field.
func TagNameHasSuffix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasSuffix(FieldTagName, v))
}

// TagNameEqualFold applies the EqualFold predicate on the "tag_name" field.
func TagNameEqualFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEqualFold(FieldTagName, v))
}

// TagNameContainsFold applies the ContainsFold predicate on the "tag_name" field.
func TagNameContainsFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContainsFold(FieldTagName, v))
}

// TargetCommitishEQ applies the EQ predicate on the "target_commitish" field.
func TargetCommitishEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldTargetCommitish, v))
}

// TargetCommitishNEQ applies the NEQ predicate on the "target_commitish" field.
func TargetCommitishNEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldTargetCommitish, v))
}

// TargetCommitishIn applies the In predicate on the "target_commitish" field.
func TargetCommitishIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldTargetCommitish, vs...))
}

// TargetCommitishNotIn applies the NotIn predicate on the "target_commitish" field.
func TargetCommitishNotIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldTargetCommitish, vs...))
}

// TargetCommitishGT applies the GT predicate on the "target_commitish" field.
func TargetCommitishGT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldTargetCommitish, v))
}

// TargetCommitishGTE applies the GTE predicate on the "target_commitish" field.
func TargetCommitishGTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldTargetCommitish, v))
}

// TargetCommitishLT applies the LT predicate on the "target_commitish" field.
func TargetCommitishLT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldTargetCommitish, v))
}

// TargetCommitishLTE applies the LTE predicate on the "target_commitish" field.
func TargetCommitishLTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldTargetCommitish, v))
}

// TargetCommitishContains applies the Contains predicate on the "target_commitish" field.
func TargetCommitishContains(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContains(FieldTargetCommitish, v))
}

// TargetCommitishHasPrefix applies the HasPrefix predicate on the "target_commitish" field.
func TargetCommitishHasPrefix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasPrefix(FieldTargetCommitish, v))
}

// TargetCommitishHasSuffix applies the HasSuffix predicate on the "target_commitish" field.
func TargetCommitishHasSuffix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasSuffix(FieldTargetCommitish, v))
}

// TargetCommitishEqualFold applies the EqualFold predicate on the "target_commitish" field.
func TargetCommitishEqualFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEqualFold(FieldTargetCommitish, v))
}

// TargetCommitishContainsFold applies the ContainsFold predicate on the "target_commitish" field.
func TargetCommitishContainsFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContainsFold(FieldTargetCommitish, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldContainsFold(FieldName, v))
}

// DraftEQ applies the EQ predicate on the "draft" field.
func DraftEQ(v bool) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldDraft, v))
}

// DraftNEQ applies the NEQ predicate on the "draft" field.
func DraftNEQ(v bool) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldDraft, v))
}

// PrereleaseEQ applies the EQ predicate on the "prerelease" field.
func PrereleaseEQ(v bool) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldPrerelease, v))
}

// PrereleaseNEQ applies the NEQ predicate on the "prerelease" field.
func PrereleaseNEQ(v bool) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldPrerelease, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldCreatedAt, v))
}

// PublishedAtEQ applies the EQ predicate on the "published_at" field.
func PublishedAtEQ(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldEQ(FieldPublishedAt, v))
}

// PublishedAtNEQ applies the NEQ predicate on the "published_at" field.
func PublishedAtNEQ(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNEQ(FieldPublishedAt, v))
}

// PublishedAtIn applies the In predicate on the "published_at" field.
func PublishedAtIn(vs ...time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldIn(FieldPublishedAt, vs...))
}

// PublishedAtNotIn applies the NotIn predicate on the "published_at" field.
func PublishedAtNotIn(vs ...time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldNotIn(FieldPublishedAt, vs...))
}

// PublishedAtGT applies the GT predicate on the "published_at" field.
func PublishedAtGT(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGT(FieldPublishedAt, v))
}

// PublishedAtGTE applies the GTE predicate on the "published_at" field.
func PublishedAtGTE(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldGTE(FieldPublishedAt, v))
}

// PublishedAtLT applies the LT predicate on the "published_at" field.
func PublishedAtLT(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLT(FieldPublishedAt, v))
}

// PublishedAtLTE applies the LTE predicate on the "published_at" field.
func PublishedAtLTE(v time.Time) predicate.GithubRelease {
	return predicate.GithubRelease(sql.FieldLTE(FieldPublishedAt, v))
}

// HasRepository applies the HasEdge predicate on the "repository" edge.
func HasRepository() predicate.GithubRelease {
	return predicate.GithubRelease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepositoryTable, RepositoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepositoryWith applies the HasEdge predicate on the "repository" edge with a given conditions (other predicates).
func HasRepositoryWith(preds ...predicate.GithubRepository) predicate.GithubRelease {
	return predicate.GithubRelease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepositoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepositoryTable, RepositoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssets applies the HasEdge predicate on the "assets" edge.
func HasAssets() predicate.GithubRelease {
	return predicate.GithubRelease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetsTable, AssetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetsWith applies the HasEdge predicate on the "assets" edge with a given conditions (other predicates).
func HasAssetsWith(preds ...predicate.GithubAsset) predicate.GithubRelease {
	return predicate.GithubRelease(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetsTable, AssetsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GithubRelease) predicate.GithubRelease {
	return predicate.GithubRelease(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GithubRelease) predicate.GithubRelease {
	return predicate.GithubRelease(func(s *sql.Selector) {
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
func Not(p predicate.GithubRelease) predicate.GithubRelease {
	return predicate.GithubRelease(func(s *sql.Selector) {
		p(s.Not())
	})
}
