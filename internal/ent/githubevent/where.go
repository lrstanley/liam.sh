// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package githubevent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLTE(FieldID, id))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldEventID, v))
}

// EventType applies equality check predicate on the "event_type" field. It's identical to EventTypeEQ.
func EventType(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldEventType, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldCreatedAt, v))
}

// Public applies equality check predicate on the "public" field. It's identical to PublicEQ.
func Public(v bool) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldPublic, v))
}

// ActorID applies equality check predicate on the "actor_id" field. It's identical to ActorIDEQ.
func ActorID(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldActorID, v))
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldRepoID, v))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNotIn(FieldEventID, vs...))
}

// EventIDGT applies the GT predicate on the "event_id" field.
func EventIDGT(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGT(FieldEventID, v))
}

// EventIDGTE applies the GTE predicate on the "event_id" field.
func EventIDGTE(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGTE(FieldEventID, v))
}

// EventIDLT applies the LT predicate on the "event_id" field.
func EventIDLT(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLT(FieldEventID, v))
}

// EventIDLTE applies the LTE predicate on the "event_id" field.
func EventIDLTE(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLTE(FieldEventID, v))
}

// EventIDContains applies the Contains predicate on the "event_id" field.
func EventIDContains(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldContains(FieldEventID, v))
}

// EventIDHasPrefix applies the HasPrefix predicate on the "event_id" field.
func EventIDHasPrefix(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldHasPrefix(FieldEventID, v))
}

// EventIDHasSuffix applies the HasSuffix predicate on the "event_id" field.
func EventIDHasSuffix(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldHasSuffix(FieldEventID, v))
}

// EventIDEqualFold applies the EqualFold predicate on the "event_id" field.
func EventIDEqualFold(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEqualFold(FieldEventID, v))
}

// EventIDContainsFold applies the ContainsFold predicate on the "event_id" field.
func EventIDContainsFold(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldContainsFold(FieldEventID, v))
}

// EventTypeEQ applies the EQ predicate on the "event_type" field.
func EventTypeEQ(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldEventType, v))
}

// EventTypeNEQ applies the NEQ predicate on the "event_type" field.
func EventTypeNEQ(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNEQ(FieldEventType, v))
}

// EventTypeIn applies the In predicate on the "event_type" field.
func EventTypeIn(vs ...string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldIn(FieldEventType, vs...))
}

// EventTypeNotIn applies the NotIn predicate on the "event_type" field.
func EventTypeNotIn(vs ...string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNotIn(FieldEventType, vs...))
}

// EventTypeGT applies the GT predicate on the "event_type" field.
func EventTypeGT(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGT(FieldEventType, v))
}

// EventTypeGTE applies the GTE predicate on the "event_type" field.
func EventTypeGTE(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGTE(FieldEventType, v))
}

// EventTypeLT applies the LT predicate on the "event_type" field.
func EventTypeLT(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLT(FieldEventType, v))
}

// EventTypeLTE applies the LTE predicate on the "event_type" field.
func EventTypeLTE(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLTE(FieldEventType, v))
}

// EventTypeContains applies the Contains predicate on the "event_type" field.
func EventTypeContains(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldContains(FieldEventType, v))
}

// EventTypeHasPrefix applies the HasPrefix predicate on the "event_type" field.
func EventTypeHasPrefix(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldHasPrefix(FieldEventType, v))
}

// EventTypeHasSuffix applies the HasSuffix predicate on the "event_type" field.
func EventTypeHasSuffix(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldHasSuffix(FieldEventType, v))
}

// EventTypeEqualFold applies the EqualFold predicate on the "event_type" field.
func EventTypeEqualFold(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEqualFold(FieldEventType, v))
}

// EventTypeContainsFold applies the ContainsFold predicate on the "event_type" field.
func EventTypeContainsFold(v string) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldContainsFold(FieldEventType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLTE(FieldCreatedAt, v))
}

// PublicEQ applies the EQ predicate on the "public" field.
func PublicEQ(v bool) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldPublic, v))
}

// PublicNEQ applies the NEQ predicate on the "public" field.
func PublicNEQ(v bool) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNEQ(FieldPublic, v))
}

// ActorIDEQ applies the EQ predicate on the "actor_id" field.
func ActorIDEQ(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldActorID, v))
}

// ActorIDNEQ applies the NEQ predicate on the "actor_id" field.
func ActorIDNEQ(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNEQ(FieldActorID, v))
}

// ActorIDIn applies the In predicate on the "actor_id" field.
func ActorIDIn(vs ...int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldIn(FieldActorID, vs...))
}

// ActorIDNotIn applies the NotIn predicate on the "actor_id" field.
func ActorIDNotIn(vs ...int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNotIn(FieldActorID, vs...))
}

// ActorIDGT applies the GT predicate on the "actor_id" field.
func ActorIDGT(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGT(FieldActorID, v))
}

// ActorIDGTE applies the GTE predicate on the "actor_id" field.
func ActorIDGTE(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGTE(FieldActorID, v))
}

// ActorIDLT applies the LT predicate on the "actor_id" field.
func ActorIDLT(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLT(FieldActorID, v))
}

// ActorIDLTE applies the LTE predicate on the "actor_id" field.
func ActorIDLTE(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLTE(FieldActorID, v))
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldEQ(FieldRepoID, v))
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNEQ(FieldRepoID, v))
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldIn(FieldRepoID, vs...))
}

// RepoIDNotIn applies the NotIn predicate on the "repo_id" field.
func RepoIDNotIn(vs ...int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldNotIn(FieldRepoID, vs...))
}

// RepoIDGT applies the GT predicate on the "repo_id" field.
func RepoIDGT(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGT(FieldRepoID, v))
}

// RepoIDGTE applies the GTE predicate on the "repo_id" field.
func RepoIDGTE(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldGTE(FieldRepoID, v))
}

// RepoIDLT applies the LT predicate on the "repo_id" field.
func RepoIDLT(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLT(FieldRepoID, v))
}

// RepoIDLTE applies the LTE predicate on the "repo_id" field.
func RepoIDLTE(v int64) predicate.GithubEvent {
	return predicate.GithubEvent(sql.FieldLTE(FieldRepoID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GithubEvent) predicate.GithubEvent {
	return predicate.GithubEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GithubEvent) predicate.GithubEvent {
	return predicate.GithubEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GithubEvent) predicate.GithubEvent {
	return predicate.GithubEvent(func(s *sql.Selector) {
		p(s.Not())
	})
}
