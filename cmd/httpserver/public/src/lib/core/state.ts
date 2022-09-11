import { defineStore } from "pinia"
import { useStorage } from "@vueuse/core"

import type { BaseQuery } from "@/lib/api"

export interface History {
  title: string
  path: string
  timestamp: string
}

export interface State {
  base: BaseQuery | null
  history: History[]
  sidebarCollapsed: boolean
}

export const useState = defineStore("state", {
  state: () => {
    return useStorage("state", {
      base: null,
      history: [],
      sidebarCollapsed: false,
    } as State)
  },
  actions: {
    addToHistory(item: History) {
      // remove any previous duplicates with the exact same path.
      for (let i = this.history.length - 1; i >= 0; i--) {
        if (this.history[i].path === item.path) {
          this.history.splice(i, 1)
        }
      }

      this.history.push(item)
    },
  },
})
