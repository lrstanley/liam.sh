// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package githubasset

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AssetID applies equality check predicate on the "asset_id" field. It's identical to AssetIDEQ.
func AssetID(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetID), v))
	})
}

// BrowserDownloadURL applies equality check predicate on the "browser_download_url" field. It's identical to BrowserDownloadURLEQ.
func BrowserDownloadURL(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrowserDownloadURL), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// ContentType applies equality check predicate on the "content_type" field. It's identical to ContentTypeEQ.
func ContentType(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentType), v))
	})
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// DownloadCount applies equality check predicate on the "download_count" field. It's identical to DownloadCountEQ.
func DownloadCount(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDownloadCount), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// AssetIDEQ applies the EQ predicate on the "asset_id" field.
func AssetIDEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetID), v))
	})
}

// AssetIDNEQ applies the NEQ predicate on the "asset_id" field.
func AssetIDNEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAssetID), v))
	})
}

// AssetIDIn applies the In predicate on the "asset_id" field.
func AssetIDIn(vs ...int64) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAssetID), v...))
	})
}

// AssetIDNotIn applies the NotIn predicate on the "asset_id" field.
func AssetIDNotIn(vs ...int64) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAssetID), v...))
	})
}

// AssetIDGT applies the GT predicate on the "asset_id" field.
func AssetIDGT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAssetID), v))
	})
}

// AssetIDGTE applies the GTE predicate on the "asset_id" field.
func AssetIDGTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAssetID), v))
	})
}

// AssetIDLT applies the LT predicate on the "asset_id" field.
func AssetIDLT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAssetID), v))
	})
}

// AssetIDLTE applies the LTE predicate on the "asset_id" field.
func AssetIDLTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAssetID), v))
	})
}

// BrowserDownloadURLEQ applies the EQ predicate on the "browser_download_url" field.
func BrowserDownloadURLEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLNEQ applies the NEQ predicate on the "browser_download_url" field.
func BrowserDownloadURLNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLIn applies the In predicate on the "browser_download_url" field.
func BrowserDownloadURLIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBrowserDownloadURL), v...))
	})
}

// BrowserDownloadURLNotIn applies the NotIn predicate on the "browser_download_url" field.
func BrowserDownloadURLNotIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBrowserDownloadURL), v...))
	})
}

// BrowserDownloadURLGT applies the GT predicate on the "browser_download_url" field.
func BrowserDownloadURLGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLGTE applies the GTE predicate on the "browser_download_url" field.
func BrowserDownloadURLGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLLT applies the LT predicate on the "browser_download_url" field.
func BrowserDownloadURLLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLLTE applies the LTE predicate on the "browser_download_url" field.
func BrowserDownloadURLLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLContains applies the Contains predicate on the "browser_download_url" field.
func BrowserDownloadURLContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLHasPrefix applies the HasPrefix predicate on the "browser_download_url" field.
func BrowserDownloadURLHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLHasSuffix applies the HasSuffix predicate on the "browser_download_url" field.
func BrowserDownloadURLHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLEqualFold applies the EqualFold predicate on the "browser_download_url" field.
func BrowserDownloadURLEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBrowserDownloadURL), v))
	})
}

// BrowserDownloadURLContainsFold applies the ContainsFold predicate on the "browser_download_url" field.
func BrowserDownloadURLContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBrowserDownloadURL), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLabel), v))
	})
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLabel), v))
	})
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLabel), v...))
	})
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLabel), v...))
	})
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLabel), v))
	})
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLabel), v))
	})
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLabel), v))
	})
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLabel), v))
	})
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLabel), v))
	})
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLabel), v))
	})
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLabel), v))
	})
}

// LabelIsNil applies the IsNil predicate on the "label" field.
func LabelIsNil() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLabel)))
	})
}

// LabelNotNil applies the NotNil predicate on the "label" field.
func LabelNotNil() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLabel)))
	})
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLabel), v))
	})
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLabel), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldState)))
	})
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldState)))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// ContentTypeEQ applies the EQ predicate on the "content_type" field.
func ContentTypeEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentType), v))
	})
}

// ContentTypeNEQ applies the NEQ predicate on the "content_type" field.
func ContentTypeNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContentType), v))
	})
}

// ContentTypeIn applies the In predicate on the "content_type" field.
func ContentTypeIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContentType), v...))
	})
}

// ContentTypeNotIn applies the NotIn predicate on the "content_type" field.
func ContentTypeNotIn(vs ...string) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContentType), v...))
	})
}

// ContentTypeGT applies the GT predicate on the "content_type" field.
func ContentTypeGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContentType), v))
	})
}

// ContentTypeGTE applies the GTE predicate on the "content_type" field.
func ContentTypeGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContentType), v))
	})
}

// ContentTypeLT applies the LT predicate on the "content_type" field.
func ContentTypeLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContentType), v))
	})
}

// ContentTypeLTE applies the LTE predicate on the "content_type" field.
func ContentTypeLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContentType), v))
	})
}

// ContentTypeContains applies the Contains predicate on the "content_type" field.
func ContentTypeContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContentType), v))
	})
}

// ContentTypeHasPrefix applies the HasPrefix predicate on the "content_type" field.
func ContentTypeHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContentType), v))
	})
}

// ContentTypeHasSuffix applies the HasSuffix predicate on the "content_type" field.
func ContentTypeHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContentType), v))
	})
}

// ContentTypeEqualFold applies the EqualFold predicate on the "content_type" field.
func ContentTypeEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContentType), v))
	})
}

// ContentTypeContainsFold applies the ContainsFold predicate on the "content_type" field.
func ContentTypeContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContentType), v))
	})
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSize), v))
	})
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSize), v...))
	})
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSize), v...))
	})
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSize), v))
	})
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSize), v))
	})
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSize), v))
	})
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSize), v))
	})
}

// DownloadCountEQ applies the EQ predicate on the "download_count" field.
func DownloadCountEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDownloadCount), v))
	})
}

// DownloadCountNEQ applies the NEQ predicate on the "download_count" field.
func DownloadCountNEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDownloadCount), v))
	})
}

// DownloadCountIn applies the In predicate on the "download_count" field.
func DownloadCountIn(vs ...int64) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDownloadCount), v...))
	})
}

// DownloadCountNotIn applies the NotIn predicate on the "download_count" field.
func DownloadCountNotIn(vs ...int64) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDownloadCount), v...))
	})
}

// DownloadCountGT applies the GT predicate on the "download_count" field.
func DownloadCountGT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDownloadCount), v))
	})
}

// DownloadCountGTE applies the GTE predicate on the "download_count" field.
func DownloadCountGTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDownloadCount), v))
	})
}

// DownloadCountLT applies the LT predicate on the "download_count" field.
func DownloadCountLT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDownloadCount), v))
	})
}

// DownloadCountLTE applies the LTE predicate on the "download_count" field.
func DownloadCountLTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDownloadCount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GithubAsset {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// HasRelease applies the HasEdge predicate on the "release" edge.
func HasRelease() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReleaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReleaseTable, ReleaseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReleaseWith applies the HasEdge predicate on the "release" edge with a given conditions (other predicates).
func HasReleaseWith(preds ...predicate.GithubRelease) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReleaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReleaseTable, ReleaseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GithubAsset) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GithubAsset) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
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
func Not(p predicate.GithubAsset) predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		p(s.Not())
	})
}
