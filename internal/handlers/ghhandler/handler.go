// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghhandler

import (
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/liam.sh/internal/database/ent"
)

type handler struct {
	db         *ent.Client
	projectSVG *template.Template
}

func New(db *ent.Client) *handler {
	return &handler{
		db:         db,
		projectSVG: template.Must(template.New("").Parse(rawSVG)),
	}
}

func (h *handler) Route(r chi.Router) {
	// SVG banner generation.
	r.Get("/svg", h.getProjectSVG)
	r.Get("/svg/{owner}/{repo}", h.getProjectSVG)

	// Releases/assets.
	r.Get("/dl/{repo:^[a-zA-Z0-9_.-]{1,100}$}", h.getReleases)
	r.Get("/dl/{repo:^[a-zA-Z0-9_.-]{1,100}$}/{version:^[a-zA-Z0-9_.-]{1,100}$}", h.getReleases)
	r.Get("/dl/{repo:^[a-zA-Z0-9_.-]{1,100}$}/{version:^[a-zA-Z0-9_.-]{1,100}$}/{asset}", h.getReleases)

	// Gists.
	r.Get("/g/{asset}", h.getGists)
	r.Get("/g/{id:^[a-zA-Z0-9]{1,100}$}/{asset}", h.getGists)

	// Legacy routes.
	r.Get("/dl-legacy", h.getReleasesLegacy)
	r.Get("/dl-legacy/{asset}", h.getReleasesLegacy)
}
