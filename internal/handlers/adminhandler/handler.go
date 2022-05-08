// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package adminhandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
)

func New(db *ent.Client, auth *chix.AuthHandler[ent.User, int]) *handler {
	return &handler{
		db:   db,
		auth: auth,
	}
}

type handler struct {
	db   *ent.Client
	auth *chix.AuthHandler[ent.User, int]
}

func (h *handler) Route(r chi.Router) {
	r.Get("/ping", h.ping)
}

func (h *handler) ping(w http.ResponseWriter, r *http.Request) {
	chix.JSON(w, r, http.StatusOK, chix.M{
		"message": "pong",
	})
}
