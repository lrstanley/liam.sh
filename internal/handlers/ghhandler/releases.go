// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghhandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
)

func (h *handler) getReleasesLegacy(w http.ResponseWriter, r *http.Request) {
	asset := chi.URLParam(r, "asset")

	if asset == "" {
		assets, err := h.db.GithubAsset.Query().Order(ent.Desc(githubasset.FieldCreatedAt)).
			Select(githubasset.FieldName).
			Strings(r.Context())
		if err != nil {
			chix.Error(w, r, err)
			return
		}

		chix.JSON(w, r, http.StatusOK, chix.M{"assets": assets})
		return
	}

	a, err := h.db.GithubAsset.Query().
		Order(ent.Asc(githubasset.FieldName)).
		Where(githubasset.Name(asset)).First(r.Context())
	if err != nil {
		chix.ErrorCode(w, r, 404, err)
		return
	}

	http.Redirect(w, r, a.BrowserDownloadURL, http.StatusTemporaryRedirect)
}

func (h *handler) getReleases(w http.ResponseWriter, r *http.Request) {
	repo := chi.URLParam(r, "repo")
	version := strings.TrimLeft(chi.URLParam(r, "version"), "v")
	asset := chi.URLParam(r, "asset")
	text, _ := strconv.ParseBool(r.FormValue("text"))

	var err error

	// Return just the releases for the given repository.
	if version == "" {
		releases, err := h.db.GithubRelease.Query().Where(
			githubrelease.HasRepositoryWith(githubrepository.NameEqualFold(repo)),
			githubrelease.Draft(false),
			githubrelease.Prerelease(false),
		).
			Order(ent.Desc(githubrelease.FieldCreatedAt)).
			Select(githubrelease.FieldTagName).Strings(r.Context())

		if chix.Error(w, r, err) {
			return
		}

		if text {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			for _, item := range releases {
				fmt.Fprintf(w, "%v\n", item)
			}
		} else {
			chix.JSON(w, r, http.StatusOK, chix.M{"data": releases})
		}
		return
	}

	// Resolve the latest non-draft and non-pre-release release for the given
	// repository.
	if version == "latest" {
		version, err = h.db.GithubRelease.Query().Where(
			githubrelease.HasRepositoryWith(githubrepository.NameEqualFold(repo)),
			githubrelease.Draft(false),
			githubrelease.Prerelease(false),
		).
			Order(ent.Desc(githubrelease.FieldCreatedAt)).
			Limit(1).Select(githubrelease.FieldTagName).
			String(r.Context())

		if chix.Error(w, r, err) {
			return
		}
	}

	// Fetch the release for the given version.
	var release *ent.GithubRelease
	release, err = h.db.GithubRelease.Query().Where(
		githubrelease.HasRepositoryWith(githubrepository.NameEqualFold(repo)),
		githubrelease.Or(
			githubrelease.TagName(version),
			githubrelease.TagName("v"+version),
		),
	).First(r.Context())

	if chix.Error(w, r, err) {
		return
	}

	// List the assets for the given release.
	if asset == "" {
		var assets []string

		assets, err = release.QueryAssets().
			Select(githubasset.FieldName).Strings(r.Context())

		if chix.Error(w, r, err) {
			return
		}

		if text {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			for _, item := range assets {
				fmt.Fprintf(w, "%v\n", item)
			}
		} else {
			chix.JSON(w, r, http.StatusOK, chix.M{
				"release": release,
				"assets":  assets,
			})
		}
		return
	}

	// Redirect to the assets download URL.
	var downloadURL string
	downloadURL, err = release.QueryAssets().Where(
		githubasset.NameEqualFold(asset),
	).Limit(1).Select(githubasset.FieldBrowserDownloadURL).String(r.Context())

	if chix.Error(w, r, err) {
		return
	}

	http.Redirect(w, r, downloadURL, http.StatusTemporaryRedirect)
}
