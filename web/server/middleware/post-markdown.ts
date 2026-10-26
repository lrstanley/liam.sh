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
import { fetchPosts, formatPostMarkdown } from "../utils/api"

export default defineEventHandler(async (event) => {
  const match = getRequestURL(event).pathname.match(/^\/p\/(.+)\.md$/)
  const slug = match?.[1]
  if (!slug) return

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
