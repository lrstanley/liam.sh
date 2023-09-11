// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"embed"
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
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/post"
	"github.com/lrstanley/liam.sh/internal/graphql"
	"github.com/lrstanley/liam.sh/internal/handlers/ghhandler"
	"github.com/lrstanley/liam.sh/internal/handlers/webhookhandler"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//go:generate sh -c "mkdir -vp public/dist;touch public/dist/index.html"
//go:embed all:public/dist
var staticFS embed.FS

func httpServer(ctx context.Context) *http.Server {
	chix.DefaultAPIPrefix = "/-/"
	r := chi.NewRouter()

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

	if len(cli.Flags.HTTP.TrustedProxies) > 0 {
		r.Use(chix.UseRealIPCLIOpts(cli.Flags.HTTP.TrustedProxies))
	}

	// Core middeware.
	r.Use(
		chix.UseDebug(cli.Debug),
		chix.UseContextIP,
		middleware.RequestID,
		chix.UseStructuredLogger(logger),
		chix.UsePrometheus,
		chix.Recoverer,
		middleware.Maybe(middleware.StripSlashes, func(r *http.Request) bool {
			return !strings.HasPrefix(r.URL.Path, "/debug/")
		}),
		middleware.Compress(5),
		chix.UseNextURL,
	)

	// Security related.
	if !cli.Debug {
		r.Use(middleware.SetHeader("Strict-Transport-Security", "max-age=31536000"))
	}
	r.Use(
		cors.AllowAll().Handler,
		chix.UseHeaders(map[string]string{
			"Content-Security-Policy": "default-src 'self'; img-src * data:; media-src * data:; style-src 'self' 'unsafe-inline'; object-src 'none'; child-src 'none'; frame-src 'none'; worker-src 'none'",
			"X-Frame-Options":         "DENY",
			"X-Content-Type-Options":  "nosniff",
			"Referrer-Policy":         "no-referrer-when-downgrade",
			"Permissions-Policy":      "clipboard-write=(self)",
		}),
		auth.AddToContext,
		httprate.LimitByIP(200, 1*time.Minute),
	)

	// Legacy redirects.
	r.Use(
		middleware.PathRewrite("/ghr/", "/-/gh/dl-legacy/"),
		middleware.PathRewrite("/ghr.json", "/-/gh/dl-legacy"),
	)

	// Misc.
	r.Use(
		chix.UseRobotsTxt(fmt.Sprintf(
			"User-agent: *\nSitemap: %s/sitemap.txt\nAllow: /\n",
			strings.TrimRight(cli.Flags.HTTP.BaseURL, "/"),
		)),
		chix.UseSecurityTxt(&chix.SecurityConfig{
			ExpiresIn: 182 * 24 * time.Hour,
			Contacts: []string{
				"mailto:me@liamstanley.io",
				"https://liam.sh/chat",
				"https://github.com/lrstanley",
			},
			KeyLinks:  []string{"https://github.com/lrstanley.gpg"},
			Languages: []string{"en"},
		}),
	)

	r.Mount("/chat", http.RedirectHandler(cli.Flags.ChatLink, http.StatusTemporaryRedirect))
	r.Mount("/-/graphql", graphql.New(db, cli))
	r.With(middleware.SetHeader(
		"Content-Security-Policy",
		"default-src * 'unsafe-inline' 'unsafe-eval' data: blob:; ",
	)).Mount("/-/playground", playground.Handler("GraphQL playground", "/-/graphql"))
	r.Mount("/-/auth", auth)
	r.Route("/-/gh", ghhandler.New(db).Route)
	r.Route("/-/webhook/discord", webhookhandler.New().Route)
	r.With(chix.UsePrivateIP).Mount("/metrics", promhttp.Handler())

	if cli.Debug {
		r.With(chix.UsePrivateIP).Mount("/debug", middleware.Profiler())
	}

	r.With(middleware.ThrottleBacklog(1, 5, 5*time.Second)).Get("/-/healthy", func(w http.ResponseWriter, r *http.Request) {
		err := database.Ping(ctx)
		if err != nil {
			chix.Error(w, r, err)
			return
		}

		chix.JSON(w, r, 200, chix.M{
			"status": "ok",
		})
	})

	r.Get("/sitemap.txt", sitemap)
	r.NotFound(chix.UseStatic(ctx, &chix.Static{
		FS:         staticFS,
		CatchAll:   true,
		AllowLocal: cli.Debug,
		Path:       "public/dist",
		SPA:        true,
		Headers: map[string]string{
			"Vary":          "Accept-Encoding",
			"Cache-Control": "public, max-age=7776000",
		},
	}).ServeHTTP)

	return &http.Server{
		Addr:    cli.Flags.HTTP.BindAddr,
		Handler: r,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func sitemap(w http.ResponseWriter, r *http.Request) {
	urls := []string{
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
