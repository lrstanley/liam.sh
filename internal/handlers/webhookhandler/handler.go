// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package webhookhandler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	"github.com/google/go-github/v48/github"
	"github.com/lrstanley/chix"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Route(r chi.Router) {
	r.Post("/api/webhooks/*", h.Discord)
}

func (h *handler) Discord(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://discord.com/api/webhooks/%s", chi.URLParam(r, "*"))

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

	sender := f.Interface().(*github.User)
	if sender != nil && sender.GetType() == "Bot" {
		chix.Log(r).WithFields(log.Fields{
			"event_type": github.WebHookType(r),
			"bot":        sender.Login,
		}).Info("ignoring bot event")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	chix.Log(r).WithField("event_type", github.WebHookType(r)).Info("forwarding webhook to discord")
	req, err := http.NewRequestWithContext(r.Context(), "POST", url, bytes.NewBuffer(payload))
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

	chix.Log(r).WithField("discord_status", resp.StatusCode).Info("discord response")
}
