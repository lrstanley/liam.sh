/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */
import { getSelf, getGithubUser } from "@/utils/http/sdk.gen"
import type { ErrorUnauthorized } from "@/utils/http/types.gen"

export default defineNuxtRouteMiddleware(async (to, from) => {
  setHTTPClientBaseURL()

  const self = useSelf()
  const githubUser = useGithubUser()

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

  if (to.meta.auth == true && self.value == null) {
    return navigateTo(`/-/auth/providers/github?next=${window.location.origin + to.path}`)
  }
})
