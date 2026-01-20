// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package apihandler

import (
	"context"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/lrstanley/chix/v2"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
)

type LabelsCountResponse struct {
	ent.Label
	PostCount             int `json:"post_count"`
	GithubRepositoryCount int `json:"githubrepository_count"`
	Total                 int `json:"total_count"`
}

func (h *handler) getLabelsCount(w http.ResponseWriter, r *http.Request) {
	v := []*LabelsCountResponse{}

	err := h.db.Label.Query().GroupBy(label.FieldID).
		Aggregate(
			func(s *sql.Selector) string {
				s.AppendSelect(label.Columns...)

				t1 := sql.Table(label.GithubRepositoriesTable)
				s.LeftJoin(t1).On(s.C(label.FieldID), t1.C(label.GithubRepositoriesPrimaryKey[0]))
				t1count := sql.Count(sql.Distinct(t1.C(label.GithubRepositoriesPrimaryKey[1])))

				t2 := sql.Table(label.PostsTable)
				s.LeftJoin(t2).On(s.C(label.FieldID), t2.C(label.PostsPrimaryKey[0]))
				t2count := sql.Count(sql.Distinct(t2.C(label.PostsPrimaryKey[1])))

				s.AppendSelectAs(t1count, "githubrepository_count")
				s.AppendSelectAs(t2count, "post_count")
				s.AppendSelectExprAs(sql.ColumnsOp(t1count, t2count, sql.OpAdd), "total_count")
				s.OrderBy(sql.Desc("total_count"))

				return ""
			},
		).Scan(context.Background(), &v)

	if chix.IfError(w, r, err) {
		return
	}

	chix.JSON(w, r, http.StatusOK, v)
}
