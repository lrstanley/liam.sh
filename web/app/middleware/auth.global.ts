/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

export default defineNuxtRouteMiddleware(async (to, from) => {
  const runtime = useRuntimeConfig()
  const self = useSelf()
  const githubUser = useGithubUser()
  const url = useRequestURL()

  const { $api } = useNuxtApp()

  await Promise.all([
    callOnce("getSelf", async () => {
      try {
        self.value = await $api('/self')
      } catch (error) {
        self.value = null
      }
    }),
    callOnce("getGithubUser", async () => {
      try {
        githubUser.value = await $api('/github-user')
      } catch (error) {
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
