/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

export default defineNitroPlugin((nitroApp) => {
  nitroApp.hooks.hook("request", (event) => {
    console.log(
      `[${new Date().toISOString()}] ${event.node.req.method || "method:n/a"} ${JSON.stringify(event.node.req.url)} ${JSON.stringify(event.node.req.headers["host"] || "host:n/a")} ${JSON.stringify(event.node.req.socket.remoteAddress || "ip:n/a")} ${JSON.stringify(event.node.req.headers["user-agent"] || "user-agent:n/a")}`
    )

    // If frontend receives a request intended for the API, fail the request.
    if (event.node.req.url?.startsWith("/-/")) {
      event.node.res.statusCode = 404
      event.node.res.statusMessage = "backend request sent to frontend"
      event.node.res.end()
      return
    }
  })

  nitroApp.hooks.hook("render:response", (response) => {
    response.headers = {
      ...response.headers,
      "X-Frame-Options": "DENY",
      "X-Content-Type-Options": "nosniff",
      "Referrer-Policy": "no-referrer-when-downgrade",
      "Permissions-Policy": "clipboard-write=(self)",
    }

    if (!import.meta.dev) {
      response.headers = {
        ...response.headers,
        "Content-Security-Policy": [
          "default-src 'self'",
          "img-src * data:",
          "media-src * data:",
          "style-src 'self' 'unsafe-inline'",
          "object-src 'none'",
          "child-src 'none'",
          "frame-src 'none'",
          "worker-src 'none'",
          "script-src 'self' 'unsafe-inline' 'unsafe-eval'",
          "script-src-elem 'self' 'unsafe-inline' https://static.cloudflareinsights.com",
          "connect-src 'self' localhost:*",
        ].join(";"),
      }
    }

    delete response.headers["x-powered-by"]
  })
})
