package schema

import (
	"context"
	"strings"

	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
)

// import (
// 	"entgo.io/ent"
// 	"entgo.io/ent/schema/mixin"
// 	"github.com/lrstanley/liam.sh/internal/ent/privacy"
// )

// // PrivacyMixin for all schemas in the graph.
// type PrivacyMixin struct {
// 	mixin.Schema
// }

// // Policy defines the privacy policy of the PrivacyMixin.
// func (PrivacyMixin) Policy() ent.Policy {
// 	return privacy.Policy{
// 		Mutation: privacy.MutationPolicy{
// 			// policies.AllowIfAdmin(),
// 			privacy.AlwaysAllowRule(),
// 		},
// 		Query: privacy.QueryPolicy{
// 			privacy.AlwaysAllowRule(),
// 		},
// 	}
// }

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
