/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { defineSitemapEventHandler } from "#imports"
import type { SchemaPostRead } from '#open-fetch-schemas/api'

export default defineSitemapEventHandler(async (event) => {
  const { $api } = useNitroApp()

  const posts: SchemaPostRead[] = []
  let page = 1

  while (true) {
    const resp = await $api('/posts', {
      query: {
        page: page,
        per_page: 100,
        "public.eq": true,
      },
    })
    posts.push(...resp.content)

    if (resp.is_last_page) break
    page++
  }

  return posts.map((p) => ({
    loc: `/p/${p.slug}`,
    _sitemap: "posts",
    lastmod: p.published_at,
  }))
})
