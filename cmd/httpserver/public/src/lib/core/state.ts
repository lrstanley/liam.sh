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
})
