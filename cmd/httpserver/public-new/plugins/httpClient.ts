/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { client } from "@/utils/http/sdk.gen"

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook("app:beforeMount", () => {
    client.setConfig({
      onRequest(config) {
        loadingBar.start()
        return config
      },
      onResponse(ctx) {
        ctx.response.status < 300 ? loadingBar.finish() : loadingBar.error()
        return ctx
      },
    })
  })
})
