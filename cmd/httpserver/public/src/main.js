import { createApp } from "vue"
import App from "@/app.vue"
import router from "@/lib/core/router.js"
import { createPinia } from "pinia"
import { MotionPlugin } from "@vueuse/motion"
import "virtual:windi.css"

import urql from "@urql/vue"
import { client } from "@/lib/api/client"

const app = createApp(App)

app.use(urql, client)
app.use(createPinia())
app.use(router)
app.use(MotionPlugin)
app.mount("#app")
