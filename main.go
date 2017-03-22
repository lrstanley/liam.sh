package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gctx "github.com/gorilla/context"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/urfave/cli"
)

const genPost = `
Title: Example Post
Show: false
Time: %s

### Markdown here
`

var tmpl *Loader

type Config struct {
	// Listen changes what ip/port the server is listening on.
	Listen string
	// TLS enables and disables the use of TLS.
	TLS bool
	// CertFile and KeyFile are the system paths to the SSL certificate and
	// key file which can be used when TLS is enabled.
	CertFile string
	KeyFile  string
	// Proxy when enabled, will allow the webserver to sit behind another
	// webserver, that supports X-Forwarded-For and similar headers to provide
	// the server behind the proxy with the real IP address.
	Proxy bool
	// If the templates/style should be cached upon first load.
	Cache bool
	// DebugLog is the log file path where debug logs will be written.
	DebugLog string
}

var conf Config

func runServer() {
	tmpl = NewLoader("static/partials/*", conf.Cache, defaultCtx)
	r := chi.NewRouter()
	if conf.Proxy {
		r.Use(middleware.RealIP)
	}

	r.Use(middleware.DefaultCompress)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: debug}))
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.Recoverer)

	r.FileServer("/static", http.Dir("static"))

	r.Get("/", getPost)
	r.Get("/post/:uid", getPost)

	r.Get("/go/:pkg", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Load(w, r, "static/views/go.html", map[string]interface{}{"pkg": chi.URLParam(r, "pkg")})
	})

	debug.Println("initializing webserver")
	if conf.TLS {
		debug.Fatal(http.ListenAndServeTLS(conf.Listen, mustFile(conf.CertFile), mustFile(conf.KeyFile), gctx.ClearHandler(r)))
	} else {
		debug.Fatal(http.ListenAndServe(conf.Listen, gctx.ClearHandler(r)))
	}
}

func main() {
	initLogger()

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
				cli.StringFlag{Name: "cert", Destination: &conf.CertFile, Usage: "certificate file `path`"},
				cli.StringFlag{Name: "key", Destination: &conf.KeyFile, Usage: "certificate key `path`"},
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

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "log", Destination: &conf.DebugLog, Usage: "Logfile `path`"},
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
