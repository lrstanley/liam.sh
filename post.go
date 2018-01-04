// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Depado/bfchroma"
	"github.com/go-chi/chi"
	"github.com/lrstanley/pt"
	zglob "github.com/mattn/go-zglob"
	bf "github.com/russross/blackfriday"
)

type GenPost struct{}

func (GenPost) Execute(_ []string) error {
	fmt.Println(fmt.Sprintf(`
Title: Example Post
Show: false
Time: %s

### Markdown here
`, FormatTime(time.Now())))
	return nil
}

type Post struct {
	Title   string
	Time    time.Time
	Show    bool
	UID     string
	Content []byte
}

// HTML generates an HTML version of the stored markdown.
func (p *Post) HTML() string {
	return string(bf.Run(p.Content, bf.WithRenderer(bfchroma.NewRenderer(bfchroma.Style("monokai")))))
}

// FormatTime is much like stand-alone FormatTime, however it uses the post
// time.
func (p *Post) FormatTime() string {
	return p.Time.Format(time.RFC822)
}

// FormatTime returns a string version of the time format we use for posts.
func FormatTime(in time.Time) string {
	return in.Format(time.RFC822)
}

var uidRe = regexp.MustCompile("^[a-z0-9-._]{3,50}$")

// IsValidUID validates and ensures the UID is "web-safe".
func IsValidUID(uid string) bool {
	return uidRe.MatchString(uid)
}

// ParseTime parses the time format we use in posts from a string.
func ParseTime(in string) (time.Time, error) {
	return time.Parse(time.RFC822, in)
}

// LoadPost loads a given post from a file path.
func LoadPost(path string) (*Post, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0555)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m, err := mail.ReadMessage(f)
	if err != nil {
		return nil, err
	}

	p := &Post{}
	p.Title = m.Header.Get("title")
	if len(p.Title) == 0 {
		return nil, errors.New("no title defined")
	}
	if p.Time, err = ParseTime(m.Header.Get("time")); err != nil {
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

	if !IsValidUID(p.UID) {
		return nil, fmt.Errorf("uid %q is invalid", p.UID)
	}

	// Read the rest of the post.
	if p.Content, err = ioutil.ReadAll(m.Body); err != nil {
		return nil, err
	}

	return p, nil
}

// IsPost ensures that a post with a given uid exists.
func IsPost(uid string) (string, bool) {
	if !IsValidUID(uid) {
		return "", false
	}

	// This would be very unsafe if we were to not use IsValidUID (as someone
	// could potentially use something like "../../../", etc).
	path := "posts/" + uid + ".md"

	if _, err := os.Stat(path); err == nil {
		return path, true
	}

	return "", false
}

// AllPosts gives us all of the public posts.
func AllPosts(path string) []*Post {
	files, err := zglob.Glob(path)
	if err != nil {
		panic(err)
	}

	posts := []*Post{}
	var post *Post

	for i := 0; i < len(files); i++ {
		post, err = LoadPost(files[i])
		if err != nil {
			panic(err)
		}

		// Skip posts we don't want to show.
		if !post.Show {
			continue
		}

		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[j].Time.Before(posts[i].Time)
	})

	return posts
}

// RecentPosts gives us all of the recent public posts.
func RecentPosts(path string, max int) []*Post {
	posts := AllPosts(path)

	var out []*Post
	var count int

	for i := 0; i < len(posts); i++ {
		count++
		out = append(out, posts[i])

		if count == max {
			break
		}
	}

	return out
}

func getPost(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")

	if len(uid) == 0 {
		posts := RecentPosts("posts/**", 5)
		if len(posts) < 1 {
			http.NotFound(w, r)
			return
		}

		tmpl.Render(w, r, "/tmpl/index.html", pt.M{"post": posts[0], "recent": posts})
		return
	}

	path, ok := IsPost(uid)
	if !ok {
		http.NotFound(w, r)
		return
	}

	post, err := LoadPost(path)
	if err != nil {
		panic(err)
	}

	posts := RecentPosts("posts/**", 5)

	tmpl.Render(w, r, "/tmpl/index.html", pt.M{"post": post, "recent": posts})
}
