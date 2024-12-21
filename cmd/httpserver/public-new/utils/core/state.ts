/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { defineStore } from "pinia"
import { useStorage } from "@vueuse/core"
import type { User, GithubUser } from "@/lib/http/types.gen"

export interface History {
  title: string
  path: string
  timestamp: string
}

export interface State {
  user: User | null
  githubUser: GithubUser | null
  history: History[]
  sidebarCollapsed: boolean
}

export const useState = defineStore("state", {
  state: () => {
    return useStorage("state", {
      user: null,
      githubUser: null,
      history: [],
      sidebarCollapsed: false,
    } as State)
  },
  actions: {
    addToHistory(item: History) {
      // Truncate to a max size.
      if (this.history.length > 4) {
        this.history.shift()
      }

      // Remove any previous duplicates with the exact same path.
      for (let i = this.history.length - 1; i >= 0; i--) {
        if (this.history[i].path === item.path) {
          this.history.splice(i, 1)
        }
      }

      this.history.push(item)
    },
  },
})
