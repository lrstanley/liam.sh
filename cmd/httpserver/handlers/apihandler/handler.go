// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/rest"
	"github.com/lrstanley/liam.sh/internal/models"
)

type handler struct {
	db      *ent.Client
	srv     *rest.Server
	version *clix.VersionInfo[models.Flags]
}

func New(db *ent.Client, version *clix.VersionInfo[models.Flags], debug bool, base string) *handler {
	restServer, err := rest.NewServer(db, &rest.ServerConfig{
		BasePath:   base,
		MaskErrors: !debug,
	})
	if err != nil {
		panic(err)
	}

	return &handler{
		db:      db,
		srv:     restServer,
		version: version,
	}
}

func (h *handler) Route(r chi.Router) {
	r.Route("/", h.srv.Handler)
	r.Get("/version", h.getVersion)
	r.Get("/self", h.getSelf)
	r.Get("/labels/count", h.getLabelsCount)
	r.Get("/github-user", h.getGithubUser)
	r.Get("/github-releases/outdated", h.getOutdatedGithubReleases)
	r.Post("/posts/regenerate", h.postsRegenerate)
	r.Get("/stats/coding", h.getStatsCoding)
	r.Get("/stats/github", h.getStatsGithub)
}
