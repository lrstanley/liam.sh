// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package wakapi

import (
	"maps"
	"slices"

	"github.com/lrstanley/x/text/fuzzy"
)

// lookupColor returns the hex color for a given language name from the internal
// language color map. If the language name is not found, it will return the
// closest match.
func lookupColor(name string) (hex string) {
	sortedKeys := slices.Collect(maps.Keys(languageColorMap))
	slices.Sort(sortedKeys)
	return languageColorMap[fuzzy.ShortestDistance(name, sortedKeys)]
}
