// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package models

import "github.com/jackc/pgerrcode"

// TODO: stub for now.
func IsDatabaseError(err error) bool {
	return pgerrcode.IsWarning(err.Error())
}
