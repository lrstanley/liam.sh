// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/v63/github"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
)

// fetchReleases fetches all repository releases for the authenticated user from
// Github. It will also iterate through all pages, returning all repository releases
// in their entirety.
func fetchReleases(
	ctx context.Context, repo *github.Repository,
) (allReleases []*github.RepositoryRelease, err error) {
	var resp *github.Response

	opts := &github.ListOptions{
		PerPage: 100,
		Page:    1,
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		time.Sleep(1 * time.Second)

		var releases []*github.RepositoryRelease

		log.FromContext(ctx).WithFields(log.Fields{
			"page": opts.Page,
			"repo": repo.GetFullName(),
		}).Info("querying repository releases")
		releases, resp, err = RestClient.Repositories.ListReleases(ctx, repo.GetOwner().GetLogin(), repo.GetName(), opts)
		if err != nil {
			return nil, err
		}

		allReleases = append(allReleases, releases...)

		if resp.NextPage < 1 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allReleases, nil
}

// storeRelease stores a repository release in the database.
func storeRelease(ctx context.Context, db *ent.Client, repoID int, release *github.RepositoryRelease) (id int, err error) {
	// Upsert release.
	return db.GithubRelease.Create().
		SetRepositoryID(repoID).
		SetReleaseID(release.GetID()).
		SetHTMLURL(release.GetHTMLURL()).
		SetTagName(release.GetTagName()).
		SetTargetCommitish(release.GetTargetCommitish()).
		SetName(release.GetName()).
		SetDraft(release.GetDraft()).
		SetPrerelease(release.GetPrerelease()).
		SetCreatedAt(release.GetCreatedAt().Time).
		SetPublishedAt(release.GetPublishedAt().Time).
		SetAuthor(release.GetAuthor()).
		OnConflictColumns(githubrelease.FieldReleaseID).
		UpdateNewValues().ID(ctx)
}

func storeAsset(ctx context.Context, db *ent.Client, releaseID int, asset *github.ReleaseAsset) (id int, err error) {
	// Upsert asset.
	return db.GithubAsset.Create().
		SetReleaseID(releaseID).
		SetAssetID(asset.GetID()).
		SetBrowserDownloadURL(asset.GetBrowserDownloadURL()).
		SetName(asset.GetName()).
		SetLabel(asset.GetLabel()).
		SetState(asset.GetState()).
		SetContentType(asset.GetContentType()).
		SetSize(int64(asset.GetSize())).
		SetDownloadCount(int64(asset.GetDownloadCount())).
		SetCreatedAt(asset.GetCreatedAt().Time).
		SetUpdatedAt(asset.GetUpdatedAt().Time).
		SetUploader(asset.GetUploader()).
		OnConflictColumns(githubasset.FieldAssetID).
		UpdateNewValues().ID(ctx)
}
