// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	gctx "github.com/gorilla/context"
	"github.com/lrstanley/pt"
)

var tmpl *pt.Loader

type HTTPArgs struct {
	Debug bool   `long:"debug" description:"enable debugging (pprof endpoints) and disables template caching"`
	Bind  string `short:"b" long:"bind" default:":8080" description:"ip:port pair to bind to"`
	Proxy bool   `short:"p" long:"behind-proxy" description:"if X-Forwarded-For headers should be trusted"`
	TLS   struct {
		Enable bool   `long:"enable" description:"run tls server rather than standard http"`
		Cert   string `short:"c" long:"cert" description:"path to ssl cert file"`
		Key    string `short:"k" long:"key" description:"path to ssl key file"`
	} `group:"TLS Options" namespace:"tls"`
}

func (h *HTTPArgs) Execute(_ []string) error {
	tmpl = pt.New("", pt.Config{
		CacheParsed: !h.Debug,
		Loader:      rice.MustFindBox("static").Bytes,
		ErrorLogger: os.Stderr,
	})

	r := chi.NewRouter()
	if h.Proxy {
		r.Use(middleware.RealIP)
	}

	r.Use(middleware.DefaultCompress)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: debug}))
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.Recoverer)
	pt.FileServer(r, "/static", rice.MustFindBox("static").HTTPBox())

	if h.Debug {
		r.Mount("/debug", middleware.Profiler())
	}

	r.Get("/", getPost)
	r.Get("/post/{uid}", getPost)

	r.Get("/g/{pkg}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Render(w, r, "/tmpl/go.html", pt.M{"pkg": chi.URLParam(r, "pkg")})
	})

	r.Get("/ghr/{asset}", func(w http.ResponseWriter, r *http.Request) {
		gc.RLock()
		asset, ok := gc.assetMap[chi.URLParam(r, "asset")]
		gc.RUnlock()

		if !ok {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, asset, http.StatusFound)
	})

	if err := updateReleaseCache(); err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(5 * time.Minute)
			if err := updateReleaseCache(); err != nil {
				fmt.Printf("error updating cache: %s", err)
			}
		}
	}()

	srv := &http.Server{
		Handler:      gctx.ClearHandler(r),
		Addr:         h.Bind,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 45 * time.Second,
		ErrorLog:     debug,
	}

	debug.Println("initializing webserver")
	if h.TLS.Enable {
		debug.Fatal(srv.ListenAndServeTLS(h.TLS.Cert, h.TLS.Key))
	}

	debug.Fatal(srv.ListenAndServe())

	return nil
}
