import { useStorage } from "@vueuse/core"
import { defineStore } from "pinia"

export const useState = defineStore("state", {
  state: () => {
    return useStorage("state", {
      base: null,
      loading: false,
      history: [],
      sidebarCollapsed: false,
    })
  },
})
