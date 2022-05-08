// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lrstanley/liam.sh/internal/ent/label"
)

// Label is the model entity for the Label schema.
type Label struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LabelQuery when eager-loading is set.
	Edges LabelEdges `json:"edges"`
}

// LabelEdges holds the relations/edges for other nodes in the graph.
type LabelEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e LabelEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Label) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case label.FieldID:
			values[i] = new(sql.NullInt64)
		case label.FieldName:
			values[i] = new(sql.NullString)
		case label.FieldCreateTime, label.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Label", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Label fields.
func (l *Label) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case label.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case label.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				l.CreateTime = value.Time
			}
		case label.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				l.UpdateTime = value.Time
			}
		case label.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		}
	}
	return nil
}

// QueryPosts queries the "posts" edge of the Label entity.
func (l *Label) QueryPosts() *PostQuery {
	return (&LabelClient{config: l.config}).QueryPosts(l)
}

// Update returns a builder for updating this Label.
// Note that you need to call Label.Unwrap() before calling this method if this Label
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Label) Update() *LabelUpdateOne {
	return (&LabelClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Label entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Label) Unwrap() *Label {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Label is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Label) String() string {
	var builder strings.Builder
	builder.WriteString("Label(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(l.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(l.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(l.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Labels is a parsable slice of Label.
type Labels []*Label

func (l Labels) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
