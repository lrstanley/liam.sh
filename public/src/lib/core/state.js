import { defineStore } from "pinia"

export const useState = defineStore("state", {
    state: () => {
        return {
            loading: false,
        }
    },
})
