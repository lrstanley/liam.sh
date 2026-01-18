// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"os"
	"time"

	"github.com/lrstanley/chix/v2"
	"github.com/lrstanley/clix/v2"
	"github.com/lrstanley/liam.sh/internal/ai"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime"
	"github.com/lrstanley/liam.sh/internal/gh"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/lrstanley/liam.sh/internal/wakapi"
	"github.com/lrstanley/x/scheduler"
)

var (
	db    *ent.Client
	aiSvc ai.Service

	cli = clix.NewWithDefaults(
		clix.WithAppInfo[models.Flags](clix.AppInfo{
			Links: clix.GithubLinks("github.com/lrstanley/liam.sh", "master", "https://liam.sh"),
		}),
	)
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	logger := cli.GetLogger()
	ctx := context.Background()

	db = database.Open(ctx, cli.Flags.Database)
	defer db.Close()
	ctx = ent.NewContext(ctx, db)

	database.Migrate(ctx, db)

	gh.NewClient(ctx, cli.Flags.Github)

	var err error
	aiSvc, err = ai.NewService(ctx, cli.Flags.AI, cli.Flags.HTTP, db, logger)
	if err != nil {
		logger.ErrorContext(ctx, "failed to create AI service", "error", err)
	}

	err = chix.Run(
		ctx,
		logger,
		httpServer(logger),
		scheduler.NewCron("users", scheduler.JobLoggerFunc(gh.UserRunner)).
			WithImmediate(true).
			WithLogger(logger).
			WithInterval(10*time.Minute),
		scheduler.NewCron("stats", scheduler.JobLoggerFunc(gh.StatsRunner)).
			WithImmediate(true).
			WithLogger(logger).
			WithInterval(4*time.Hour),
		scheduler.NewCron("events", scheduler.JobLoggerFunc(gh.EventsRunner)).
			WithImmediate(true).
			WithLogger(logger).
			WithInterval(10*time.Minute),
		scheduler.NewCron("repositories", scheduler.JobLoggerFunc(gh.RepositoryRunner)).
			WithLogger(logger).
			WithInterval(30*time.Minute),
		scheduler.NewCron("gists", scheduler.JobLoggerFunc(gh.GistRunner)).
			WithImmediate(true).
			WithLogger(logger).
			WithInterval(120*time.Minute),
		scheduler.NewCron("wakapi", wakapi.Runner(cli.Flags.WakAPI)).
			WithImmediate(true).
			WithLogger(logger).
			WithExitOnError(true).
			WithInterval(30*time.Minute),
	)
	if err != nil {
		logger.ErrorContext(ctx, "shutting down", "error", err)
		exitCode = 1
	}
}
