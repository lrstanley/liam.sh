/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { defineSitemapEventHandler } from "#sitemap/server"
import { fetchAllPosts } from "../../utils/api"

export default defineSitemapEventHandler(async () => {
  const posts = await fetchAllPosts({ "public.eq": true })

  return posts.map((p) => ({
    loc: `/p/${p.slug}`,
    _sitemap: "posts",
    lastmod: p.published_at,
  }))
})
