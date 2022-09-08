import { retryExchange } from "@urql/exchange-retry"
import { cacheExchange, createClient, dedupExchange, fetchExchange } from "@urql/vue"

function fetchWithTimeout(url: RequestInfo, opts: RequestInit): Promise<Response> {
  const controller = new AbortController()
  const id = setTimeout(() => controller.abort(), 5000)

  const promise = new Promise<Response>((resolve, reject) => {
    fetch(url, {
      ...opts,
      signal: controller.signal,
    })
      .then((resp) => {
        clearTimeout(id)
        resolve(resp)
      })
      .catch((err) => {
        resolve(err)
      })
  })
  return promise
}

export const client = createClient({
  url: "/-/graphql",
  requestPolicy: "cache-and-network",
  fetch: fetchWithTimeout,
  exchanges: [
    dedupExchange,
    cacheExchange,
    retryExchange({
      initialDelayMs: 1000,
      maxDelayMs: 15000,
      randomDelay: true,
      maxNumberAttempts: 3,
      retryIf: (err) => (err && err.networkError ? true : false),
    }),
    fetchExchange,
  ],
})
