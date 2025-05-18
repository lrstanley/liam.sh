/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { client } from "#hey-api/client.gen"

export default defineNuxtRouteMiddleware(async (to, from) => {
  const runtime = useRuntimeConfig()

  client.setConfig({
    baseURL: runtime.API_URL
      ? runtime.API_URL.replace(/\/$/, "")
      : runtime.public.API_URL.replace(/\/$/, ""),
    credentials: "include",
    retry: 3,
    retryDelay: 1000,
    cache: "no-cache",
  })

  const self = useSelf()
  const githubUser = useGithubUser()
  const url = useRequestURL()

  await Promise.all([
    callOnce("getSelf", async () => {
      let apiError: ErrorUnauthorized | undefined

      try {
        const results = await getSelf({
          composable: "useFetch",
          headers: useRequestHeaders(["cookie"]), // Allows pass-through of cookies from origin browser.
        })
        self.value = results.data.value
      } catch (error) {
        apiError = error as ErrorUnauthorized
        self.value = null
      }
    }),
    callOnce("getGithubUser", async () => {
      let apiError: ErrorUnauthorized | undefined

      try {
        const results = await getGithubUser({
          composable: "useFetch",
          headers: useRequestHeaders(["cookie"]), // Allows pass-through of cookies from origin browser.
        })
        githubUser.value = results.data.value
      } catch (error) {
        if (!apiError) apiError = error as ErrorUnauthorized
        githubUser.value = null
        throw error
      }
    }),
  ])

  if (to.path.startsWith("/admin") && self.value == null) {
    return navigateTo(
      `${runtime.public.API_URL.replace(/\/$/, "")}/auth/providers/github?next=${url.origin + to.path}`,
      {
        external: true,
      }
    )
  }
})
