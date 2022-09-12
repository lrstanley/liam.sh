// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

type CodingStats struct {
	Languages      []LanguageStat `json:"languages"`
	TotalSeconds   int            `json:"total_seconds"`
	TotalDuration  string         `json:"total_duration"`
	CalculatedDays int            `json:"calculated_days"`
}

type LanguageStat struct {
	Language      string `json:"key"`
	TotalSeconds  int    `json:"total"`
	TotalDuration string `json:"total_duration"`
}
