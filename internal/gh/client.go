// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"sync"

	"github.com/google/go-github/v50/github"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	ghql "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const (
	cacheDir = ".gitapicache"
)

var (
	RestClient  *github.Client
	GraphClient *ghql.Client
	clientOnce  sync.Once
	SyncOnStart = false
)

func NewClient(ctx context.Context, token string) {
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
	tc.Transport = &oauth2.Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
		Base:   httpcache.NewTransport(diskcache.New(cacheDir)),
	}

	clientOnce.Do(func() {
		RestClient = github.NewClient(tc)
		GraphClient = ghql.NewClient(tc)
	})
}
