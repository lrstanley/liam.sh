// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package githubhandler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/gh"
)

func New(db *ent.Client) *handler {
	return &handler{
		db: db,
	}
}

type handler struct {
	db *ent.Client
}

func (h *handler) Route(r chi.Router) {
	r.Get("/me", h.me)
}

func (h *handler) me(w http.ResponseWriter, r *http.Request) {
	user := gh.User.Load()
	if user == nil {
		chix.Error(w, r, http.StatusServiceUnavailable, errors.New("information not available yet"))
		return
	}

	w.WriteHeader(http.StatusOK)
	chix.JSON(w, r, http.StatusOK, chix.M{
		"user": user,
	})
}
