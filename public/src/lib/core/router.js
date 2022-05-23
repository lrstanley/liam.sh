import { createRouter, createWebHistory } from "vue-router"
import routes from "~pages"
import { titleCase } from "@/lib/core/util"
import { client } from "@/lib/api/client"
import { BaseDocument } from "@/lib/api"

const router = createRouter({
  history: createWebHistory("/"),
  routes,
})

router.beforeEach(async (to, from, next) => {
  const state = useState()

  if (from.name != to.name || JSON.stringify(from.params) != JSON.stringify(to.params)) {
    state.loading = true
  }

  let error

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
    window.location.href = `/api/auth/providers/github?next=${window.location.origin + to.path}`
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

router.afterEach((to) => {
  const state = useState()

  let title = to.path
    .split("/")
    .reverse()
    .filter((item) => item != "")

  if (title.length > 2) {
    title = title.slice(0, 2)
  }

  title = titleCase(title.reverse().join(" · ").replace(/-/g, " "))

  if (title.length < 2) {
    title = "Home"
  }

  document.title = `${title} · Liam Stanley`

  if (state.history.length > 4) {
    state.history.shift()
  }

  // remove any previous duplicates with the exact same path.
  for (let i = state.history.length - 1; i >= 0; i--) {
    if (state.history[i].path === to.path) {
      state.history.splice(i, 1)
    }
  }

  state.history.push({ title, path: to.path, timestamp: new Date().toISOString() })

  setTimeout(() => {
    state.loading = false
  }, 100)
  return
})

export default router
