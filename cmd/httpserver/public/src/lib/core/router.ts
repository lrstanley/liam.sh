/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { setupLayouts } from "virtual:generated-layouts"
import { createRouter, createWebHistory } from "vue-router/auto"
import { BaseDocument, client } from "@/lib/api"
import { useState } from "@/lib/core/state"
import { loadingBar } from "@/lib/core/status"
import { titleCase } from "@/lib/util"

import type { RouteRecordRaw } from "vue-router/auto"
import type { CombinedError } from "@urql/vue"

type RouteCallback = (route: RouteRecordRaw) => RouteRecordRaw

function recurseRoute(route: RouteRecordRaw, cb: RouteCallback): RouteRecordRaw {
  if (route.children) {
    for (let i = 0; i < route.children.length; i++) {
      route.children[i] = recurseRoute(route.children[i], cb)
    }
  }

  return cb(route)
}

const router = createRouter({
  history: createWebHistory("/"),
  extendRoutes(routes) {
    return routes.map((r) =>
      recurseRoute(r, (route) => {
        if (route.path.startsWith("/admin")) {
          route = {
            ...route,
            meta: {
              auth: true,
              admin: true,
              ...route.meta,
            },
          }
        }

        return route.component ? setupLayouts([route])[0] : route
      })
    )
  },
})

router.beforeEach(async (to, from, next) => {
  const state = useState()

  if (from.name != to.name || JSON.stringify(from.params) != JSON.stringify(to.params)) {
    loadingBar.start()
  }

  let error: CombinedError

  if (state.base == null || (from.path == "/" && from.name == undefined)) {
    await client
      .query(BaseDocument, {}, { requestPolicy: "network-only" })
      .toPromise()
      .then((resp) => {
        state.base = resp.data

        if (resp.error !== null) {
          error = resp.error
        }
      })
  }

  if (to.meta.auth == true && state.base.self == null) {
    window.location.href = `/-/auth/providers/github?next=${window.location.origin + to.path}`
    return
  }

  if (
    error !== undefined &&
    !error.graphQLErrors?.some((e) => e.path?.includes("self")) &&
    to.name !== "catchall"
  ) {
    console.log(error)
    next({ name: "catchall", params: { catchall: error.name } })
    return
  }

  next()
})

router.afterEach((to, from) => {
  const state = useState()

  let args = to.path
    .split("/")
    .reverse()
    .filter((item) => item != "")

  if (args.length > 2) {
    args = args.slice(0, 2)
  }

  for (let i = 0; i < args.length; i++) {
    if (args[i] == "p") {
      args[i] = "Posts"
    }
  }

  let title = titleCase(args.reverse().join(" · ").replace(/-/g, " "))

  if (title.length < 2) {
    title = "Home"
  }

  document.title = `${title} · Liam Stanley` // TODO: switch to dynamic name.
  state.addToHistory({ title, path: to.path, timestamp: new Date().toISOString() })

  // Scroll to anchor, just in case the page happens to not render fast enough.
  nextTick(() => {
    loadingBar.finish()

    if (location.hash && !to.meta.disableAnchor) {
      const el = document.getElementById(location.hash.slice(1))
      if (el) {
        el.scrollIntoView({ behavior: "smooth" })
      }
    } else if (from.name !== undefined && !to.meta.disableAnchor) {
      window.scrollTo(0, 0)
    }
  })
  return
})

export default router
