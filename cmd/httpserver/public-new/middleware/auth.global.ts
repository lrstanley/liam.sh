/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */
import { getSelf, getGithubUser } from "@/utils/http/services.gen"
import type { ErrorUnauthorized } from "@/utils/http/types.gen"
import { client } from "@/utils/http/services.gen"

export default defineNuxtRouteMiddleware(async (to, from) => {
  // TODO: this is fucking stupid.
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
    // fetch: (request) => {
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
    //       },
    //       onResponseError: ({ response: r }) => {
    //         reject(
    //           new Response(JSON.stringify(toRaw(r._data)), {
    //             status: r.status,
    //             statusText: r.statusText,
    //             headers: r.headers,
    //           })
    //         )
    //       },
    //     })
    //   })
    // },
  })

  const self = useSelf()
  const githubUser = useGithubUser()

  let apiError: ErrorUnauthorized | undefined

  if (self.value === undefined || (from.path == "/" && from.name == undefined)) {
    const { data, error } = await getSelf()

    if (error) {
      apiError = error as ErrorUnauthorized
      self.value = null
    } else {
      self.value = data
    }
  }

  const { data, error } = await getGithubUser()
  if (error) {
    if (!apiError) apiError = error as ErrorUnauthorized
  } else {
    githubUser.value = data
  }

  if (to.meta.auth == true && self.value == null) {
    return navigateTo(`/-/auth/providers/github?next=${window.location.origin + to.path}`)
  }

  if (apiError !== undefined && to.name !== "/[...catchall]" && apiError?.code !== 401) {
    console.log(apiError)
    return navigateTo({ name: "/[...catchall]", params: { catchall: apiError.error } })
  }
})
