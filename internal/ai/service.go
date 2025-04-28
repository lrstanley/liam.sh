// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ai

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/cache"
	"github.com/tmc/langchaingo/llms/cache/inmemory"
	"github.com/tmc/langchaingo/llms/openai"
)

var _ Service = (*service)(nil)

type Service interface {
	PostSuggestLabels(ctx context.Context, content string) ([]SuggestedLabel, error)
}

const cacheTimeout = 30 * time.Minute

type service struct {
	config     models.ConfigAI
	httpConfig models.ConfigHTTP
	cache      *inmemory.InMemory
	db         *ent.Client
}

func NewService(ctx context.Context, config models.ConfigAI, httpConfig models.ConfigHTTP, db *ent.Client) (s *service, err error) {
	s = &service{
		config:     config,
		httpConfig: httpConfig,
		db:         db,
	}
	s.cache, err = inmemory.New(ctx, inmemory.WithExpiration(cacheTimeout))
	if err != nil {
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}
	return s, nil
}

func (s *service) newClient(ctx context.Context, opts ...openai.Option) (llms.Model, error) {
	baseOpts := []openai.Option{
		openai.WithBaseURL(s.config.BaseURL),
		openai.WithToken(s.config.APIKey),
		openai.WithHTTPClient(&http.Client{
			Timeout: 60 * time.Second,
			Transport: &httpTransport{
				headers: map[string]string{
					"HTTP-Referer": s.httpConfig.BaseURL,
					"X-Title":      "Liam Stanley's Personal Website & Blog",
				},
				logger: log.FromContext(ctx).WithField("src", "ai-client"),
			},
		}),
	}

	if s.config.DefaultModel != "" {
		baseOpts = append(baseOpts, openai.WithModel(s.config.DefaultModel))
	}

	client, err := openai.New(append(baseOpts, opts...)...)
	if err != nil {
		return nil, fmt.Errorf("failed to create AI client: %w", err)
	}

	return cache.New(client, s.cache), nil
}
