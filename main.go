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

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	gctx "github.com/gorilla/context"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

const genPost = `
Title: Example Post
Show: false
Time: %s

### Markdown here
`

type Config struct {
	// Listen changes what ip/port the server is listening on.
	Listen string
	// TLS enables and disables the use of TLS.
	TLS bool
	// CertFile and KeyFile are the system paths to the SSL certificate and
	// key file which can be used when TLS is enabled.
	CertFile, KeyFile string
	// Proxy when enabled, will allow the webserver to sit behind another
	// webserver, that supports X-Forwarded-For and similar headers to provide
	// the server behind the proxy with the real IP address.
	Proxy bool
	// If the templates/style should be cached upon first load.
	Cache bool
}

var conf Config
var debug = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

func runServer() {
	r := chi.NewRouter()
	if conf.Proxy {
		r.Use(middleware.RealIP)
	}

	r.Use(middleware.DefaultCompress)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: debug}))
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.Recoverer)
	FileServer(r, "/static", http.Dir("static"))

	r.Get("/", getPost)
	r.Get("/post/{uid}", getPost)

	r.Get("/g/{pkg}", func(w http.ResponseWriter, r *http.Request) {
		tmpl(w, r, "static/views/go.html", map[string]interface{}{"pkg": chi.URLParam(r, "pkg")})
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

	debug.Println("initializing webserver")
	if conf.TLS {
		debug.Fatal(http.ListenAndServeTLS(conf.Listen, mustFile(conf.CertFile), mustFile(conf.KeyFile), gctx.ClearHandler(r)))
	} else {
		debug.Fatal(http.ListenAndServe(conf.Listen, gctx.ClearHandler(r)))
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		debug.Printf("warn: error loading .env file: %s", err)
	}

	app := cli.NewApp()
	app.HideVersion = true
	app.Name = "liam.sh"
	app.Usage = "Liam's personal website"
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run the webserver",
			Action: func(ctx *cli.Context) {
				runServer()
			},
			Aliases: []string{"start"},
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "tls", Destination: &conf.TLS, Usage: "Enable TLS"},
				cli.BoolFlag{Name: "proxy", Destination: &conf.Proxy, Usage: "Use X-Forwarded-For (ONLY IF PROXYING!)"},
				cli.BoolFlag{Name: "cache", Destination: &conf.Cache, Usage: "Cache templates"},
				cli.StringFlag{Name: "b", Destination: &conf.Listen, Usage: "Bind `address`. E.g. 0.0.0.0:8080"},
				cli.StringFlag{Name: "cert", Destination: &conf.CertFile, Usage: "Certificate file `path`"},
				cli.StringFlag{Name: "key", Destination: &conf.KeyFile, Usage: "Certificate key `path`"},
			},
		},
		{
			Name:  "gen",
			Usage: "Generate an example post and print to STDOUT",
			Action: func(ctx *cli.Context) {
				fmt.Println(fmt.Sprintf(genPost, FormatTime(time.Now())))
			},
			Aliases: []string{"gen-post", "generate"},
		},
	}

	app.Action = func(ctx *cli.Context) error {
		if len(os.Args) == 1 {
			return errors.New("error: please supply a sub-command (see --help)")
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		debug.Printf("unable to instantiate cli parser: %s", err)
		log.Fatal(err)
	}
}
