// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

import (
	"net/http"

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

		return 0
	})
}
