package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/flosch/pongo2"
	_ "github.com/flosch/pongo2-addons"
	"github.com/go-chi/chi"
)

func mustFile(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(f)
}

func tmpl(w http.ResponseWriter, r *http.Request, path string, ctx map[string]interface{}) {
	var tpl *pongo2.Template

	if conf.Cache {
		tpl = pongo2.Must(pongo2.FromCache(path))
	} else {
		tpl = pongo2.Must(pongo2.FromFile(path))
	}

	out, err := tpl.ExecuteBytes(ctx)
	if err != nil {
		panic(err)
	}

	w.Write(out)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
