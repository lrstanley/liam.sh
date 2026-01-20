//go:build !skippolicy

// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
)

func (Post) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			AllowAuthenticated(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			AllowAuthenticated(),
			FilterPublicOnly(),
			privacy.AlwaysAllowRule(),
		},
	}
}
