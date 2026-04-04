// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"log/slog"
	"sync/atomic"

	"github.com/google/go-github/v82/github"
)

// User is a cache of the current GitHub user for the authenticated user.
var User atomic.Pointer[github.User]

// UserRunner fetches the current GitHub user for the authenticated user
// and stores it in the User cache, in memory.
func UserRunner(ctx context.Context, logger *slog.Logger) error {
	user, _, err := RestClient.Users.Get(ctx, "lrstanley")
	if err != nil {
		logger.ErrorContext(ctx, "failed to get user", "error", err)
		return nil
	}

	logger.InfoContext(ctx, "got user", "user", user.GetLogin())
	User.Store(user)
	return nil
}
