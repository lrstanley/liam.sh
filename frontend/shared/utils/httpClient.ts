/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

export function getBackendURL(): string {
  const runtime = useRuntimeConfig()

  if (runtime.API_URL) {
    return runtime.API_URL.replace(/\/$/, "")
  } else if (runtime.public.API_URL) {
    return runtime.public.API_URL.replace(/\/$/, "")
  }

  return `${useRequestURL().origin}/-`
}

// don't import client here until https://github.com/hey-api/openapi-ts/issues/1985 is resolved.
export async function setHTTPClientConfig(client: any) {
  // Allows pass-through of cookies from origin browser, to API requests (mainly helpful when in SSR).
  const headers = useRequestHeaders(["cookie"])

  client.setConfig({
    baseURL: getBackendURL(),
    credentials: "include",
    retry: 3,
    retryDelay: 1000,
    headers: headers,
  })
}
