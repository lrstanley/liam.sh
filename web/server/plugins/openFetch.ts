/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { defineNitroPlugin, useRuntimeConfig, createOpenFetch } from '#imports'
import type { NitroApp } from "nitropack";
import type { FetchOptions, ResponseType } from 'ofetch'

declare module 'nitropack' {
  interface NitroApp {
    $api: ReturnType<typeof createOpenFetch>
  }
}

export default defineNitroPlugin((nitroApp: NitroApp) => {
  const clients = useRuntimeConfig().public.openFetch as Record<string, FetchOptions<ResponseType, any>>

  Object.entries(clients).forEach(([name, client]) => {
    (nitroApp as any)[`$${name}`] = createOpenFetch(
      {
        ...client, baseURL: useRuntimeConfig().API_URL.replace(/\/$/, "")
      },
      nitroApp.localFetch as any,
      name,
      nitroApp.hooks,
    )
  })
})
