import { defineStore } from "pinia"
import { useStorage } from "@vueuse/core"

export const useState = defineStore("state", {
  state: () => {
    return useStorage("state", {
      base: null,
      history: [],
      sidebarCollapsed: false,
    })
  },
})
