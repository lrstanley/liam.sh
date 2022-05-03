import { createRouter, createWebHistory } from "vue-router"
import Index from "@/views/Index.vue"

import { api } from "@/lib/http/index.js"

const routes = [
    {
        path: "/",
        name: "index",
        component: Index,
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

    await api
        .get("/auth/self")
        .then((resp) => {
            console.log(resp)
            state.auth = resp.data.auth
            return next()
        })
        .catch((error) => {
            console.log(error.response)
            if (error.response.status === 401) {
                state.auth = null

                if (to.meta.auth == true) {
                    window.location.href = `/api/auth/providers/github?next=${window.location.origin + to.path}`
                    return
                }
                return next()
            } else {
                if (to.name == "catchall") return next()

                return next({ name: "catchall" })
            }
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
