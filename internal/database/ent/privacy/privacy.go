// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package privacy

import (
	"context"

	"github.com/lrstanley/liam.sh/internal/database/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The GithubAssetQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubAssetQueryRuleFunc func(context.Context, *ent.GithubAssetQuery) error

// EvalQuery return f(ctx, q).
func (f GithubAssetQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubAssetQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubAssetQuery", q)
}

// The GithubAssetMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubAssetMutationRuleFunc func(context.Context, *ent.GithubAssetMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubAssetMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubAssetMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubAssetMutation", m)
}

// The GithubEventQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubEventQueryRuleFunc func(context.Context, *ent.GithubEventQuery) error

// EvalQuery return f(ctx, q).
func (f GithubEventQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubEventQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubEventQuery", q)
}

// The GithubEventMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubEventMutationRuleFunc func(context.Context, *ent.GithubEventMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubEventMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubEventMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubEventMutation", m)
}

// The GithubGistQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubGistQueryRuleFunc func(context.Context, *ent.GithubGistQuery) error

// EvalQuery return f(ctx, q).
func (f GithubGistQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubGistQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubGistQuery", q)
}

// The GithubGistMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubGistMutationRuleFunc func(context.Context, *ent.GithubGistMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubGistMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubGistMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubGistMutation", m)
}

// The GithubReleaseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubReleaseQueryRuleFunc func(context.Context, *ent.GithubReleaseQuery) error

// EvalQuery return f(ctx, q).
func (f GithubReleaseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubReleaseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubReleaseQuery", q)
}

// The GithubReleaseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubReleaseMutationRuleFunc func(context.Context, *ent.GithubReleaseMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubReleaseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubReleaseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubReleaseMutation", m)
}

// The GithubRepositoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubRepositoryQueryRuleFunc func(context.Context, *ent.GithubRepositoryQuery) error

// EvalQuery return f(ctx, q).
func (f GithubRepositoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubRepositoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubRepositoryQuery", q)
}

// The GithubRepositoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubRepositoryMutationRuleFunc func(context.Context, *ent.GithubRepositoryMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubRepositoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubRepositoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubRepositoryMutation", m)
}

// The LabelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LabelQueryRuleFunc func(context.Context, *ent.LabelQuery) error

// EvalQuery return f(ctx, q).
func (f LabelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LabelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LabelQuery", q)
}

// The LabelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LabelMutationRuleFunc func(context.Context, *ent.LabelMutation) error

// EvalMutation calls f(ctx, m).
func (f LabelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LabelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LabelMutation", m)
}

// The PostQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PostQueryRuleFunc func(context.Context, *ent.PostQuery) error

// EvalQuery return f(ctx, q).
func (f PostQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PostQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PostQuery", q)
}

// The PostMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PostMutationRuleFunc func(context.Context, *ent.PostMutation) error

// EvalMutation calls f(ctx, m).
func (f PostMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PostMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PostMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.GithubAssetQuery:
		return q.Filter(), nil
	case *ent.GithubEventQuery:
		return q.Filter(), nil
	case *ent.GithubGistQuery:
		return q.Filter(), nil
	case *ent.GithubReleaseQuery:
		return q.Filter(), nil
	case *ent.GithubRepositoryQuery:
		return q.Filter(), nil
	case *ent.LabelQuery:
		return q.Filter(), nil
	case *ent.PostQuery:
		return q.Filter(), nil
	case *ent.UserQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.GithubAssetMutation:
		return m.Filter(), nil
	case *ent.GithubEventMutation:
		return m.Filter(), nil
	case *ent.GithubGistMutation:
		return m.Filter(), nil
	case *ent.GithubReleaseMutation:
		return m.Filter(), nil
	case *ent.GithubRepositoryMutation:
		return m.Filter(), nil
	case *ent.LabelMutation:
		return m.Filter(), nil
	case *ent.PostMutation:
		return m.Filter(), nil
	case *ent.UserMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
