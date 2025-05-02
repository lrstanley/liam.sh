/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

// don't import client here until https://github.com/hey-api/openapi-ts/issues/1985 is resolved.
export async function setHTTPClientConfig(client: any) {
  const runtime = useRuntimeConfig()

  const baseURL = runtime.API_URL
    ? runtime.API_URL.replace(/\/$/, "")
    : runtime.public.API_URL
      ? runtime.public.API_URL.replace(/\/$/, "")
      : `${useRequestURL().origin}/-`

  // Allows pass-through of cookies from origin browser, to API requests (mainly helpful when in SSR).
  const headers = useRequestHeaders(["cookie"])

  client.setConfig({
    baseURL: baseURL,
    credentials: "include",
    retry: 3,
    retryDelay: 1000,
    headers: headers,
  })
}
