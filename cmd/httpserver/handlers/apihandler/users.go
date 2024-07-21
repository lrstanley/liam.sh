// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"net/http"

	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/rest"
	"github.com/lrstanley/liam.sh/internal/gh"
)

func (h *handler) getSelf(w http.ResponseWriter, r *http.Request) {
	ident := chix.IdentFromContext[ent.User](r.Context())
	if ident == nil {
		chix.ErrorCode(w, r, http.StatusUnauthorized, chix.WrapCode(http.StatusUnauthorized))
		return
	}

	// Right now we don't eager load additional things for users, so don't bother with this.
	// user, err := rest.EagerLoadUser(h.db.User.Query().Where(user.ID(ident.ID))).Only(r.Context())
	// if err != nil {
	// 	chix.Error(w, r, err)
	// 	return
	// }
	// rest.JSON(w, r, http.StatusOK, user)

	rest.JSON(w, r, http.StatusOK, ident)
}

func (h *handler) getGithubUser(w http.ResponseWriter, r *http.Request) {
	user := gh.User.Load()
	if user == nil {
		chix.ErrorCode(w, r, http.StatusNotFound, chix.WrapCode(http.StatusNotFound))
		return
	}
	rest.JSON(w, r, http.StatusOK, user)
}
