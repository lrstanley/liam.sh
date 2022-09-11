// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"context"
	"strings"

	"entgo.io/ent/entql"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

// AllowRoles is a rule that returns Allow decision if the authenticated client
// has ONE of the specified roles.
func AllowRoles(allowed []string, deny bool) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		roles := chix.RolesFromContext(ctx)
		if roles != nil {
			for _, role := range roles {
				for _, allowedRole := range allowed {
					if strings.EqualFold(role, allowedRole) {
						return privacy.Allow
					}
				}
			}
		}

		if deny {
			return privacy.Deny
		}
		return privacy.Skip
	})
}

func AllowPublicOnly() privacy.QueryRule {
	type PublicFilter interface {
		WherePublic(p entql.BoolP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		public, ok := f.(PublicFilter)
		if !ok {
			return privacy.Denyf("missing public field in filter")
		}

		public.WherePublic(entql.BoolEQ(true))
		return privacy.Skip
	})
}
