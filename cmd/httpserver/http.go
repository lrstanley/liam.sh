// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
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
	"github.com/lrstanley/liam.sh/internal/ent/post"
	"github.com/lrstanley/liam.sh/internal/graphql"
	"github.com/lrstanley/liam.sh/internal/handlers/ghhandler"
	"github.com/lrstanley/liam.sh/internal/httpware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

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
	r.Use(middleware.Compress(9))
	r.Use(httprate.LimitByIP(400, 5*time.Minute))
	r.Use(chix.UseRobotsTxt(fmt.Sprintf(
		"User-agent: *\nSitemap: %s/sitemap.txt\nDisallow: /-/\nAllow: /\n",
		strings.TrimRight(cli.Flags.HTTP.BaseURL, "/"),
	)))
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
			cli.Flags.HTTP.BaseURL+"/-/auth/providers/github/callback",
		),
	)

	auth := chix.NewAuthHandler[ent.User, int](
		database.NewAuthService(db, cli.Flags.Github.User),
		cli.Flags.HTTP.ValidationKey,
		cli.Flags.HTTP.EncryptionKey,
	)
	r.Use(auth.AddToContext)
	r.Use(httpware.UseEvictCacheAdmin)

	r.Mount("/chat", http.RedirectHandler(cli.Flags.ChatLink, http.StatusTemporaryRedirect))
	r.Mount("/-/graphql", graphql.New(db, cli))
	r.Mount("/-/playground", playground.Handler("GraphQL playground", "/-/graphql"))
	r.Mount("/-/auth", auth)
	r.Route("/-/gh", ghhandler.New(db).Route)

	// Legacy functionality.
	r.Handle("/ghr", http.RedirectHandler("/-/gh/dl-legacy", http.StatusPermanentRedirect))
	r.Handle("/ghr.json", http.RedirectHandler("/-/gh/dl-legacy", http.StatusPermanentRedirect))
	r.Get("/ghr/{asset}", func(w http.ResponseWriter, r *http.Request) {
		http.RedirectHandler(fmt.Sprintf("/-/gh/dl-legacy/%s", chi.URLParam(r, "asset")), http.StatusPermanentRedirect)
	})

	if !cli.Debug {
		r.With(chix.UsePrivateIP).Mount("/debug", middleware.Profiler())
	}

	r.Get("/sitemap.txt", sitemap)
	r.NotFound(catchAll)

	return &http.Server{
		Addr:    cli.Flags.HTTP.BindAddr,
		Handler: r,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/-/") {
		chix.Error(w, r, chix.WrapCode(404))
		return
	}

	if r.Method != http.MethodGet {
		chix.Error(w, r, chix.WrapCode(405))
		return
	}
	// if strings.HasSuffix(r.URL.Path, ".ico") {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// w.Write(rice.MustFindBox("public").MustBytes("index.html"))
}

func sitemap(w http.ResponseWriter, r *http.Request) {
	urls := []string{
		"/",
		"/repos",
		"/posts",
	}

	postSlugs, err := db.Post.Query().Select(post.FieldSlug).Strings(r.Context())
	if chix.Error(w, r, err) {
		return
	}

	for _, slug := range postSlugs {
		urls = append(urls, fmt.Sprintf("/p/%s", slug))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	for _, url := range urls {
		fmt.Fprintf(w, "%s%s\n", strings.TrimRight(cli.Flags.HTTP.BaseURL, "/"), url)
	}
}
