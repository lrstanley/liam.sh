package schema

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
