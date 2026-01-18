// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package wakapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/lrstanley/x/http/utils/httpclog"
	"github.com/lrstanley/x/scheduler"
)

const maxSummaryLength = 5

var Statistics atomic.Pointer[models.CodingStats]

func Runner(config models.ConfigWakAPI) scheduler.JobLoggerFunc {
	uri := fmt.Sprintf("%s/api/summary?interval=last_30_days", strings.TrimRight(config.URL, "/"))
	level := slog.LevelInfo

	return scheduler.JobLoggerFunc(func(ctx context.Context, logger *slog.Logger) error {
		client := httpclog.NewClient(&httpclog.Config{
			Level:         &level,
			BaseTransport: http.DefaultTransport,
			Logger:        logger,
		})

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
		if err != nil {
			return err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set(
			"Authorization",
			fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(config.APIKey))),
		)

		resp, err := client.Do(req)
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
	})
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
