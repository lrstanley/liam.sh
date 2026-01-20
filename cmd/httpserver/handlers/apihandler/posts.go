// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/lrstanley/chix/v2"
	"github.com/lrstanley/liam.sh/internal/database/ent/post"
)

func (h *handler) postsRegenerate(w http.ResponseWriter, r *http.Request) {
	p := []struct {
		ID          int       `json:"id"`
		Content     string    `json:"content"`
		PublishedAt time.Time `json:"published_at"`
	}{}

	tx, err := h.db.Tx(r.Context())
	if chix.IfError(w, r, err) {
		return
	}

	err = tx.Post.Query().
		Select(
			post.FieldID,
			post.FieldContent,
			post.FieldPublishedAt,
		).
		Scan(r.Context(), &p)
	if err != nil {
		_ = tx.Rollback()
		chix.Error(w, r, err)
		return
	}

	for _, v := range p {
		if v.Content == "" {
			continue
		}

		err = tx.Post.UpdateOneID(v.ID).
			SetContent(v.Content).
			SetPublishedAt(v.PublishedAt.UTC()).
			Exec(r.Context())
		if err != nil {
			_ = tx.Rollback()
			chix.Error(w, r, err)
			return
		}
		chix.LogDebug(r.Context(), "regenerated post", slog.Int("id", v.ID))
	}

	err = tx.Commit()
	if chix.IfError(w, r, err) {
		_ = tx.Rollback()
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) postsSuggestLabels(w http.ResponseWriter, r *http.Request) {
	v := &struct {
		Content string `json:"content"`
	}{}
	if err := chix.Bind(r, v); err != nil {
		chix.Error(w, r, err)
		return
	}

	// Don't reuse context, because we want it to continue in the background and cache if
	// possible.
	suggestions, err := h.aiSvc.PostSuggestLabels(context.WithoutCancel(r.Context()), v.Content)
	if err != nil {
		chix.Error(w, r, err)
		return
	}
	chix.JSON(w, r, http.StatusOK, suggestions)
}
