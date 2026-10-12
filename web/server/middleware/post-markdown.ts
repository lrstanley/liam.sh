/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import {
  createError,
  defineEventHandler,
  getRequestURL,
} from "nuxt/server"
import type { SchemaPostRead } from "#open-fetch-schemas/api"
import { fetchPosts } from "../utils/api"

function formatPostMarkdown(post: SchemaPostRead): string {
  const lines = [`# ${post.title}`, ""]

  if (post.summary) {
    lines.push(`> ${post.summary}`, "")
  }

  lines.push(post.content)

  return lines.join("\n")
}

export default defineEventHandler(async (event) => {
  const match = getRequestURL(event).pathname.match(/^\/p\/(.+)\.md$/)
  if (!match) return

  const slug = match[1]
  const resp = await fetchPosts({
    "slug.eq": slug,
  })

  const post = resp.content?.[0]
  if (!post) {
    throw createError({ status: 404, statusText: "post not found" })
  }

  event.res.headers.set("content-type", "text/markdown; charset=utf-8")
  event.res.headers.set("cache-control", "public, max-age=300")

  return formatPostMarkdown(post)
})
