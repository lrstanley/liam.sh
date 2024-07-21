// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package rest

import (
	"github.com/lrstanley/liam.sh/internal/database/ent"
)

// EagerLoadGithubAsset eager-loads the edges of a GithubAsset entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadGithubAsset(query *ent.GithubAssetQuery) *ent.GithubAssetQuery {
	return query
}

// EagerLoadGithubEvent eager-loads the edges of a GithubEvent entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadGithubEvent(query *ent.GithubEventQuery) *ent.GithubEventQuery {
	return query
}

// EagerLoadGithubGist eager-loads the edges of a GithubGist entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadGithubGist(query *ent.GithubGistQuery) *ent.GithubGistQuery {
	return query
}

// EagerLoadGithubRelease eager-loads the edges of a GithubRelease entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadGithubRelease(query *ent.GithubReleaseQuery) *ent.GithubReleaseQuery {
	return query.WithRepository().WithAssets()
}

// EagerLoadGithubRepository eager-loads the edges of a GithubRepository entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadGithubRepository(query *ent.GithubRepositoryQuery) *ent.GithubRepositoryQuery {
	return query.WithLabels()
}

// EagerLoadLabel eager-loads the edges of a Label entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadLabel(query *ent.LabelQuery) *ent.LabelQuery {
	return query
}

// EagerLoadPost eager-loads the edges of a Post entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadPost(query *ent.PostQuery) *ent.PostQuery {
	return query.WithAuthor().WithLabels()
}

// EagerLoadUser eager-loads the edges of a User entity, if any edges
// were requested to be eager-loaded, based off associated annotations.
func EagerLoadUser(query *ent.UserQuery) *ent.UserQuery {
	return query
}
