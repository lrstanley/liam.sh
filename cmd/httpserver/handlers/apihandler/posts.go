// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"net/http"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database/ent/post"
)

func (h *handler) postsRegenerate(w http.ResponseWriter, r *http.Request) {
	p := []struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}{}

	tx, err := h.db.Tx(r.Context())
	if chix.Error(w, r, err) {
		return
	}

	if err = tx.Post.Query().Select(post.FieldID, post.FieldContent).Scan(r.Context(), &p); err != nil {
		_ = tx.Rollback()
		chix.Error(w, r, err)
		return
	}

	for _, v := range p {
		if v.Content == "" {
			continue
		}

		if err = tx.Post.UpdateOneID(v.ID).SetContent(v.Content).Exec(r.Context()); err != nil {
			_ = tx.Rollback()
			chix.Error(w, r, err)
			return
		}
		log.FromContext(r.Context()).WithField("id", v.ID).Debug("regenerated post")
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		chix.Error(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
