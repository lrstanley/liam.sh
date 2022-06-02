// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghhandler

import (
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/liam.sh/internal/ent"
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
	r.Get("/svg", h.getProjectSVG)
	r.Get("/svg/{owner}/{repo}", h.getProjectSVG)
}
