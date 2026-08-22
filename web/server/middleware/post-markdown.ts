/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { createError, defineEventHandler, setHeader } from "#imports"
import type { SchemaPostRead } from "#open-fetch-schemas/api"

function formatPostMarkdown(post: SchemaPostRead): string {
  const lines = [`# ${post.title}`, ""]

  if (post.summary) {
    lines.push(`> ${post.summary}`, "")
  }

  lines.push(post.content)

  return lines.join("\n")
}

export default defineEventHandler(async (event) => {
  const match = event.path.match(/^\/p\/(.+)\.md$/)
  if (!match) return

  const slug = match[1]
  const { $api } = useNitroApp()
  const resp = await $api("/posts", {
    query: {
      "slug.eq": slug,
    },
  })

  const post = resp.content?.[0]
  if (!post) {
    throw createError({ statusCode: 404, statusMessage: "post not found" })
  }

  setHeader(event, "content-type", "text/markdown; charset=utf-8")
  setHeader(event, "cache-control", "public, max-age=300")

  return formatPostMarkdown(post)
})
