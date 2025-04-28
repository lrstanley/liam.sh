// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/liam.sh/internal/ai"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime"
	"github.com/lrstanley/liam.sh/internal/gh"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/lrstanley/liam.sh/internal/wakapi"
)

var (
	db     *ent.Client
	logger log.Interface
	aiSvc  ai.Service

	cli = &clix.CLI[models.Flags]{
		Links: clix.GithubLinks("github.com/lrstanley/liam.sh", "master", "https://liam.sh"),
	}
)

func main() {
	cli.Parse()
	logger = cli.Logger

	ctx := log.NewContext(context.Background(), logger)

	db = database.Open(ctx, cli.Flags.Database)
	defer db.Close()
	ctx = ent.NewContext(ctx, db)

	database.Migrate(ctx, db)

	gh.NewClient(ctx, cli.Flags.Github)

	var err error
	aiSvc, err = ai.NewService(ctx, cli.Flags.AI, cli.Flags.HTTP, db)
	if err != nil {
		logger.WithError(err).Fatal("failed to create AI service")
	}

	if err := chix.RunContext(
		ctx, httpServer(ctx),
		chix.RunnerInterval("users", gh.UserRunner, 10*time.Minute, true, false),
		chix.RunnerInterval("stats", gh.StatsRunner, 4*time.Hour, true, false),
		chix.RunnerInterval("events", gh.EventsRunner, 30*time.Minute, cli.Flags.Github.SyncOnStart, false),
		chix.RunnerInterval("repositories", gh.RepositoryRunner, 30*time.Minute, cli.Flags.Github.SyncOnStart, false),
		chix.RunnerInterval("gists", gh.GistRunner, 120*time.Minute, cli.Flags.Github.SyncOnStart, false),
		chix.RunnerInterval("wakapi", wakapi.NewRunner(logger, cli.Flags.WakAPI).Run, 30*time.Minute, true, true),
	); err != nil {
		logger.WithError(err).Fatal("shutting down")
	}
}
