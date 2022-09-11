// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/v44/github"
)

const userInterval = 10 * time.Minute

var User atomic.Pointer[github.User]

func UserRunner(ctx context.Context) error {
	logger := log.FromContext(ctx).WithField("runner", "github_user")

	getUser(ctx, logger)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(userInterval):
			getUser(ctx, logger)
		}
	}
}

func getUser(ctx context.Context, logger log.Interface) {
	user, _, err := Client.Users.Get(ctx, "lrstanley")
	if err != nil {
		logger.WithError(err).Error("failed to get user")
		return
	}

	logger.WithField("user", user.GetLogin()).Info("got user")
	User.Store(user)
}
