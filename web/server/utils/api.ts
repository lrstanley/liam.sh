/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in the
 * LICENSE file.
 */

import { createError, useRuntimeConfig } from "nuxt/server"
import type { SchemaPostList } from "#open-fetch-schemas/api"

export async function fetchPosts(
  query: Record<string, string | number | boolean> = {},
): Promise<SchemaPostList> {
  const config = useRuntimeConfig()
  const url = new URL(`${String(config.API_URL).replace(/\/+$/, "")}/posts`)

  for (const [key, value] of Object.entries(query)) {
    url.searchParams.set(key, String(value))
  }

  const response = await fetch(url)
  if (!response.ok) {
    throw createError({
      status: response.status,
      statusText: response.statusText || "api request failed",
    })
  }

  return (await response.json()) as SchemaPostList
}
