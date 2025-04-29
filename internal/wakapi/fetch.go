// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package wakapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/models"
)

const (
	fetchInterval    = 30 * time.Minute
	maxSummaryLength = 5
)

var Statistics atomic.Pointer[models.CodingStats]

func NewRunner(logger log.Interface, config models.ConfigWakAPI) *Runner {
	return &Runner{
		config: config,
		logger: logger.WithField("runner", "wakapi"),
		client: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

type Runner struct {
	config models.ConfigWakAPI
	logger log.Interface
	client http.Client
}

func (r *Runner) Run(ctx context.Context) (err error) {
	if err = r.fetch(ctx); err != nil {
		r.logger.WithError(err).Error("failed to fetch wakapi stats")
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(fetchInterval):
			if err = r.fetch(ctx); err != nil {
				r.logger.WithError(err).Error("failed to fetch wakapi stats")
				return err
			}
		}
	}
}

func (r *Runner) url() string {
	return fmt.Sprintf("%s/api/summary?interval=last_30_days", strings.TrimRight(r.config.URL, "/"))
}

func (r *Runner) fetch(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.url(), http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(r.config.APIKey))),
	)

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var stats models.CodingStats
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		return err
	}

	// Calculate misc stats for the frontend.
	for i := range stats.Languages {
		stats.Languages[i].HexColor = lookupColor(stats.Languages[i].Language)
		stats.Languages[i].TotalDuration = formatDuration(stats.Languages[i].TotalSeconds)
		stats.TotalSeconds += stats.Languages[i].TotalSeconds
	}
	stats.TotalDuration = formatDuration(stats.TotalSeconds)
	stats.TotalDurationShort = formatDurationShort(stats.TotalSeconds)
	stats.CalculatedDays = 30

	maxTitleLength := 5

	for _, stat := range stats.Languages {
		if len(stats.Summary) > maxSummaryLength {
			stats.Summary[maxSummaryLength].Key = "Other"
			stats.Summary[maxSummaryLength].HexColor = stat.HexColor
			stats.Summary[maxSummaryLength].TotalSeconds = stat.TotalSeconds
			continue
		}

		maxTitleLength = max(maxTitleLength, len(stat.Language))
		stats.Summary = append(stats.Summary, &models.CodingStatsSummary{
			Key:           stat.Language,
			HexColor:      stat.HexColor,
			TotalSeconds:  stat.TotalSeconds,
			TotalDuration: stat.TotalDuration,
			TitleLength:   len(stat.Language),
		})
	}

	for _, stat := range stats.Summary {
		stat.Percentage = int(math.Round(float64(stat.TotalSeconds) / float64(stats.TotalSeconds) * 100))
	}

	Statistics.Store(&stats)
	return nil
}

func formatDuration(seconds int) (out string) {
	hours := seconds / 3600
	minutes := (seconds / 60) % 60

	if hours < 1 && minutes < 1 {
		minutes = 1
	}

	return fmt.Sprintf("%d hrs %d mins", hours, minutes)
}

func formatDurationShort(seconds int) (out string) {
	hours := seconds / 3600
	minutes := (seconds / 60) % 60

	if hours < 1 {
		return fmt.Sprintf("%d mins", minutes)
	}

	return fmt.Sprintf("%d hrs", hours)
}
