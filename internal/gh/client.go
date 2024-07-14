// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"sync"

	"github.com/google/go-github/v52/github"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/lrstanley/liam.sh/internal/models"
	ghql "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var (
	RestClient  *github.Client
	GraphClient *ghql.Client
	clientOnce  sync.Once
	SyncOnStart = false
)

func NewClient(ctx context.Context, config models.ConfigGithub) {
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Token}))
	tc.Transport = &oauth2.Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Token}),
		Base:   httpcache.NewTransport(diskcache.New(config.CachePath)),
	}

	clientOnce.Do(func() {
		RestClient = github.NewClient(tc)
		GraphClient = ghql.NewClient(tc)
	})
}
