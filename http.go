// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/didip/tollbooth/v6"
	"github.com/didip/tollbooth/v6/limiter"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/lrstanley/liam.sh/internal/ent/ogent"
	"github.com/lrstanley/liam.sh/internal/handlers/adminhandler"
	"github.com/lrstanley/liam.sh/internal/handlers/authhandler"
	"github.com/lrstanley/liam.sh/internal/httpware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func httpServer(ctx context.Context) error {
	if !cli.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.RedirectTrailingSlash = true
	r.HandleMethodNotAllowed = true

	if err := r.SetTrustedProxies(cli.HTTP.TrustedProxies); err != nil {
		logger.WithError(err).Fatal("failed to set trusted proxies")
	}

	r.Use(gin.Recovery())
	r.Use(requestid.New())
	r.Use(httpware.ErrorCatcher)
	r.Use(httpware.StructuredLogger(logger, false))

	limit := tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	r.Use(httpware.LimitHandler(limit))

	// r.Use(timeout.New(timeout.WithTimeout(20 * time.Second)))

	// Setup session store and oauth2 configuration.
	authStore := sessions.NewCookieStore([]byte(cli.HTTP.SessionKey))
	authStore.MaxAge(30 * 86400)
	authStore.Options.Path = "/"
	authStore.Options.HttpOnly = true
	authStore.Options.Secure = !cli.Debug
	gothic.Store = authStore
	goth.UseProviders(
		github.New(
			cli.Github.ClientID,
			cli.Github.ClientSecret,
			cli.HTTP.BaseURL+"/api/auth/providers/github/callback",
		),
	)

	if cli.Debug {
		pprof.Register(r)
	}

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Misc standard routes.
	r.GET("/robots.txt", func(c *gin.Context) {
		c.String(http.StatusOK, "User-agent: *\nDisallow: /api\nAllow /\n")
	})
	r.GET("/security.txt", httpware.SecurityText)
	r.GET("/.well-known/security.txt", httpware.SecurityText)

	r.Use(httpware.InjectUser(db))
	adminhandler.New(db, r.Group("/api/admin"))
	authhandler.New(db, r.Group("/api/auth"), &cli.Github)

	dbHandler, err := ogent.NewServer(ogent.NewOgentHandler(db))
	if err != nil {
		logger.WithError(err).Fatal("failed to create ogent handler")
	}

	r.Any("/api/query/*any", gin.WrapH(http.StripPrefix("/api/query", dbHandler)))

	srv := &http.Server{
		Addr:    cli.HTTP.BindAddr,
		Handler: r,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ch := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.WithError(err).Fatal("http server error")
			ch <- err
		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
	case err := <-ch:
		return err
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctxTimeout)
}

// func catchAll(w http.ResponseWriter, r *http.Request) {
// 	if strings.HasPrefix(r.URL.Path, "/api/") {
// 		httpware.Error(w, r, http.StatusNotFound, nil)
// 		return
// 	}
// 	if r.Method != http.MethodGet {
// 		httpware.Error(w, r, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
// 		return
// 	}
// 	if strings.HasSuffix(r.URL.Path, ".ico") {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	w.Write(rice.MustFindBox("public").MustBytes("index.html"))
// }
