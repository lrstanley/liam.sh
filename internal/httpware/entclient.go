// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package httpware

import (
	"net/http"

	"github.com/lrstanley/liam.sh/internal/ent"
)

func UseEntContext(db *ent.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(ent.NewContext(r.Context(), db)))
		})
	}
}
