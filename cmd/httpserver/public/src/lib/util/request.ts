/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { shallowEqual } from "@/lib/util/equal"
import type { RequestResult } from "@hey-api/client-fetch"
import type {
  ErrorBadRequest,
  ErrorForbidden,
  ErrorInternalServerError,
  ErrorTooManyRequests,
  ErrorUnauthorized,
  SortOrder,
} from "@/lib/http/types.gen"
import { useRouteQuery } from "@vueuse/router"
import type { WatchStopHandle } from "vue"

type ErrorResponse = ErrorBadRequest &
  ErrorUnauthorized &
  ErrorForbidden &
  ErrorTooManyRequests &
  ErrorInternalServerError

class ApiError extends Error {
  constructor(
    public readonly response: ErrorResponse,
    options?: ErrorOptions
  ) {
    super(response.error, options)
  }
}

export async function unwrapErrors<R>(prom: RequestResult<R, ErrorResponse>): Promise<R> {
  const resp = await prom

  if (resp.error) {
    // it's a paginated response, which should be a 404, but has a regular body.
    if ((resp.error as any).content !== undefined && (resp.error as any).page !== undefined) {
      return resp.error as R
    }
    throw new ApiError(resp.error)
  }
  return resp.data
}

export type PaginationOptions<T> = {
  page?: number
  perPage?: number
  sort?: T
  order?: SortOrder
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
    order: useRouteQuery<SortOrder>("order", order),
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
