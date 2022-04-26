// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package policies

import (
	"context"

	"github.com/lrstanley/liam.sh/internal/ent/privacy"
	"github.com/lrstanley/liam.sh/internal/models"
)

// AllowIfAdmin is a rule that returns Allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if x := ctx.Value(models.UserContextKey); x != nil {
			return privacy.Allow
		}
		return privacy.Deny
	})
}
