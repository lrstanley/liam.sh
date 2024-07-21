// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"net/http"

	"github.com/lrstanley/liam.sh/internal/database/ent/rest"
)

func (h *handler) getVersion(w http.ResponseWriter, r *http.Request) {
	rest.JSON(w, r, http.StatusOK, h.version.NonSensitive())
}
