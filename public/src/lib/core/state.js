import { useStorage } from "@vueuse/core"
import { defineStore } from "pinia"

export const useState = defineStore("state", {
  state: () => {
    return useStorage("state", {
      loading: false,
      auth: null,
      me: null,
      history: [],
      sidebarCollapsed: false,
    })
  },
})
