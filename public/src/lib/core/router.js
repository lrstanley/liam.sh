import { createRouter, createWebHistory } from "vue-router"

import { api } from "@/lib/http"

const routes = [
    {
        path: "/",
        name: "index",
        component: () => import("@/views/Index.vue"),
        meta: {
            title: "Home",
        },
    },
    {
        path: "/about",
        name: "about",
        component: () => import("@/views/About.vue"),
        meta: {
            title: "About",
            auth: true,
        },
    },
    {
        path: "/admin",
        name: "admin",
        component: () => import("@/views/admin/Index.vue"),
        meta: {
            title: "Admin",
            auth: true,
        },
    },
    {
        path: "/admin/edit",
        name: "admin-edit",
        component: () => import("@/views/admin/EditPost.vue"),
        meta: {
            title: "Edit Post",
            auth: true,
        },
    },
    {
        path: "/:pathMatch(.*)",
        name: "catchall",
        // redirect: "/404",
        component: () => import("@/views/Error.vue"),
    },
]

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
        api.get("/auth/self")
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
        api.get("/gh/me")
            .then(({ data }) => {
                state.me = data.user
                return resolve()
            })
            .catch(({ response }) => {
                console.log(response)
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
    if (to.meta.title) {
        document.title = `${to.meta.title} · Liam Stanley`
    }

    setTimeout(() => {
        const state = useState()
        state.loading = false
    }, 100)
    return
})

export default router
