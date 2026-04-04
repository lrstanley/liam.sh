/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { FetchOptions, ResponseType } from 'ofetch'

export default defineNuxtPlugin({
  enforce: 'pre',
  setup(nuxtApp) {
    const options = useRuntimeConfig().public.openFetch.api as FetchOptions<ResponseType, any>
    const localFetch = useRequestFetch()

    options.headers = options.headers ? { ...options.headers } : {}

    options.retry = 3
    options.retryDelay = 1000

    if (import.meta.server) {
      const incomingCookie = useRequestHeaders(['cookie']).cookie
      if (incomingCookie) {
        (options.headers as Record<string, string>).Cookie = incomingCookie
      }

      options.baseURL = useRuntimeConfig().API_URL.replace(/\/$/, "")
    } else {
      const browserCookie = useCookie('session', { path: '/' }).value
      if (browserCookie) {
        (options.headers as Record<string, string>).Cookie = browserCookie
      }

      options.baseURL = useRuntimeConfig().public.API_URL.replace(/\/$/, "")
    }

    return {
      provide: {
        api: createOpenFetch(localOptions => ({
          ...options,
          ...localOptions,
          credentials: 'include',
        }), localFetch as any, 'core', nuxtApp.hooks)
      }
    }
  }
})
