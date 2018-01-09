// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Depado/bfchroma"
	"github.com/go-chi/chi"
	"github.com/lrstanley/pt"
	zglob "github.com/mattn/go-zglob"
	bf "gopkg.in/russross/blackfriday.v2"
)

type GenPost struct{}

func (GenPost) Execute(_ []string) error {
	fmt.Println(fmt.Sprintf(`
Title: Example Post
Show: false
Time: %s

### Markdown here
`, formatPostTime(time.Now())))
	return nil
}

var pc = &postCache{posts: make(map[string]*Post)}

type postCache struct {
	sync.RWMutex
	posts map[string]*Post
}

func (p *postCache) update(path string) {
	defer log.Print("update of post index cache complete")
	posts := make(map[string]*Post)

	files, err := zglob.Glob(path)
	if err != nil {
		log.Printf("error indexing posts: %s", err)
		return
	}

	var post *Post

	for i := 0; i < len(files); i++ {
		log.Printf("updating index for: %s", files[i])
		post, err = loadPost(files[i])
		if err != nil {
			log.Printf("error loading post %s: %s", files[i], err)
		}

		// Skip posts we don't want to show.
		if !post.Show {
			continue
		}

		posts[post.UID] = post
	}

	p.Lock()
	p.posts = posts
	p.Unlock()
}

func (p *postCache) exists(uid string) bool {
	if !isValidPostUID(uid) {
		return false
	}

	p.RLock()
	_, ok := p.posts[uid]
	p.RUnlock()
	return ok
}

func (p *postCache) get(uid string) *Post {
	p.RLock()
	post := p.posts[uid]
	p.RUnlock()
	return post
}

func (p *postCache) all() (posts []*Post) {
	p.RLock()
	posts = make([]*Post, 0, len(p.posts))
	for uid := range p.posts {
		posts = append(posts, p.posts[uid])
	}
	p.RUnlock()

	return posts
}

func (p *postCache) recent(max int) []*Post {
	posts := []*Post{}
	var count int

	all := p.all()
	sortPosts(all)

	for i := 0; i < len(all); i++ {
		count++
		posts = append(posts, all[i])

		if count == max {
			break
		}
	}

	return posts
}

type Post struct {
	Title   string
	Summary string
	Time    time.Time
	Show    bool
	UID     string
	Content []byte

	TOCCache  atomic.Value
	HTMLCache atomic.Value
}

var tocSplitAt = []byte("</nav>")
var tocReplacer = strings.NewReplacer(
	"<nav>", "",
	"<ul>", `<ul class="nav flex-column">`,
	"<li>", `<li class="nav-item">`,
	"<a ", `<a class="nav-link" `,
)

// HTML generates an HTML version of the stored markdown.
func (p *Post) HTML() (toc, body string) {
	if tocCache := p.TOCCache.Load(); tocCache != nil {
		toc = tocCache.(string)
	}
	if htmlCache := p.HTMLCache.Load(); htmlCache != nil {
		body = htmlCache.(string)
	}
	if toc != "" || body != "" {
		log.Printf("loaded post from generated cache")
		return toc, body
	}

	raw := bf.Run(
		p.Content, bf.WithRenderer(bfchroma.NewRenderer(
			bfchroma.Style("monokai"),
			bfchroma.Extend(bf.NewHTMLRenderer(bf.HTMLRendererParameters{
				Flags: bf.CommonHTMLFlags | bf.TOC,
			})),
		)),
	)

	if i := bytes.Index(raw, tocSplitAt); i > -1 {
		toc = tocReplacer.Replace(string(raw[:i]))
		body = string(raw[i+len(tocSplitAt)+1:])

		p.TOCCache.Store(toc)
		p.HTMLCache.Store(body)

		return toc, body
	}

	body = string(raw)
	p.HTMLCache.Store(body)
	return "", body
}

// FormatTime is much like stand-alone FormatTime, however it uses the post
// time.
func (p *Post) FormatTime() string {
	return p.Time.Format(time.RFC822)
}

func sortPosts(posts []*Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[j].Time.Before(posts[i].Time)
	})
}

// formatPostTime returns a string version of the time format we use for posts.
func formatPostTime(in time.Time) string {
	return in.Format(time.RFC822)
}

var uidRe = regexp.MustCompile("^[a-z0-9-._]{3,50}$")

// isValidPostUID validates and ensures the UID is "web-safe".
func isValidPostUID(uid string) bool {
	return uidRe.MatchString(uid)
}

// parsePostTime parses the time format we use in posts from a string.
func parsePostTime(in string) (time.Time, error) {
	return time.Parse(time.RFC822, in)
}

// loadPost loads a given post from a file path.
func loadPost(path string) (*Post, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0555)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m, err := mail.ReadMessage(f)
	if err != nil {
		return nil, err
	}

	p := &Post{
		Title:   m.Header.Get("title"),
		Summary: m.Header.Get("Summary"),
	}
	if len(p.Title) == 0 {
		return nil, errors.New("no title defined")
	}
	if p.Time, err = parsePostTime(m.Header.Get("time")); err != nil {
		return nil, err
	}
	if p.Show, err = strconv.ParseBool(m.Header.Get("show")); err != nil {
		return nil, err
	}

	// Now the uid.
	base := filepath.Base(path)
	j := strings.Index(base, ".")
	if j < 3 {
		return nil, fmt.Errorf("uid of filename %q invalid", path)
	}
	p.UID = base[0:j]

	if !isValidPostUID(p.UID) {
		return nil, fmt.Errorf("uid %q is invalid", p.UID)
	}

	// Read the rest of the post.
	if p.Content, err = ioutil.ReadAll(m.Body); err != nil {
		return nil, err
	}

	return p, nil
}

func getPost(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")

	if len(uid) == 0 {
		posts := pc.recent(10)
		if len(posts) < 1 {
			notFoundHandler(w, r)
			return
		}

		tmpl.Render(w, r, "/tmpl/index.html", pt.M{"recent": posts})
		return
	}

	if !pc.exists(uid) {
		notFoundHandler(w, r)
		return
	}

	post := pc.get(uid)
	recent := pc.recent(5)

	toc, body := post.HTML()
	tmpl.Render(w, r, "/tmpl/post.html", pt.M{
		"toc": toc, "body": body,
		"post": post, "recent": recent,
	})
}
