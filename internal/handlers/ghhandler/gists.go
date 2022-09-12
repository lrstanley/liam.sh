// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghhandler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent/githubgist"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

func (h *handler) getGists(w http.ResponseWriter, r *http.Request) {
	asset := chi.URLParam(r, "asset")
	id := chi.URLParam(r, "id")

	filter := []predicate.GithubGist{
		githubgist.NameEqualFold(asset),
	}

	if id != "" {
		filter = append(filter, githubgist.GistID(id))
	}

	gist, err := h.db.GithubGist.Query().Where(filter...).First(r.Context())
	if chix.Error(w, r, err) {
		return
	}

	if r.FormValue("latest") != "" {
		http.Redirect(w, r, gist.RawURL, http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("Content-Type", gist.Type)
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Last-Modified", gist.UpdatedAt.Format(http.TimeFormat))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, gist.Content)
}
