import path from "path"
import { defineConfig } from "vite"

import Vue from "@vitejs/plugin-vue"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import WindiCSS from "vite-plugin-windicss"
import { NaiveUiResolver } from "unplugin-vue-components/resolvers"
import { IconComponentResolver } from "./src/lib/resolvers/icon-component-resolver.ts"
import IconResolver from "./src/lib/resolvers/icon-resolver.ts"

export default defineConfig({
    resolve: {
        alias: {
            "@/": `${path.resolve(__dirname, "src")}/`,
        },
    },
    publicDir: `${path.resolve(__dirname, "src")}/assets`,
    plugins: [
        Vue({
            reactivityTransform: true,
        }),
        Components({
            directoryAsNamespace: true,
            resolvers: [NaiveUiResolver(), IconComponentResolver({ pkg: "@vicons/ionicons5" })],
        }),
        AutoImport({
            imports: [
                "vue",
                "vue/macros",
                "vue-router",
                "@vueuse/core",
                {
                    pinia: ["storeToRefs"],
                },
                {
                    "@/lib/core/state.js": ["useState"],
                },
                IconResolver("@vicons/ionicons5"),
            ],
            eslintrc: {
                enabled: true,
            },
            exclude: [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/],
        }),
        WindiCSS(),
    ],
    server: {
        port: 8081,
        strictPort: true,
        proxy: {
            "^/api/.*": {
                target: "http://localhost:8080",
                xfwd: true,
            },
        },
        force: true,
    },
})
