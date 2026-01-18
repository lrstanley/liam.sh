// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ai

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/lrstanley/x/http/utils/httpclog"
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
	logger     *slog.Logger
}

func NewService(ctx context.Context, config models.ConfigAI, httpConfig models.ConfigHTTP, db *ent.Client, logger *slog.Logger) (s *service, err error) {
	s = &service{
		config:     config,
		httpConfig: httpConfig,
		db:         db,
		logger:     logger,
	}
	s.cache, err = inmemory.New(ctx, inmemory.WithExpiration(cacheTimeout))
	if err != nil {
		return nil, fmt.Errorf("failed to create cache: %w", err)
	}
	return s, nil
}

func (s *service) newClient(ctx context.Context, opts ...openai.Option) (llms.Model, error) {
	level := slog.LevelInfo

	baseOpts := []openai.Option{
		openai.WithBaseURL(s.config.BaseURL),
		openai.WithToken(s.config.APIKey),
		openai.WithHTTPClient(httpclog.NewClient(&httpclog.Config{
			Level:         &level,
			BaseTransport: http.DefaultTransport,
			Logger:        s.logger,
		})),
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
