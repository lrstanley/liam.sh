//go:build !skippolicy

// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/lrstanley/chix/xauth/v2"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

func AllowAuthenticated() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if xauth.IdentFromContext[ent.User](ctx) != nil {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

func FilterPublicOnly() privacy.QueryRule {
	type PublicFilter interface {
		WherePublic(p entql.BoolP)
	}
	return privacy.FilterFunc(func(_ context.Context, f privacy.Filter) error {
		public, ok := f.(PublicFilter)
		if !ok {
			return privacy.Denyf("missing public field in filter")
		}

		public.WherePublic(entql.BoolEQ(true))
		return privacy.Skip
	})
}
