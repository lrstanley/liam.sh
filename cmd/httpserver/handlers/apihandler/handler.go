// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrstanley/chix/xauth/v2"
	"github.com/lrstanley/clix/v2"
	"github.com/lrstanley/liam.sh/internal/ai"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/rest"
)

type handler struct {
	db      *ent.Client
	aiSvc   ai.Service
	srv     *rest.Server
	version *clix.NonSensitiveVersion
}

func New(db *ent.Client, aiSvc ai.Service, version *clix.NonSensitiveVersion, debug bool, base string) *handler {
	restServer, err := rest.NewServer(db, &rest.ServerConfig{
		BasePath:   base,
		MaskErrors: !debug,
	})
	if err != nil {
		panic(err)
	}

	return &handler{
		db:      db,
		aiSvc:   aiSvc,
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
	r.With(
		xauth.UseAuthRequired[ent.User](),
		middleware.Throttle(1),
	).Post("/posts/suggest-labels", h.postsSuggestLabels)
	r.Get("/stats/coding", h.getStatsCoding)
	r.Get("/stats/github", h.getStatsGithub)
}
