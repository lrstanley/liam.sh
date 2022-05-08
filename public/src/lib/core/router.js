import { createRouter, createWebHistory } from "vue-router"
import routes from "~pages"
import { api } from "@/lib/http"
import { titleCase } from "@/lib/core/util"

const router = createRouter({
  history: createWebHistory("/"),
  routes,
})

router.beforeEach(async (to, from, next) => {
  const state = useState()
  state.loading = true

  if (state.auth !== null) {
    return next()
  }

  const self = new Promise((resolve, reject) => {
    api
      .get("/auth/self")
      .then(({ data }) => {
        state.auth = data.auth
        return resolve()
      })
      .catch(({ response }) => {
        if (response.status === 401) {
          state.auth = null

          if (to.meta.auth == true) {
            window.location.href = `/api/auth/providers/github?next=${window.location.origin + to.path}`
            return
          }
          return resolve()
        } else {
          if (to.name == "catchall") return resolve()

          return reject(response)
        }
      })
  })

  const me = new Promise((resolve, reject) => {
    api
      .get("/gh/me")
      .then(({ data }) => {
        state.me = data.user
        return resolve()
      })
      .catch(({ response }) => {
        if (to.name == "catchall") return resolve()

        return reject(response)
      })
  })

  await Promise.all([self, me])
    .then(() => {
      return next()
    })
    .catch((error) => {
      return next({ name: "catchall", params: { pathMatch: error.status } })
    })
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
