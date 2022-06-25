import "virtual:windi.css"
import "@/main.css"
import { createPinia } from "pinia"
import { createApp } from "vue"
import { client } from "@/lib/api/client"
import router from "@/lib/core/router"
import App from "@/main.vue"
import urql from "@urql/vue"
import { MotionPlugin } from "@vueuse/motion"

const app = createApp(App)

app.use(urql, client)
app.use(createPinia())
app.use(router)
app.use(MotionPlugin)
app.mount("#app")
