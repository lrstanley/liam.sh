// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apex/log"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/ent"
	_ "github.com/lrstanley/liam.sh/internal/ent/runtime"
	"github.com/lrstanley/liam.sh/internal/models"
	"golang.org/x/sync/errgroup"
)

var (
	version = "master"
	commit  = "latest"
	date    = "-"
	cli     = &models.Flags{}

	db     *ent.Client
	logger log.Interface
)

func main() {
	_ = cli.Parse()
	logger = cli.Log.Parse(cli.Debug).WithFields(log.Fields{
		"build_version": fmt.Sprintf("%s/%s (%s)", version, commit, date),
	})

	db = database.Open(logger, cli.Database)
	defer db.Close()
	database.Migrate(context.TODO(), logger, db)

	grp, ctx := errgroup.WithContext(context.Background())

	grp.Go(func() error {
		return httpServer(ctx)
	})
	grp.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case sig := <-quit:
			return fmt.Errorf("received signal %v", sig)
		case <-ctx.Done():
			return nil
		}
	})

	if err := grp.Wait(); err != nil {
		logger.WithError(err).Fatal("shutting down")
	}
}
