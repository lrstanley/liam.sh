import { createRouter, createWebHistory } from "vue-router"
import Index from "@/views/Index.vue"

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

router.beforeEach(() => {
    const state = useState()
    state.loading = true
    return
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
