// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package httpware

import (
	"net/http"

	"ariga.io/entcache"
	"github.com/lrstanley/chix"
)

func UseEvictCacheAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if chix.RolesFromContext(r.Context()).Has("admin") {
			next.ServeHTTP(w, r.WithContext(entcache.Evict(r.Context())))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func UseSkipCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(entcache.Skip(r.Context())))
	})
}
