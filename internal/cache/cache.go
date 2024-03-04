// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package cache

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

type Cache[K comparable, V any] struct {
	rc         *ristretto.Cache
	defaultTTL time.Duration
	fillFn     func(K) (V, bool)
}

func New[K comparable, V any](size int64, defaultTTL time.Duration, fillFn func(K) (V, bool)) *Cache[K, V] {
	rc, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: size,
		MaxCost:     size,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	return &Cache[K, V]{
		rc:         rc,
		defaultTTL: defaultTTL,
		fillFn:     fillFn,
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	_ = c.rc.SetWithTTL(key, value, 1, c.defaultTTL)
}

func (c *Cache[K, V]) Get(key K) (val V) {
	tmp, ok := c.rc.Get(key)
	if !ok {
		if c.fillFn != nil {
			val, ok = c.fillFn(key)
			if ok {
				c.Set(key, val)
			}
		}

		return val
	}

	return tmp.(V)
}
