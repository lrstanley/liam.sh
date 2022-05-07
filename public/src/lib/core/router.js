import { createRouter, createWebHistory } from "vue-router"
import routes from "~pages"
import { api } from "@/lib/http"

function titleCase(str) {
  return str
    .toLowerCase()
    .split(" ")
    .map(function (word) {
      return word.charAt(0).toUpperCase() + word.slice(1)
    })
    .join(" ")
}

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
  let title = to.path
    .split("/")
    .at(-1)
    .replace(/-/g, " ")
    .replace(/:([^a-zA-Z0-9]+)/g, (p1) => {
      return ` ${to.params[p1]} `
    })

  if (title.length < 2) {
    title = "Home"
  }

  document.title = `${titleCase(title)} · Liam Stanley`

  setTimeout(() => {
    const state = useState()
    state.loading = false
  }, 100)
  return
})

export default router
