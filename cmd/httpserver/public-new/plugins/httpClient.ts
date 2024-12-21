/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { client } from "@/utils/http/services.gen"

type Mutable = {
  -readonly [key in keyof Response]: Response[key]
}

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
      // fetch: async (request) => {
      //   return new Promise((resolve, reject) => {
      //     useFetch(request, {
      //       onResponse: ({ response: r }) => {
      //         resolve(
      //           new Response(JSON.stringify(toRaw(r._data)), {
      //             status: r.status,
      //             statusText: r.statusText,
      //             headers: r.headers,
      //           })
      //         )
      //         // resolve(r.clone())
      //         // response = r.clone()
      //         // response = new Response(JSON.stringify(toRaw(resp.data.value)), {
      //         //   status: r.status,
      //         //   statusText: r.statusText,
      //         //   headers: r.headers,
      //         // })
      //       },
      //       onResponseError: ({ response: r }) => {
      //         reject(
      //           new Response(JSON.stringify(toRaw(r._data)), {
      //             status: r.status,
      //             statusText: r.statusText,
      //             headers: r.headers,
      //           })
      //         )
      //         // reject(r.clone())
      //         // response = r.clone()
      //         // response = new Response(JSON.stringify(toRaw(resp.error.value)), {
      //         //   status: r.status,
      //         //   statusText: r.statusText,
      //         //   headers: r.headers,
      //         // })
      //       },
      //     })
      //   })
      // },
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
