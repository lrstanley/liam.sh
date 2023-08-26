// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"

	"github.com/apex/log"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/models"
)

var (
	logger log.Interface

	cli = &clix.CLI[models.FlagsMigration]{}
)

func main() {
	cli.Parse()
	logger = cli.Logger

	ctx := log.NewContext(context.Background(), logger)

	if len(cli.Args) < 1 {
		logger.Fatal("no migration name specified")
	}

	database.GenerateMigrations(ctx, logger, cli.Flags.Database, cli.Flags.MigrationPath, cli.Args[0])

	// db = database.Open(context.Background(), logger, cli.Flags.Database)
	// defer db.Close()

	// database.RegisterHooks(ctx)
	// database.Migrate(ctx, logger)

	// gh.SyncOnStart = cli.Flags.Github.SyncOnStart
	// gh.NewClient(ctx, cli.Flags.Github.Token)

	// if err := chix.RunCtx(
	// 	ctx, httpServer(ctx),
	// 	gh.UserRunner,
	// 	gh.StatsRunner,
	// 	gh.EventsRunner,
	// 	gh.RepositoryRunner,
	// 	gh.GistRunner,
	// 	wakapi.NewRunner(logger, cli.Flags.WakAPI).Run,
	// ); err != nil {
	// 	logger.WithError(err).Fatal("shutting down")
	// }
}
