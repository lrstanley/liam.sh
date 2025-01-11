/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { client } from "@/utils/http/sdk.gen"

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook("app:beforeMount", () => {
    const runtime = useRuntimeConfig()

    let url: string

    if (runtime.API_URL) {
      url = runtime.API_URL as string
    } else if (runtime.public.API_URL) {
      url = runtime.public.API_URL as string
    } else {
      url = `${useRequestURL().origin}/-`
    }

    client.setConfig({
      baseUrl: url,
    })

    client.interceptors.request.use((request) => {
      loadingBar.start()
      return request
    })
    client.interceptors.response.use((response) => {
      response.status < 300 ? loadingBar.finish() : loadingBar.error()
      return response
    })
  })
})
