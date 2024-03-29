// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package githubrepository

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the githubrepository type in the database.
	Label = "github_repository"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldOwnerLogin holds the string denoting the owner_login field in the database.
	FieldOwnerLogin = "owner_login"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldHTMLURL holds the string denoting the html_url field in the database.
	FieldHTMLURL = "html_url"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFork holds the string denoting the fork field in the database.
	FieldFork = "fork"
	// FieldHomepage holds the string denoting the homepage field in the database.
	FieldHomepage = "homepage"
	// FieldStarCount holds the string denoting the star_count field in the database.
	FieldStarCount = "star_count"
	// FieldDefaultBranch holds the string denoting the default_branch field in the database.
	FieldDefaultBranch = "default_branch"
	// FieldIsTemplate holds the string denoting the is_template field in the database.
	FieldIsTemplate = "is_template"
	// FieldHasIssues holds the string denoting the has_issues field in the database.
	FieldHasIssues = "has_issues"
	// FieldArchived holds the string denoting the archived field in the database.
	FieldArchived = "archived"
	// FieldPushedAt holds the string denoting the pushed_at field in the database.
	FieldPushedAt = "pushed_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLicense holds the string denoting the license field in the database.
	FieldLicense = "license"
	// EdgeLabels holds the string denoting the labels edge name in mutations.
	EdgeLabels = "labels"
	// EdgeReleases holds the string denoting the releases edge name in mutations.
	EdgeReleases = "releases"
	// Table holds the table name of the githubrepository in the database.
	Table = "github_repositories"
	// LabelsTable is the table that holds the labels relation/edge. The primary key declared below.
	LabelsTable = "label_github_repositories"
	// LabelsInverseTable is the table name for the Label entity.
	// It exists in this package in order to avoid circular dependency with the "label" package.
	LabelsInverseTable = "labels"
	// ReleasesTable is the table that holds the releases relation/edge.
	ReleasesTable = "github_releases"
	// ReleasesInverseTable is the table name for the GithubRelease entity.
	// It exists in this package in order to avoid circular dependency with the "githubrelease" package.
	ReleasesInverseTable = "github_releases"
	// ReleasesColumn is the table column denoting the releases relation/edge.
	ReleasesColumn = "github_repository_releases"
)

// Columns holds all SQL columns for githubrepository fields.
var Columns = []string{
	FieldID,
	FieldRepoID,
	FieldName,
	FieldFullName,
	FieldOwnerLogin,
	FieldOwner,
	FieldPublic,
	FieldHTMLURL,
	FieldDescription,
	FieldFork,
	FieldHomepage,
	FieldStarCount,
	FieldDefaultBranch,
	FieldIsTemplate,
	FieldHasIssues,
	FieldArchived,
	FieldPushedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLicense,
}

var (
	// LabelsPrimaryKey and LabelsColumn2 are the table columns denoting the
	// primary key for the labels relation (M2M).
	LabelsPrimaryKey = []string{"label_id", "github_repository_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/lrstanley/liam.sh/internal/database/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// RepoIDValidator is a validator for the "repo_id" field. It is called by the builders before save.
	RepoIDValidator func(int64) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	FullNameValidator func(string) error
	// OwnerLoginValidator is a validator for the "owner_login" field. It is called by the builders before save.
	OwnerLoginValidator func(string) error
	// DefaultPublic holds the default value on creation for the "public" field.
	DefaultPublic bool
	// HTMLURLValidator is a validator for the "html_url" field. It is called by the builders before save.
	HTMLURLValidator func(string) error
	// DefaultFork holds the default value on creation for the "fork" field.
	DefaultFork bool
	// DefaultStarCount holds the default value on creation for the "star_count" field.
	DefaultStarCount int
	// StarCountValidator is a validator for the "star_count" field. It is called by the builders before save.
	StarCountValidator func(int) error
	// DefaultIsTemplate holds the default value on creation for the "is_template" field.
	DefaultIsTemplate bool
	// DefaultHasIssues holds the default value on creation for the "has_issues" field.
	DefaultHasIssues bool
	// DefaultArchived holds the default value on creation for the "archived" field.
	DefaultArchived bool
)

// OrderOption defines the ordering options for the GithubRepository queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRepoID orders the results by the repo_id field.
func ByRepoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRepoID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByOwnerLogin orders the results by the owner_login field.
func ByOwnerLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerLogin, opts...).ToFunc()
}

// ByPublic orders the results by the public field.
func ByPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublic, opts...).ToFunc()
}

// ByHTMLURL orders the results by the html_url field.
func ByHTMLURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTMLURL, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByFork orders the results by the fork field.
func ByFork(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFork, opts...).ToFunc()
}

// ByHomepage orders the results by the homepage field.
func ByHomepage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHomepage, opts...).ToFunc()
}

// ByStarCount orders the results by the star_count field.
func ByStarCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStarCount, opts...).ToFunc()
}

// ByDefaultBranch orders the results by the default_branch field.
func ByDefaultBranch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultBranch, opts...).ToFunc()
}

// ByIsTemplate orders the results by the is_template field.
func ByIsTemplate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsTemplate, opts...).ToFunc()
}

// ByHasIssues orders the results by the has_issues field.
func ByHasIssues(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasIssues, opts...).ToFunc()
}

// ByArchived orders the results by the archived field.
func ByArchived(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArchived, opts...).ToFunc()
}

// ByPushedAt orders the results by the pushed_at field.
func ByPushedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPushedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLabelsCount orders the results by labels count.
func ByLabelsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLabelsStep(), opts...)
	}
}

// ByLabels orders the results by labels terms.
func ByLabels(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLabelsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReleasesCount orders the results by releases count.
func ByReleasesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReleasesStep(), opts...)
	}
}

// ByReleases orders the results by releases terms.
func ByReleases(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReleasesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLabelsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LabelsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, LabelsTable, LabelsPrimaryKey...),
	)
}
func newReleasesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReleasesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReleasesTable, ReleasesColumn),
	)
}
