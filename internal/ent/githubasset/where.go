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
	return predicate.GithubAsset(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldID, id))
}

// AssetID applies equality check predicate on the "asset_id" field. It's identical to AssetIDEQ.
func AssetID(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldAssetID, v))
}

// BrowserDownloadURL applies equality check predicate on the "browser_download_url" field. It's identical to BrowserDownloadURLEQ.
func BrowserDownloadURL(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldBrowserDownloadURL, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldName, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldState, v))
}

// ContentType applies equality check predicate on the "content_type" field. It's identical to ContentTypeEQ.
func ContentType(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldContentType, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldSize, v))
}

// DownloadCount applies equality check predicate on the "download_count" field. It's identical to DownloadCountEQ.
func DownloadCount(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldDownloadCount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldUpdatedAt, v))
}

// AssetIDEQ applies the EQ predicate on the "asset_id" field.
func AssetIDEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldAssetID, v))
}

// AssetIDNEQ applies the NEQ predicate on the "asset_id" field.
func AssetIDNEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldAssetID, v))
}

// AssetIDIn applies the In predicate on the "asset_id" field.
func AssetIDIn(vs ...int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldAssetID, vs...))
}

// AssetIDNotIn applies the NotIn predicate on the "asset_id" field.
func AssetIDNotIn(vs ...int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldAssetID, vs...))
}

// AssetIDGT applies the GT predicate on the "asset_id" field.
func AssetIDGT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldAssetID, v))
}

// AssetIDGTE applies the GTE predicate on the "asset_id" field.
func AssetIDGTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldAssetID, v))
}

// AssetIDLT applies the LT predicate on the "asset_id" field.
func AssetIDLT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldAssetID, v))
}

// AssetIDLTE applies the LTE predicate on the "asset_id" field.
func AssetIDLTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldAssetID, v))
}

// BrowserDownloadURLEQ applies the EQ predicate on the "browser_download_url" field.
func BrowserDownloadURLEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLNEQ applies the NEQ predicate on the "browser_download_url" field.
func BrowserDownloadURLNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLIn applies the In predicate on the "browser_download_url" field.
func BrowserDownloadURLIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldBrowserDownloadURL, vs...))
}

// BrowserDownloadURLNotIn applies the NotIn predicate on the "browser_download_url" field.
func BrowserDownloadURLNotIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldBrowserDownloadURL, vs...))
}

// BrowserDownloadURLGT applies the GT predicate on the "browser_download_url" field.
func BrowserDownloadURLGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLGTE applies the GTE predicate on the "browser_download_url" field.
func BrowserDownloadURLGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLLT applies the LT predicate on the "browser_download_url" field.
func BrowserDownloadURLLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLLTE applies the LTE predicate on the "browser_download_url" field.
func BrowserDownloadURLLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLContains applies the Contains predicate on the "browser_download_url" field.
func BrowserDownloadURLContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContains(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLHasPrefix applies the HasPrefix predicate on the "browser_download_url" field.
func BrowserDownloadURLHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasPrefix(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLHasSuffix applies the HasSuffix predicate on the "browser_download_url" field.
func BrowserDownloadURLHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasSuffix(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLEqualFold applies the EqualFold predicate on the "browser_download_url" field.
func BrowserDownloadURLEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEqualFold(FieldBrowserDownloadURL, v))
}

// BrowserDownloadURLContainsFold applies the ContainsFold predicate on the "browser_download_url" field.
func BrowserDownloadURLContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContainsFold(FieldBrowserDownloadURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContainsFold(FieldName, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelIsNil applies the IsNil predicate on the "label" field.
func LabelIsNil() predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIsNull(FieldLabel))
}

// LabelNotNil applies the NotNil predicate on the "label" field.
func LabelNotNil() predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotNull(FieldLabel))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContainsFold(FieldLabel, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasSuffix(FieldState, v))
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIsNull(FieldState))
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotNull(FieldState))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContainsFold(FieldState, v))
}

// ContentTypeEQ applies the EQ predicate on the "content_type" field.
func ContentTypeEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldContentType, v))
}

// ContentTypeNEQ applies the NEQ predicate on the "content_type" field.
func ContentTypeNEQ(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldContentType, v))
}

// ContentTypeIn applies the In predicate on the "content_type" field.
func ContentTypeIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldContentType, vs...))
}

// ContentTypeNotIn applies the NotIn predicate on the "content_type" field.
func ContentTypeNotIn(vs ...string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldContentType, vs...))
}

// ContentTypeGT applies the GT predicate on the "content_type" field.
func ContentTypeGT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldContentType, v))
}

// ContentTypeGTE applies the GTE predicate on the "content_type" field.
func ContentTypeGTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldContentType, v))
}

// ContentTypeLT applies the LT predicate on the "content_type" field.
func ContentTypeLT(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldContentType, v))
}

// ContentTypeLTE applies the LTE predicate on the "content_type" field.
func ContentTypeLTE(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldContentType, v))
}

// ContentTypeContains applies the Contains predicate on the "content_type" field.
func ContentTypeContains(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContains(FieldContentType, v))
}

// ContentTypeHasPrefix applies the HasPrefix predicate on the "content_type" field.
func ContentTypeHasPrefix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasPrefix(FieldContentType, v))
}

// ContentTypeHasSuffix applies the HasSuffix predicate on the "content_type" field.
func ContentTypeHasSuffix(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldHasSuffix(FieldContentType, v))
}

// ContentTypeEqualFold applies the EqualFold predicate on the "content_type" field.
func ContentTypeEqualFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEqualFold(FieldContentType, v))
}

// ContentTypeContainsFold applies the ContainsFold predicate on the "content_type" field.
func ContentTypeContainsFold(v string) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldContainsFold(FieldContentType, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldSize, v))
}

// DownloadCountEQ applies the EQ predicate on the "download_count" field.
func DownloadCountEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldDownloadCount, v))
}

// DownloadCountNEQ applies the NEQ predicate on the "download_count" field.
func DownloadCountNEQ(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldDownloadCount, v))
}

// DownloadCountIn applies the In predicate on the "download_count" field.
func DownloadCountIn(vs ...int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldDownloadCount, vs...))
}

// DownloadCountNotIn applies the NotIn predicate on the "download_count" field.
func DownloadCountNotIn(vs ...int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldDownloadCount, vs...))
}

// DownloadCountGT applies the GT predicate on the "download_count" field.
func DownloadCountGT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldDownloadCount, v))
}

// DownloadCountGTE applies the GTE predicate on the "download_count" field.
func DownloadCountGTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldDownloadCount, v))
}

// DownloadCountLT applies the LT predicate on the "download_count" field.
func DownloadCountLT(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldDownloadCount, v))
}

// DownloadCountLTE applies the LTE predicate on the "download_count" field.
func DownloadCountLTE(v int64) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldDownloadCount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.GithubAsset {
	return predicate.GithubAsset(sql.FieldNotNull(FieldUpdatedAt))
}

// HasRelease applies the HasEdge predicate on the "release" edge.
func HasRelease() predicate.GithubAsset {
	return predicate.GithubAsset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
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
