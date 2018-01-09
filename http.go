// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	rice "github.com/GeertJohan/go.rice"
	humanize "github.com/flosch/go-humanize"
	"github.com/flosch/pongo2"
	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/go-github/github"
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
	GitArgs GitArgs `group:"Git Options" namespace:"git"`
}

func (h *HTTPArgs) Execute(_ []string) error {
	tmpl = pt.New("", pt.Config{
		CacheParsed: !h.Debug,
		Loader:      rice.MustFindBox("static").Bytes,
		ErrorLogger: os.Stderr,
		DefaultCtx: func(w http.ResponseWriter, r *http.Request) (ctx map[string]interface{}) {
			return pt.M{
				"git": gc.user.Load().(*github.User),
			}
		},
	})

	fsw, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer fsw.Close()

	if err = fsw.Add("posts/"); err != nil {
		panic(err)
	}

	go func() {
		for err := range fsw.Errors {
			if err != nil {
				log.Printf("error in fsnotify: %s", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	go func() {
		// Make sure that initially it gets updated too.
		pc.update("posts/**")
		time.Sleep(5 * time.Second)

		for event := range fsw.Events {
			log.Printf("new fsevent: %v", event)
			pc.update("posts/**")
			time.Sleep(5 * time.Second)
		}
	}()

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

	r.NotFound(notFoundHandler)
	r.Get("/", getPost)
	r.Get("/post/{uid}", getPost)

	r.Get("/g/{pkg}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Render(w, r, "/tmpl/go.html", pt.M{"pkg": chi.URLParam(r, "pkg")})
	})

	// r.Get("/api/user", func(w http.ResponseWriter, r *http.Request) {
	// 	pt.JSON(w, r, gc.user.Load().(*github.User))
	// })
	// r.Get("/api/repos", func(w http.ResponseWriter, r *http.Request) {
	// 	gc.RLock()
	// 	repos := gc.repos
	// 	gc.RUnlock()

	// 	pt.JSON(w, r, repos)
	// })

	r.Get("/ghr.json", gitRequestAllAssets)
	r.Get("/ghr", gitRequestAllAssets)
	r.Get("/ghr/{asset}", gitRequestAsset)

	// Pre-fetch git data just to make sure when we're starting that credentials
	// and similar are functioning.
	if !h.GitArgs.ReleaseSkip {
		if err := gitUpdateReleaseCache(); err != nil {
			panic(err)
		}
		go func() {
			for {
				time.Sleep(h.GitArgs.ReleaseCheck)
				if err := gitUpdateReleaseCache(); err != nil {
					fmt.Printf("error updating release cache: %s", err)
				}
			}
		}()
	}

	if err := gitUpdateUserCache(); err != nil {
		panic(err)
	}
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			if err := gitUpdateUserCache(); err != nil {
				fmt.Printf("error updating user cache: %s", err)
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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl.Render(w, r, "/tmpl/404.html", nil)
}

func init() {
	pongo2.RegisterFilter("field", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		if in == nil || in.IsNil() {
			return pongo2.AsValue(nil), nil
		}

		obj, isMap := in.Interface().(map[string]string)
		if !isMap {
			return pongo2.AsValue(nil), nil
		}

		res, found := obj[param.String()]
		if found {
			return pongo2.AsValue(res), nil
		}

		return pongo2.AsValue(nil), nil
	})

	pongo2.RegisterFilter("naturaltime", func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		basetime, isTime := in.Interface().(time.Time)
		if !isTime {
			return nil, &pongo2.Error{
				Sender:    "filter:timeuntil/timesince",
				OrigError: errors.New("time-value is not a time.Time-instance"),
			}
		}
		var paramtime time.Time
		if !param.IsNil() {
			paramtime, isTime = param.Interface().(time.Time)
			if !isTime {
				return nil, &pongo2.Error{
					Sender:    "filter:timeuntil/timesince",
					OrigError: errors.New("time-parameter is not a time.Time-instance"),
				}
			}
		} else {
			paramtime = time.Now()
		}

		return pongo2.AsValue(humanize.TimeDuration(basetime.Sub(paramtime))), nil
	})
}
