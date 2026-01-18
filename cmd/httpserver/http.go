// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"embed"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/lrstanley/chix/v2"
	"github.com/lrstanley/chix/xauth/v2"
	"github.com/lrstanley/liam.sh/cmd/httpserver/handlers/apihandler"
	"github.com/lrstanley/liam.sh/cmd/httpserver/handlers/ghhandler"
	"github.com/lrstanley/liam.sh/cmd/httpserver/handlers/webhookhandler"
	"github.com/lrstanley/liam.sh/internal/auth"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

//go:embed all:public
var frontendFS embed.FS

func httpServer(logger *slog.Logger) *http.Server {
	r := chi.NewRouter()

	r.Use(
		chix.NewConfig().
			SetAPIBasePath("/-/").
			SetLogger(logger).
			SetMaskPrivateErrors(cli.Debug).
			SetErrorResolvers(models.ErrorResolver).
			Use(),
	)

	goth.UseProviders(
		github.New(
			cli.Flags.Github.ClientID,
			cli.Flags.Github.ClientSecret,
			cli.Flags.HTTP.BaseURL+"/-/auth/providers/github/callback",
			"read:user",
			"user:email",
		),
	)

	authSvc := auth.NewService(db, cli.Flags.Github.User)

	if len(cli.Flags.HTTP.TrustedProxies) > 0 {
		r.Use(chix.UseRealIPStringOpts(cli.Flags.HTTP.TrustedProxies))
	}

	r.Use(
		chix.UseContextIP(),
		chix.UseRequestID(),
		chix.UseStripSlashes(),
		chix.UseStructuredLogger(chix.DefaultLogConfig()),
		middleware.Compress(5),
		chix.UseNextURL(),
		chix.UseCrossOriginProtection("http://localhost:8081", "https://liam.sh", "https://*.liam.sh"),
		chix.UseCrossOriginResourceSharing(&chix.CORSConfig{
			AllowedOrigins: []string{"http://localhost:8081", "https://liam.sh", "https://*.liam.sh"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
			},
			AllowedHeaders: []string{
				"Accept",
				"Content-Type",
				"Origin",
			},
			// AllowPrivateNetwork: true,
			AllowCredentials: true,
		}),
		xauth.UseAuthContext(authSvc),
		httprate.LimitByIP(400, 1*time.Minute),
	)

	// Legacy redirects.
	r.Use(
		middleware.PathRewrite("/ghr/", "/-/gh/dl-legacy/"),
		middleware.PathRewrite("/ghr.json", "/-/gh/dl-legacy"),
	)

	// Misc.
	r.Use(
		chix.UseSecurityText(chix.SecurityTextConfig{
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

	r.Route("/-", apihandler.New(db, aiSvc, cli.GetVersion().NonSensitive(), cli.Debug, "/-").Route)
	r.Mount("/-/auth", xauth.NewGothHandler(&xauth.GothConfig[ent.User, int]{
		Service:        authSvc,
		SessionStorage: xauth.NewCookieStore(cli.Flags.HTTP.ValidationKey, cli.Flags.HTTP.EncryptionKey),
	}))
	r.Route("/-/gh", ghhandler.New(db).Route)
	r.Route("/-/webhook/discord", webhookhandler.New().Route)

	if cli.Debug {
		r.With(chix.UsePrivateIP()).Mount("/debug", middleware.Profiler())
	}

	r.With(middleware.ThrottleBacklog(1, 5, 5*time.Second)).Get("/-/healthy", func(w http.ResponseWriter, r *http.Request) {
		chix.JSON(w, r, 200, chix.M{
			"status": "ok",
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		chix.ErrorWithCode(w, r, http.StatusNotFound)
	})

	// r.With(chix.UseHeaders(map[string]string{
	// 	"Vary":          "Accept-Encoding",
	// 	"Cache-Control": "public, max-age=3600",
	// })).Mount("/", chix.UseStatic(&chix.StaticConfig{
	// 	FS:         frontendFS,
	// 	Prefix:     "/",
	// 	SPA:        true,
	// 	AllowLocal: true,
	// 	Path:       "public",
	// }))

	return &http.Server{
		Addr:    cli.Flags.HTTP.BindAddr,
		Handler: r,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
