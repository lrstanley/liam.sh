// @ts-check - enable TS check for js file
import { defineConfig } from "windicss/helpers"

export default defineConfig({
    theme: {
        extend: {},
    },
    plugins: [require("windicss/plugin/typography")],
    attributify: {
        prefix: "w:",
    },
    // shortcuts: {
    //     "text-gradient": "",
    // },
})
