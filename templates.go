package main

import (
	"io/ioutil"
	"math"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/aymerick/raymond"
	humanize "github.com/dustin/go-humanize"
	"github.com/go-chi/chi"
)

func init() {
	raymond.RegisterHelper("hrt", func(t time.Time) string {
		if t.IsZero() {
			return ""
		}

		return humanize.Time(t)
	})

	raymond.RegisterHelper("ellipsis", func(max int, text string) string {
		if len(text) > max {
			return text[0:max] + "..."
		}

		return text
	})

	raymond.RegisterHelper("ne", func(a interface{}, b interface{}, options *raymond.Options) interface{} {
		if raymond.Str(a) != raymond.Str(b) {
			return options.Fn()
		}

		return ""
	})

	raymond.RegisterHelper("nicetime", func(t time.Time) string {
		return t.Format(time.ANSIC)
	})

	raymond.RegisterHelper("simpletime", func(t time.Time) string {
		return t.Format("January _2, 2006")
	})
}

func mustFile(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(f)
}

func ListPartials(path string) map[string]string {
	globs, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}

	var j int
	var name string

	out := make(map[string]string)
	for i := 0; i < len(globs); i++ {
		name = filepath.Base(globs[i])
		j = strings.Index(name, ".")
		if j > -1 {
			name = name[0:j]
		}

		out[name] = mustFile(globs[i])
	}

	return out
}

type Loader struct {
	Partials string
	cached   bool

	ctx func(r *http.Request) map[string]interface{}
}

func NewLoader(partials string, cache bool, defaultCtx func(r *http.Request) map[string]interface{}) *Loader {
	if cache {
		raymond.RegisterPartials(ListPartials(partials))

		return &Loader{
			Partials: partials,
			cached:   cache,
			ctx:      defaultCtx,
		}
	}

	return &Loader{Partials: partials, ctx: defaultCtx}
}

func (l *Loader) Get(path string) *raymond.Template {
	if l.cached {
		return raymond.MustParse(mustFile(path))
	}

	tmpl := raymond.MustParse(mustFile(path))
	tmpl.RegisterPartials(ListPartials(l.Partials))

	return tmpl
}

func (l *Loader) Load(w http.ResponseWriter, r *http.Request, path string, ctx interface{}) {
	if l.ctx == nil {
		w.Write([]byte(l.Get(path).MustExec(ctx)))
		return
	}

	out := l.ctx(r)
	if out == nil {
		panic("Loader default context returned nil")
	}

	out["ctx"] = ctx

	w.Write([]byte(l.Get(path).MustExec(out)))
}

type Page struct {
	// N is the page number.
	N int
	// Current is if the page should be active (the requested page).
	Current bool
}

type Paginate struct {
	// Current, Last, and Next pages.
	Current, Last, Next int
	// Total number of pages.
	Total int
	Items int
	// Offset you should use (e.g. for a db).
	Offset int
	// ItemLimit is the is the amount per page.
	ItemLimit int
	// Padding of pages on left and right side.
	Padding int

	Pages []Page
}

func newPaginate(request int, limit int, items int, pad int) *Paginate {
	out := &Paginate{Current: request, Items: items, ItemLimit: limit, Padding: pad}

	if out.Padding < 1 {
		out.Padding = 1
	}
	if out.ItemLimit < 2 {
		out.ItemLimit = 2
	}

	out.Total = int(math.Ceil(float64(out.Items) / float64(out.ItemLimit)))
	if out.Total < 1 {
		out.Total = 1
	}

	if out.Current == 1 {
		out.Offset = 0
	} else {
		out.Offset = (out.Current - 1) * limit
	}

	out.Last = out.Current - 1
	out.Next = out.Current + 1

	if out.Current < 1 {
		out.Current = 1
	}
	if out.Current > out.Total {
		out.Current = out.Total
	}
	if out.Last < 1 {
		out.Last = 1
	}
	if out.Next > out.Total {
		out.Next = out.Total
	}

	pmin := 1
	pmax := 0
	pmin = request - out.Padding
	if pmin < 1 {
		pmax = int(math.Abs(float64(pmin))) + 1
		pmin = 1
	}

	pmax = request + out.Padding + pmax

	out.Pages = []Page{}

	for p := 1; p <= out.Total; p++ {
		if p == out.Current {
			// It's the current page, mark as current.
			out.Pages = append(out.Pages, Page{N: p, Current: true})
			continue
		}

		if p >= pmin && p <= pmax ||
			// If it is only 1 away from first page.
			(p == 1 && out.Current-out.Padding-1 == 1) ||
			// If it is only 1 away from last page.
			(p == out.Total && p+out.Padding+1 == out.Total) {

			out.Pages = append(out.Pages, Page{N: p, Current: false})
			continue
		}
	}

	return out
}

func defaultCtx(r *http.Request) map[string]interface{} {
	return map[string]interface{}{
		"full_url": r.URL.String(),
		"url":      r.URL.Path,
	}
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
