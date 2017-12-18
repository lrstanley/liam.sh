// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/go-chi/chi"
	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/lrstanley/pt"
	"golang.org/x/oauth2"
)

func newGit() (*github.Client, context.Context) {
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")}))
	tc.Transport = &oauth2.Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")}),
		Base:   httpcache.NewTransport(diskcache.New(".gitapicache")),
	}

	return github.NewClient(tc), ctx
}

type gitCache struct {
	sync.RWMutex

	repos    []*github.Repository
	assetMap map[string]string
}

var gc = &gitCache{}

func updateReleaseCache() error {
	debug.Println("updating github release cache")
	git, ctx := newGit()
	repositories := []*github.Repository{}
	assetMap := map[string]string{}

	debug.Println("fetching repo list")
	var repoPage int
	for {
		repos, resp, err := git.Repositories.List(
			ctx, os.Getenv("GITHUB_USER"), &github.RepositoryListOptions{
				ListOptions: github.ListOptions{Page: repoPage}, Sort: "updated",
			},
		)
		if err != nil {
			return err
		}

		repositories = append(repositories, repos...)

		if resp.NextPage == 0 {
			break
		}
		repoPage = resp.NextPage
	}

	for _, repo := range repositories {
		releases := []*github.RepositoryRelease{}

		debug.Printf("fetching release info for %s", repo.GetFullName())
		var relPage int
		for {
			rel, resp, err := git.Repositories.ListReleases(
				ctx, os.Getenv("GITHUB_USER"), *repo.Name, &github.ListOptions{Page: relPage},
			)
			if err != nil {
				return err
			}

			releases = append(releases, rel...)

			if resp.NextPage == 0 {
				break
			}
			relPage = resp.NextPage
		}

		assets := []*github.ReleaseAsset{}
		relPage = 0
		for _, release := range releases {
			debug.Printf("fetching asset info for %s/%s/%s", repo.GetFullName(), release.GetTagName(), release.GetName())

			for {
				assetList, resp, err := git.Repositories.ListReleaseAssets(
					ctx, os.Getenv("GITHUB_USER"), *repo.Name, *release.ID,
					&github.ListOptions{Page: relPage},
				)
				if err != nil {
					return err
				}

				assets = append(assets, assetList...)
				if resp.NextPage == 0 {
					break
				}
				relPage = resp.NextPage
			}
		}

		for _, asset := range assets {
			assetMap[asset.GetName()] = asset.GetBrowserDownloadURL()
		}
	}

	gc.Lock()
	gc.repos = repositories
	gc.assetMap = assetMap
	gc.Unlock()

	debug.Println("cache update complete")

	return nil
}

func gitRequestAsset(w http.ResponseWriter, r *http.Request) {
	gc.RLock()
	asset, ok := gc.assetMap[chi.URLParam(r, "asset")]
	gc.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, asset, http.StatusFound)
}

func gitRequestAllAssets(w http.ResponseWriter, r *http.Request) {
	gc.RLock()
	mapRef := make(map[string]string, len(gc.assetMap))
	keys := make([]string, 0, len(mapRef))
	for key := range gc.assetMap {
		mapRef[key] = gc.assetMap[key]
		keys = append(keys, key)
	}
	gc.RUnlock()

	if strings.HasSuffix(strings.ToLower(r.URL.Path), ".json") {
		pt.JSON(w, r, mapRef)
		return
	}

	sort.Strings(keys)
	tmpl.Render(w, r, "/tmpl/ghr.html", pt.M{"assets": mapRef, "assetsSorted": keys})
}
