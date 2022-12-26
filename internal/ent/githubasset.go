// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/go-github/v48/github"
	"github.com/lrstanley/liam.sh/internal/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
)

// GithubAsset is the model entity for the GithubAsset schema.
type GithubAsset struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AssetID holds the value of the "asset_id" field.
	AssetID int64 `json:"asset_id,omitempty"`
	// BrowserDownloadURL holds the value of the "browser_download_url" field.
	BrowserDownloadURL string `json:"browser_download_url,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Label holds the value of the "label" field.
	Label string `json:"label,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// ContentType holds the value of the "content_type" field.
	ContentType string `json:"content_type,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// DownloadCount holds the value of the "download_count" field.
	DownloadCount int64 `json:"download_count,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Uploader holds the value of the "uploader" field.
	Uploader *github.User `json:"uploader,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GithubAssetQuery when eager-loading is set.
	Edges                 GithubAssetEdges `json:"edges"`
	github_release_assets *int
}

// GithubAssetEdges holds the relations/edges for other nodes in the graph.
type GithubAssetEdges struct {
	// Release holds the value of the release edge.
	Release *GithubRelease `json:"release,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// ReleaseOrErr returns the Release value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GithubAssetEdges) ReleaseOrErr() (*GithubRelease, error) {
	if e.loadedTypes[0] {
		if e.Release == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: githubrelease.Label}
		}
		return e.Release, nil
	}
	return nil, &NotLoadedError{edge: "release"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GithubAsset) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case githubasset.FieldUploader:
			values[i] = new([]byte)
		case githubasset.FieldID, githubasset.FieldAssetID, githubasset.FieldSize, githubasset.FieldDownloadCount:
			values[i] = new(sql.NullInt64)
		case githubasset.FieldBrowserDownloadURL, githubasset.FieldName, githubasset.FieldLabel, githubasset.FieldState, githubasset.FieldContentType:
			values[i] = new(sql.NullString)
		case githubasset.FieldCreatedAt, githubasset.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case githubasset.ForeignKeys[0]: // github_release_assets
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GithubAsset", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GithubAsset fields.
func (ga *GithubAsset) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case githubasset.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = int(value.Int64)
		case githubasset.FieldAssetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field asset_id", values[i])
			} else if value.Valid {
				ga.AssetID = value.Int64
			}
		case githubasset.FieldBrowserDownloadURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser_download_url", values[i])
			} else if value.Valid {
				ga.BrowserDownloadURL = value.String
			}
		case githubasset.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ga.Name = value.String
			}
		case githubasset.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				ga.Label = value.String
			}
		case githubasset.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				ga.State = value.String
			}
		case githubasset.FieldContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_type", values[i])
			} else if value.Valid {
				ga.ContentType = value.String
			}
		case githubasset.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				ga.Size = value.Int64
			}
		case githubasset.FieldDownloadCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field download_count", values[i])
			} else if value.Valid {
				ga.DownloadCount = value.Int64
			}
		case githubasset.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ga.CreatedAt = value.Time
			}
		case githubasset.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ga.UpdatedAt = value.Time
			}
		case githubasset.FieldUploader:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uploader", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ga.Uploader); err != nil {
					return fmt.Errorf("unmarshal field uploader: %w", err)
				}
			}
		case githubasset.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field github_release_assets", value)
			} else if value.Valid {
				ga.github_release_assets = new(int)
				*ga.github_release_assets = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryRelease queries the "release" edge of the GithubAsset entity.
func (ga *GithubAsset) QueryRelease() *GithubReleaseQuery {
	return (&GithubAssetClient{config: ga.config}).QueryRelease(ga)
}

// Update returns a builder for updating this GithubAsset.
// Note that you need to call GithubAsset.Unwrap() before calling this method if this GithubAsset
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *GithubAsset) Update() *GithubAssetUpdateOne {
	return (&GithubAssetClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the GithubAsset entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *GithubAsset) Unwrap() *GithubAsset {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: GithubAsset is not a transactional entity")
	}
	ga.config.driver = _tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *GithubAsset) String() string {
	var builder strings.Builder
	builder.WriteString("GithubAsset(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("asset_id=")
	builder.WriteString(fmt.Sprintf("%v", ga.AssetID))
	builder.WriteString(", ")
	builder.WriteString("browser_download_url=")
	builder.WriteString(ga.BrowserDownloadURL)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ga.Name)
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(ga.Label)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(ga.State)
	builder.WriteString(", ")
	builder.WriteString("content_type=")
	builder.WriteString(ga.ContentType)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", ga.Size))
	builder.WriteString(", ")
	builder.WriteString("download_count=")
	builder.WriteString(fmt.Sprintf("%v", ga.DownloadCount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ga.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ga.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uploader=")
	builder.WriteString(fmt.Sprintf("%v", ga.Uploader))
	builder.WriteByte(')')
	return builder.String()
}

// GithubAssets is a parsable slice of GithubAsset.
type GithubAssets []*GithubAsset

func (ga GithubAssets) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
