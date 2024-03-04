// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

import (
	"errors"
	"net/http"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
)

func init() {
	chix.AddErrorResolver(func(err error) (status int) {
		if ent.IsConstraintError(err) || ent.IsValidationError(err) {
			return http.StatusBadRequest
		}

		if ent.IsNotFound(err) {
			return http.StatusNotFound
		}

		if IsDatabaseError(err) {
			if code := unwrapDBError(err); code != "" {
				if pgerrcode.IsDataException(code) {
					return http.StatusBadRequest
				}
			}
			return http.StatusInternalServerError
		}

		return 0
	})
}

// TODO: is pgerrcode needed for anything?

func unwrapDBError(err error) string {
	var pge *pgconn.PgError
	if errors.As(err, &pge) {
		return pge.Code
	}
	return ""
}

func IsDatabaseError(err error) bool {
	return unwrapDBError(err) != ""
}
