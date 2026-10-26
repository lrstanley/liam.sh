/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in the
 * LICENSE file.
 */

import { defineEventHandler } from "nuxt/server"
import { getSiteConfig } from "#site-config/server/composables"
import { fetchAllPosts, formatPostMarkdown } from "../utils/api"

// Served by us rather than nuxt-ai-ready (aiReady.llmsFullTxt is disabled), as
// the module builds llms-full.txt from a prerender/sqlite page index, and posts
// are fetched from the API at runtime. This uses the post source markdown.
export default defineEventHandler(async (event) => {
  const site = getSiteConfig(event)
  const origin = String(site.url || "").replace(/\/+$/, "")
  const posts = await fetchAllPosts({ "public.eq": true })

  const parts = [`# ${site.name}`, ""]
  if (site.description) {
    parts.push(`> ${site.description}`, "")
  }
  parts.push("Full content of every published blog post, in markdown format.", "")

  for (const post of posts) {
    parts.push(
      "---",
      "",
      `- **Source:** ${origin}/p/${post.slug}`,
      `- **Published:** ${post.published_at}`,
      "",
      formatPostMarkdown(post).trim(),
      "",
    )
  }

  event.res.headers.set("content-type", "text/plain; charset=utf-8")
  event.res.headers.set("cache-control", "public, max-age=600, stale-while-revalidate=3600")

  return parts.join("\n")
})
