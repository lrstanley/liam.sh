/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */
import { getSelf, getGithubUser } from "@/utils/http/sdk.gen"
import type { ErrorUnauthorized } from "@/utils/http/types.gen"
import { client } from "@/utils/http/sdk.gen"

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
  })

  const self = useSelf()
  const githubUser = useGithubUser()

  let apiError: ErrorUnauthorized | undefined

  if (self.value === undefined || (from.path == "/" && from.name == undefined)) {
    const selfResults = await getSelf({ composable: "$fetch" })

    // if (error) {
    //   apiError = error as ErrorUnauthorized
    //   self.value = null
    // } else {
    //   self.value = data
    // }
    self.value = selfResults
  }

  const githubUserResult = await getGithubUser({ composable: "$fetch" })
  // if (error) {
  //   if (!apiError) apiError = error as ErrorUnauthorized
  // } else {
  //   githubUser.value = data
  // }
  githubUser.value = githubUserResult

  if (to.meta.auth == true && self.value == null) {
    return navigateTo(`/-/auth/providers/github?next=${window.location.origin + to.path}`)
  }

  if (apiError !== undefined && to.name !== "/[...catchall]" && apiError?.code !== 401) {
    console.log(apiError)
    return navigateTo({ name: "/[...catchall]", params: { catchall: apiError.error } })
  }
})
