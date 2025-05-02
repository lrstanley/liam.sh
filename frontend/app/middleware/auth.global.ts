/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { client } from "#hey-api/client.gen"

export default defineNuxtRouteMiddleware(async (to, from) => {
  const runtime = useRuntimeConfig()
  setHTTPClientConfig(client)

  const self = useSelf()
  const githubUser = useGithubUser()
  const url = useRequestURL()

  // TODO: switch these from $fetch to useFetch, after the bugs with hey-api/openapi-ts are fixed.
  await Promise.all([
    callOnce("getSelf", async () => {
      let apiError: ErrorUnauthorized | undefined

      try {
        self.value = await getSelf({ composable: "$fetch" })
      } catch (error) {
        apiError = error as ErrorUnauthorized
        self.value = null
      }
    }),
    callOnce("getGithubUser", async () => {
      let apiError: ErrorUnauthorized | undefined

      try {
        githubUser.value = await getGithubUser({ composable: "$fetch" })
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
