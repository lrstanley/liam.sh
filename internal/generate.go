// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package internal

//go:generate go run -mod=readonly database/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen generate --config graphql/gqlgen.yml

import (
	_ "entgo.io/contrib/entgql" // tools.
	_ "entgo.io/ent/entc/gen"   // tools.
)
