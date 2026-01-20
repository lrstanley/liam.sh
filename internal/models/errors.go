// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

import (
	"net/http"

	"github.com/lrstanley/chix/v2"
	"github.com/lrstanley/liam.sh/internal/database/ent"
)

func ErrorResolver(oerr *chix.ResolvedError) *chix.ResolvedError {
	switch {
	case ent.IsConstraintError(oerr.Err) || ent.IsValidationError(oerr.Err):
		return &chix.ResolvedError{
			StatusCode: http.StatusBadRequest,
			Err:        oerr.Err,
			Errs:       oerr.Errs,
			Visibility: chix.ErrorMasked,
		}
	case ent.IsNotFound(oerr.Err):
		return &chix.ResolvedError{
			StatusCode: http.StatusNotFound,
			Err:        oerr.Err,
			Errs:       oerr.Errs,
			Visibility: chix.ErrorMasked,
		}
	default:
		return oerr
	}
}
