/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { client } from "#hey-api/client.gen"

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook("app:beforeMount", () => {
    const loading = useLoadingIndicator()

    client.setConfig({
      onRequest(config) {
        loading.start()
        return config
      },
      onResponse(ctx) {
        ctx.response.status < 300 ? loading.finish() : loading.finish({ error: true })
        return ctx
      },
    })
  })
})
