// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/google/go-github/v84/github"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/lrstanley/x/http/utils/httpccache"
	"github.com/lrstanley/x/http/utils/httpclog"
	ghql "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var (
	RestClient  *github.Client
	GraphClient *ghql.Client
	clientOnce  sync.Once
)

// NewClient returns a new pre-configured GitHub client.
func NewClient(ctx context.Context, logger *slog.Logger, config models.ConfigGithub) {
	storage, err := httpccache.NewFileStorage(config.CachePath, 250, 24*time.Hour)
	if err != nil {
		panic(err)
	}

	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Token}))
	tc.Transport = &oauth2.Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Token}),
		Base: httpccache.NewTransport(&httpccache.Config{
			BaseTransport: httpclog.NewTransport(&httpclog.Config{
				BaseTransport: http.DefaultTransport,
				Logger:        logger,
			}),
			Storage:                   storage,
			Logger:                    logger,
			AllowAuthorizationCaching: true,
			MaxObjectSize:             1024 * 1024 * 10, // 10MB
			AllowHeuristicFreshness:   true,
		}),
	}

	clientOnce.Do(func() {
		RestClient = github.NewClient(tc)
		GraphClient = ghql.NewClient(tc)
	})
}
