// import urql from "@urql/vue"
import { createClient, provideClient } from '@urql/vue'
import { dedupExchange, cacheExchange, fetchExchange } from "@urql/vue"
import { retryExchange } from "@urql/exchange-retry"

function fetchWithTimeout(url: RequestInfo, opts: RequestInit): Promise<any> {
    const controller = new AbortController()
    const id = setTimeout(() => controller.abort(), 5000)

    const promise = new Promise((resolve, reject) => {
        fetch(url, {
            ...opts,
            signal: controller.signal
        }).then((resp) => {
            clearTimeout(id)
            resolve(resp)
        }).catch((err) => {
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
            retryIf: (err) => err && err.networkError ? true : false,
        }),
        fetchExchange,
    ],
})
