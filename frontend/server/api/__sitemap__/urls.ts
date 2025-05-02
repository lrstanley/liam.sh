/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { defineSitemapEventHandler } from "#imports"
import type { PagedResponse, ListPostsResponse, ListPostsData } from "#hey-api/types.gen"

type PagableResponse<T> = PagedResponse & {
  content: T[]
}

// TODO: cannot use hey-api client until this is fixed: https://github.com/hey-api/openapi-ts/issues/1985
//
// type PagableQuery = {
//   composable: "$fetch"
//   query: {
//     page?: number
//     per_page?: number
//   }
// }
//
//
// type PagableFn<T> = (config: PagableQuery) => Promise<PagableResponse<T>>
// type PagedType<T> = T extends PagableFn<infer T> ? T : never
//
// async function paginate<F extends PagableFn<T>, T = PagedType<F>>(
//   fn: F,
//   baseOptions: Omit<Parameters<F>[0], "composable">
// ): Promise<T[]> {
//   const items: T[] = []
//   let page = 1
//   let per_page = 100

//   while (true) {
//     const resp = await fn({
//       ...baseOptions,
//       composable: "$fetch",
//       query: { ...baseOptions.query, page, per_page },
//     })
//     items.push(...resp.content)

//     if (resp.is_last_page) break
//   }

//   return items
// }

// const result = await $fetch<ListPostsResponse>(`${getBackendURL()}/-/posts`, {
//   method: "GET",
//   query: {
//     "public.eq": true,
//   },
// })

type PagableQuery = {
  page?: number
  per_page?: number
}

type PagedType<T> = T extends PagableResponse<infer T> ? T : never

async function paginate<
  R extends PagableResponse<T>,
  Q extends PagableQuery | undefined = {},
  T = PagedType<R>,
>(route: string, params?: Q): Promise<T[]> {
  const BASE_URL = getBackendURL()

  const items: T[] = []
  let page = 1
  let per_page = 100

  while (true) {
    const resp = await $fetch<R>(`${BASE_URL}${route}`, {
      method: "GET",
      query: { ...params, page, per_page },
    })
    items.push(...resp.content)

    if (resp.is_last_page) break
    page++
  }

  return items
}

export default defineSitemapEventHandler(async (event) => {
  return [
    ...(await paginate<ListPostsResponse, ListPostsData["query"]>("/posts", { "public.eq": true })).map(
      (p) => ({
        loc: `/p/${p.slug}`,
        _sitemap: "posts",
      })
    ),
  ]
})
