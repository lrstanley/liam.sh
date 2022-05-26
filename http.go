// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	_ "embed"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/graphql"
	"github.com/lrstanley/liam.sh/internal/httpware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

// go:embed internal/ent/openapi.json
var openapi []byte

func httpServer() *http.Server {
	r := chi.NewRouter()

	if len(cli.Flags.HTTP.TrustedProxies) > 0 {
		r.Use(chix.UseRealIP(cli.Flags.HTTP.TrustedProxies, chix.OptUseXForwardedFor))
	}

	r.Use(chix.UseContextIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(chix.UseStructuredLogger(logger))
	r.Use(chix.UseNextURL)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Compress(5))
	r.Use(httprate.LimitByIP(400, 5*time.Minute))
	r.Use(chix.UseRobotsTxt(""))
	r.Use(chix.UseSecurityTxt(&chix.SecurityConfig{
		ExpiresIn: 182 * 24 * time.Hour,
		Contacts: []string{
			"mailto:me@liamstanley.io",
			"https://liam.sh/chat",
			"https://github.com/lrstanley",
		},
		KeyLinks:  []string{"https://github.com/lrstanley.gpg"},
		Languages: []string{"en"},
	}))

	r.Use(cors.AllowAll().Handler)

	goth.UseProviders(
		github.New(
			cli.Flags.Github.ClientID,
			cli.Flags.Github.ClientSecret,
			cli.Flags.HTTP.BaseURL+"/api/auth/providers/github/callback",
		),
	)

	auth := chix.NewAuthHandler[ent.User, int](
		database.NewAuthService(db, cli.Flags.Github.User),
		cli.Flags.HTTP.ValidationKey,
		cli.Flags.HTTP.EncryptionKey,
	)
	r.Use(auth.AddToContext)
	r.Use(httpware.EvictCacheAdmin)

	r.Mount("/api/graphql", graphql.New(db, cli))
	r.Mount("/api/playground", playground.Handler("GraphQL playground", "/api/graphql"))
	r.Mount("/api/auth", auth)
	r.Mount("/chat", http.RedirectHandler(cli.Flags.ChatLink, http.StatusTemporaryRedirect))

	if !cli.Debug {
		r.With(chix.UsePrivateIP).Mount("/debug", middleware.Profiler())
	}

	r.NotFound(catchAll)

	return &http.Server{
		Addr:    cli.Flags.HTTP.BindAddr,
		Handler: r,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/") {
		chix.Error(w, r, http.StatusNotFound, chix.ErrMatchStatus)
		return
	}

	if r.Method != http.MethodGet {
		chix.Error(w, r, http.StatusMethodNotAllowed, chix.ErrMatchStatus)
		return
	}
	// if strings.HasSuffix(r.URL.Path, ".ico") {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// w.Write(rice.MustFindBox("public").MustBytes("index.html"))
}
