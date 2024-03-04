// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/go-github/v52/github"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubgist"
)

// GithubGist is the model entity for the GithubGist schema.
type GithubGist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GistID holds the value of the "gist_id" field.
	GistID string `json:"gist_id,omitempty"`
	// HTMLURL holds the value of the "html_url" field.
	HTMLURL string `json:"html_url,omitempty"`
	// Public holds the value of the "public" field.
	Public bool `json:"public,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner *github.User `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// RawURL holds the value of the "raw_url" field.
	RawURL string `json:"raw_url,omitempty"`
	// Content holds the value of the "content" field.
	Content      string `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GithubGist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case githubgist.FieldOwner:
			values[i] = new([]byte)
		case githubgist.FieldPublic:
			values[i] = new(sql.NullBool)
		case githubgist.FieldID, githubgist.FieldSize:
			values[i] = new(sql.NullInt64)
		case githubgist.FieldGistID, githubgist.FieldHTMLURL, githubgist.FieldDescription, githubgist.FieldName, githubgist.FieldType, githubgist.FieldLanguage, githubgist.FieldRawURL, githubgist.FieldContent:
			values[i] = new(sql.NullString)
		case githubgist.FieldCreatedAt, githubgist.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GithubGist fields.
func (gg *GithubGist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case githubgist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gg.ID = int(value.Int64)
		case githubgist.FieldGistID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gist_id", values[i])
			} else if value.Valid {
				gg.GistID = value.String
			}
		case githubgist.FieldHTMLURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html_url", values[i])
			} else if value.Valid {
				gg.HTMLURL = value.String
			}
		case githubgist.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				gg.Public = value.Bool
			}
		case githubgist.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gg.CreatedAt = value.Time
			}
		case githubgist.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gg.UpdatedAt = value.Time
			}
		case githubgist.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gg.Description = value.String
			}
		case githubgist.FieldOwner:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gg.Owner); err != nil {
					return fmt.Errorf("unmarshal field owner: %w", err)
				}
			}
		case githubgist.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gg.Name = value.String
			}
		case githubgist.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				gg.Type = value.String
			}
		case githubgist.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				gg.Language = value.String
			}
		case githubgist.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				gg.Size = value.Int64
			}
		case githubgist.FieldRawURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw_url", values[i])
			} else if value.Valid {
				gg.RawURL = value.String
			}
		case githubgist.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				gg.Content = value.String
			}
		default:
			gg.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GithubGist.
// This includes values selected through modifiers, order, etc.
func (gg *GithubGist) Value(name string) (ent.Value, error) {
	return gg.selectValues.Get(name)
}

// Update returns a builder for updating this GithubGist.
// Note that you need to call GithubGist.Unwrap() before calling this method if this GithubGist
// was returned from a transaction, and the transaction was committed or rolled back.
func (gg *GithubGist) Update() *GithubGistUpdateOne {
	return NewGithubGistClient(gg.config).UpdateOne(gg)
}

// Unwrap unwraps the GithubGist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gg *GithubGist) Unwrap() *GithubGist {
	_tx, ok := gg.config.driver.(*txDriver)
	if !ok {
		panic("ent: GithubGist is not a transactional entity")
	}
	gg.config.driver = _tx.drv
	return gg
}

// String implements the fmt.Stringer.
func (gg *GithubGist) String() string {
	var builder strings.Builder
	builder.WriteString("GithubGist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gg.ID))
	builder.WriteString("gist_id=")
	builder.WriteString(gg.GistID)
	builder.WriteString(", ")
	builder.WriteString("html_url=")
	builder.WriteString(gg.HTMLURL)
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", gg.Public))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gg.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gg.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(gg.Description)
	builder.WriteString(", ")
	builder.WriteString("owner=")
	builder.WriteString(fmt.Sprintf("%v", gg.Owner))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gg.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(gg.Type)
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(gg.Language)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", gg.Size))
	builder.WriteString(", ")
	builder.WriteString("raw_url=")
	builder.WriteString(gg.RawURL)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(gg.Content)
	builder.WriteByte(')')
	return builder.String()
}

// GithubGists is a parsable slice of GithubGist.
type GithubGists []*GithubGist
