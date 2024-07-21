/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import "@/css/main.css"
import router from "@/lib/core/router"
import App from "@/main.vue"
import { MotionPlugin } from "@vueuse/motion"
import { VueQueryPlugin } from "@tanstack/vue-query"
import { createClient } from "@hey-api/client-fetch"
import { createPinia } from "pinia"
import { createApp } from "vue"
import { loadingBar } from "@/lib/core/status"

const client = createClient({
  global: true,
  baseUrl: `${window.location.origin}/-`,
})

client.interceptors.request.use((request) => {
  loadingBar.start()
  return request
})

client.interceptors.response.use((response) => {
  response.status < 300 ? loadingBar.finish() : loadingBar.error()
  return response
})

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(MotionPlugin)
app.use(VueQueryPlugin)
app.mount("#app")
