/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { shallowEqual } from "@/utils/equal"
import { useRouteQuery } from "@vueuse/router"
import type { WatchStopHandle } from "vue"
import { client } from "@/utils/http/sdk.gen"

export function setHTTPClientBaseURL() {
  const runtime = useRuntimeConfig()

  let url: string
  if (runtime.API_URL) {
    url = runtime.API_URL as string
  } else if (runtime.public.API_URL) {
    url = runtime.public.API_URL as string
  } else {
    url = `${useRequestURL().origin}/-`
  }

  client.setConfig({ baseURL: url })
}

export type PaginationOptions<T> = {
  page?: number
  perPage?: number
  sort?: T
  order?: "asc" | "desc"
}

export function usePagination<T extends string = string>({
  page = 1,
  perPage = 15,
  sort,
  order = "desc",
}: PaginationOptions<T>) {
  return {
    page: useRouteQuery<number>("page", page, { transform: Number }),
    perPage: useRouteQuery<number>("per_page", perPage, { transform: Number }),
    sort: useRouteQuery<T>("sort", sort),
    order: useRouteQuery<"asc" | "desc">("order", order),
    sizes: [5, 15, 25, 50, 100, 250],
  }
}

/**
 * resetPagination - reset the pagination when any of the provided refs change, as well
 * as perPage, sort, and order from the pagination object itself.
 */
export function resetPagination(
  pagination: ReturnType<typeof usePagination>,
  refs: Ref<any>[]
): WatchStopHandle {
  return watch([...refs, pagination.perPage, pagination.sort, pagination.order], (newv, oldv) => {
    for (let i = 0; i < newv.length; i++) {
      if (!shallowEqual(newv[i], oldv[i])) {
        pagination.page.value = 1
        return
      }
    }
  })
}
