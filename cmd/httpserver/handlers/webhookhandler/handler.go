// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package webhookhandler

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
	"github.com/google/go-github/v82/github"
	"github.com/lrstanley/chix/v2"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Route(r chi.Router) {
	r.Post("/api/webhooks/*", h.Discord)
}

func (h *handler) Discord(w http.ResponseWriter, r *http.Request) {
	url := "https://discord.com/api/webhooks/" + chi.URLParam(r, "*")

	payload, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		chix.Error(w, r, err)
		return
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		chix.Error(w, r, err)
		return
	}

	e := reflect.ValueOf(event)
	f := reflect.Indirect(e).FieldByName("Sender")

	sender, ok := f.Interface().(*github.User)
	if ok && sender != nil && sender.GetType() == "Bot" {
		chix.LogInfo(
			r.Context(), "ignoring bot event",
			slog.String("event_type", github.WebHookType(r)),
			slog.String("bot", sender.GetLogin()),
		)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	chix.LogInfo(
		r.Context(), "forwarding webhook to discord",
		slog.String("event_type", github.WebHookType(r)),
	)
	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		chix.Error(w, r, err)
		return
	}

	req.Header = r.Header.Clone()

	w.WriteHeader(http.StatusOK)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		chix.Error(w, r, err)
		return
	}
	defer resp.Body.Close()
	_, _ = io.Copy(w, resp.Body)

	chix.LogInfo(
		r.Context(), "discord response",
		slog.Int("discord_status", resp.StatusCode),
	)
}
