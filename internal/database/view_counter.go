// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

var (
	postViewMu = &sync.Mutex{}
	postView   = map[int]map[string]time.Time{}
)

func PostViewCounter(ctx context.Context, post *ent.Post) {
	ip := chix.GetContextIP(ctx)
	if ip == nil {
		return
	}

	nctx, cancel := context.WithTimeout(
		privacy.DecisionContext(context.Background(), privacy.Allow),
		time.Second*5,
	)
	defer cancel()

	postViewMu.Lock()
	if _, ok := postView[post.ID]; !ok {
		postView[post.ID] = map[string]time.Time{}
	}

	if t, ok := postView[post.ID][ip.String()]; ok {
		if time.Since(t) < 5*time.Minute {
			postViewMu.Unlock()
			return
		}
	}
	postView[post.ID][ip.String()] = time.Now()
	postViewMu.Unlock()

	if _, err := post.Update().AddViewCount(1).Save(nctx); err != nil {
		log.FromContext(ctx).WithError(err).WithField("post_id", post.ID).
			Error("failed to update post view count")
	}
}
