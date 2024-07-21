// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"net/http"
	"sort"

	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/database/ent/rest"
)

type OutdatedRepositoryRelease struct {
	Repository *ent.GithubRepository `json:"repository"`
	Release    *ent.GithubRelease    `json:"release"`
}

func (h *handler) getOutdatedGithubReleases(w http.ResponseWriter, r *http.Request) {
	// Find all github repos that have releases, sorted by the oldest release.

	repositories, err := rest.EagerLoadGithubRepository(h.db.GithubRepository.Query().Where(
		githubrepository.Public(true),
		githubrepository.Archived(false),
		githubrepository.IsTemplate(false),
		githubrepository.HasReleases(),
	)).All(r.Context())
	if err != nil {
		chix.Error(w, r, err)
		return
	}

	var results []OutdatedRepositoryRelease
	var release *ent.GithubRelease

	// Get the oldest release for each repository.
	for _, repository := range repositories {
		release, err = rest.EagerLoadGithubRelease(h.db.GithubRelease.Query().Where(
			githubrelease.HasRepositoryWith(githubrepository.ID(repository.ID)),
		)).Order(ent.Desc(githubrelease.FieldCreatedAt)).First(r.Context())
		if err != nil {
			chix.Error(w, r, err)
			return
		}

		results = append(results, OutdatedRepositoryRelease{
			Repository: repository,
			Release:    release,
		})
	}

	// Sort by oldest release first.
	sort.Slice(results, func(i, j int) bool {
		return results[i].Release.CreatedAt.Before(results[j].Release.CreatedAt)
	})

	rest.JSON(w, r, http.StatusOK, results)
}
