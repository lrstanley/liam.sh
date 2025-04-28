/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { WatchStopHandle } from "vue"
import { client } from "#hey-api/client.gen"

export function getBackendURL(): string {
  const runtime = useRuntimeConfig()

  if (runtime.API_URL) {
    return (runtime.API_URL as string).replace(/\/$/, "")
  } else if (runtime.public.API_URL) {
    return (runtime.public.API_URL as string).replace(/\/$/, "")
  }

  return `${useRequestURL().origin}/-`
}

export function setHTTPClientConfig() {
  // Allows pass-through of cookies from origin browser, to API requests (mainly helpful when in SSR).
  const headers = useRequestHeaders(["cookie"])

  client.setConfig({
    baseURL: getBackendURL(),
    credentials: "include",
    retry: 3,
    retryDelay: 1000,
    headers: headers,
  })
}

export type PaginationOptions<T> = {
  page?: number
  perPage?: number
  sort?: T
  order?: "asc" | "desc"
  resetChanged?: Ref<any>[]
}

export function usePagination<T extends string = string>({
  page = 1,
  perPage = 15,
  sort,
  order = "desc",
  resetChanged = [],
}: PaginationOptions<T>) {
  const results = {
    page: useRouteQuery<number>("page", page, { transform: Number }),
    perPage: useRouteQuery<number>("per_page", perPage, { transform: Number }),
    sort: useRouteQuery<T>("sort", sort),
    order: useRouteQuery<"asc" | "desc">("order", order),
    sizes: [5, 15, 25, 50, 100, 250],
  }

  if (resetChanged.length) {
    watch([...resetChanged, results.perPage, results.sort, results.order], (newv, oldv) => {
      for (let i = 0; i < newv.length; i++) {
        if (!shallowEqual(newv[i], oldv[i])) {
          results.page.value = 1
          return
        }
      }
    })
  }

  return results
}
