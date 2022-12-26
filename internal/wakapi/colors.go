// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package wakapi

import (
	"sort"

	"github.com/agnivade/levenshtein"
)

// lookupColor returns the hex color for a given language name from the internal
// language color map. If the language name is not found, it will return the
// closest match.
func lookupColor(name string) (hex string) {
	lowestCalc := 1000

	sortedKeys := make([]string, 0, len(languageColorMap))
	for k := range languageColorMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, lang := range sortedKeys {
		calc := levenshtein.ComputeDistance(name, lang)

		if calc < lowestCalc {
			lowestCalc = calc
			hex = languageColorMap[lang]
		}

		if calc == 0 {
			break
		}
	}

	return hex
}
