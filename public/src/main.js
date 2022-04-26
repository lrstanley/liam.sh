import { createApp } from "vue"
import App from "@/App.vue"
import router from "@/lib/core/router.js"
import { createPinia } from "pinia"
import { MotionPlugin } from "@vueuse/motion"
import "virtual:windi.css"

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(MotionPlugin)
app.mount("#app")
