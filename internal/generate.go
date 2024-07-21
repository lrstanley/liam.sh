// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package internal

//go:generate go run -mod=readonly database/entc.go

import (
	_ "entgo.io/ent/entc/gen"        // tools.
	_ "github.com/lrstanley/entrest" // tools.
)
