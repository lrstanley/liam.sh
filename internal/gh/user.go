// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gh

import (
	"context"
	"sync/atomic"

	"github.com/apex/log"
	"github.com/google/go-github/v63/github"
)

// User is a cache of the current GitHub user for the authenticated user.
var User atomic.Pointer[github.User]

// UserRunner fetches the current GitHub user for the authenticated user
// and stores it in the User cache, in memory.
func UserRunner(ctx context.Context) error {
	user, _, err := RestClient.Users.Get(ctx, "lrstanley")
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to get user")
		return nil
	}

	log.FromContext(ctx).WithField("user", user.GetLogin()).Info("got user")
	User.Store(user)
	return nil
}
