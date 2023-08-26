// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/clix"
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

	cli = &clix.CLI[models.Flags]{
		Links: clix.GithubLinks("github.com/lrstanley/liam.sh", "master", "https://liam.sh"),
	}
)

func main() {
	cli.Parse()
	logger = cli.Logger

	db = database.Open(context.Background(), logger, cli.Flags.Database)
	defer db.Close()

	ctx := ent.NewContext(log.NewContext(context.Background(), logger), db)

	database.RegisterHooks(ctx)

	if cli.Flags.Database.AutoMigrate {
		database.Migrate(ctx, logger)
	}

	gh.SyncOnStart = cli.Flags.Github.SyncOnStart
	gh.NewClient(ctx, cli.Flags.Github.Token)

	if err := chix.RunCtx(
		ctx, httpServer(ctx),
		gh.UserRunner,
		gh.StatsRunner,
		gh.EventsRunner,
		gh.RepositoryRunner,
		gh.GistRunner,
		wakapi.NewRunner(logger, cli.Flags.WakAPI).Run,
	); err != nil {
		logger.WithError(err).Fatal("shutting down")
	}
}
