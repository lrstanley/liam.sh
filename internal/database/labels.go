// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
)

var reLabelSanitize = regexp.MustCompile(`[^a-z\d\-]+`)

// LabelCreator is a helper for creating and populating a label cache.
type LabelCreator struct {
	cache map[string]*ent.Label
}

func NewLabelCreator() *LabelCreator {
	return &LabelCreator{
		cache: make(map[string]*ent.Label),
	}
}

// Filter sanitizes, sorts, and deduplicates the input.
func (lc *LabelCreator) Filter(input []string) (output []string) {
	for i := range input {
		input[i] = strings.Trim(reLabelSanitize.ReplaceAllString(strings.ToLower(input[i]), "-"), "-")
	}

	slices.Sort(input)
	return slices.Compact(input)
}

// Get maps label names to ent objects.
func (lc *LabelCreator) Get(input []string) (output []*ent.Label) {
	for _, l := range lc.Filter(input) {
		if item, ok := lc.cache[l]; ok {
			output = append(output, item)
		}
	}

	return output
}

// Populate queries or creates labels based off the provided input.
func (lc *LabelCreator) Populate(ctx context.Context, db *ent.Tx, input []string) (err error) {
	var item *ent.Label
	var ok bool

	for _, l := range lc.Filter(input) {
		if _, ok = lc.cache[l]; ok {
			continue
		}

		item, err = db.Label.Query().Where(label.NameEqualFold(l)).First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return fmt.Errorf("failed to query label: %q; %w", l, err)
		}

		if item != nil {
			lc.cache[l] = item
			continue
		}

		item, err = db.Label.Create().SetName(l).Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create label: %q; %w", l, err)
		}

		lc.cache[l] = item
	}

	return nil
}

func (lc *LabelCreator) Diff(latest []*ent.Label, previous []int) (add, remove []int) {
	for _, l := range latest {
		contains := false
		for _, p := range previous {
			if l.ID == p {
				contains = true
				break
			}
		}
		if !contains {
			add = append(add, l.ID)
		}
	}

	for _, p := range previous {
		contains := false
		for _, l := range latest {
			if l.ID == p {
				contains = true
				break
			}
		}
		if !contains {
			remove = append(remove, p)
		}
	}

	return add, remove
}
