/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in the
 * LICENSE file.
 */

import { createError, defineEventHandler, sendRedirect, useRuntimeConfig } from "nuxt/server"

// Resolved per request rather than through routeRules, as routeRules are fixed
// at build time, and the chat URL is only provided at runtime ($NUXT_CHAT_URL).
export default defineEventHandler((event) => {
  const url = String(useRuntimeConfig().CHAT_URL || "").trim()
  if (!url) {
    throw createError({ status: 404, statusText: "chat url not configured" })
  }

  return sendRedirect(event, url, 302)
})
