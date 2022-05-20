// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"

	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/hook"
	"github.com/lrstanley/liam.sh/internal/ent/post"
	"github.com/lrstanley/liam.sh/internal/markdown"
)

// RegisterHooks initializes all runtime hooks into the database client.
func RegisterHooks(ctx context.Context, db *ent.Client) {
	log.FromContext(ctx).Info("registering database hooks")
	db.Use(
		hook.On(
			hook.If(func(next ent.Mutator) ent.Mutator {
				return hook.PostFunc(func(ctx context.Context, m *ent.PostMutation) (ent.Value, error) {
					content, ok := m.Content()
					if !ok {
						return next.Mutate(ctx, m)
					}

					md, err := markdown.Generate(ctx, content)
					if err != nil {
						return m, err
					}

					m.SetContentHTML(md)

					return next.Mutate(ctx, m)
				})
			}, hook.HasFields(post.FieldContent)),
			ent.OpCreate|ent.OpUpdateOne|ent.OpUpdate,
		),
	)
}
