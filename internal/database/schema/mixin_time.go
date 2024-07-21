// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/lrstanley/entrest"
)

var _ ent.Mixin = (*MixinTime)(nil)

type MixinTime struct {
	mixin.Schema
}

func (MixinTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Default(time.Now).
			Immutable().
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("Time the entity was created."),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqualExact|entrest.FilterGroupLength),
			).
			Comment("Time the entity was last updated."),
	}
}
