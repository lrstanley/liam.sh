// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/cmd/httpserver/handlers/apihandler"
	"github.com/lrstanley/liam.sh/cmd/httpserver/handlers/ghhandler"
	"github.com/lrstanley/liam.sh/cmd/httpserver/handlers/webhookhandler"
	"github.com/lrstanley/liam.sh/internal/auth"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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

	authSvc := auth.NewService(db, cli.Flags.Github.User)

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
		chix.UseRecoverer,
		middleware.Maybe(middleware.StripSlashes, func(r *http.Request) bool {
			return !strings.HasPrefix(r.URL.Path, "/debug/")
		}),
		middleware.Compress(5),
		chix.UseNextURL,
	)

	// Security related.
	r.Use(
		cors.New(cors.Options{
			AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodOptions,
			},
			AllowedHeaders:   []string{"*"},
			MaxAge:           300,
			AllowCredentials: true,
		}).Handler,
		chix.UseAuthContext(authSvc),
		httprate.LimitByIP(400, 1*time.Minute),
	)

	// Legacy redirects.
	r.Use(
		middleware.PathRewrite("/ghr/", "/-/gh/dl-legacy/"),
		middleware.PathRewrite("/ghr.json", "/-/gh/dl-legacy"),
	)

	// Misc.
	r.Use(
		chix.UseSecurityTxt(&chix.SecurityConfig{
			ExpiresIn: 182 * 24 * time.Hour,
			Contacts: []string{
				"mailto:liam@liam.sh",
				"https://liam.sh/chat",
				"https://github.com/lrstanley",
			},
			KeyLinks:  []string{"https://github.com/lrstanley.gpg"},
			Languages: []string{"en"},
		}),
	)

	r.Route("/-", apihandler.New(db, aiSvc, cli.VersionInfo, cli.Debug, "/-").Route)
	r.Mount("/-/auth", chix.NewAuthHandler(
		authSvc,
		cli.Flags.HTTP.ValidationKey,
		cli.Flags.HTTP.EncryptionKey,
	))
	r.Route("/-/gh", ghhandler.New(db).Route)
	r.Route("/-/webhook/discord", webhookhandler.New().Route)
	r.With(chix.UsePrivateIP).Mount("/metrics", promhttp.Handler())

	if cli.Debug {
		r.With(chix.UsePrivateIP).Mount("/debug", middleware.Profiler())
	}

	r.With(middleware.ThrottleBacklog(1, 5, 5*time.Second)).Get("/-/healthy", func(w http.ResponseWriter, r *http.Request) {
		chix.JSON(w, r, 200, chix.M{
			"status": "ok",
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		chix.Error(w, r, chix.WrapCode(http.StatusNotFound))
	})

	return &http.Server{
		Addr:    cli.Flags.HTTP.BindAddr,
		Handler: r,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
