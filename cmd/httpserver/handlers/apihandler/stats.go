// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"net/http"

	"github.com/lrstanley/chix/v2"
	"github.com/lrstanley/liam.sh/internal/gh"
	"github.com/lrstanley/liam.sh/internal/wakapi"
)

func (h *handler) getStatsCoding(w http.ResponseWriter, r *http.Request) {
	chix.JSON(w, r, http.StatusOK, wakapi.Statistics.Load())
}

func (h *handler) getStatsGithub(w http.ResponseWriter, r *http.Request) {
	chix.JSON(w, r, http.StatusOK, gh.Stats.Load())
}
