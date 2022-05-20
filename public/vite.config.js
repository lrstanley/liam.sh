import path from "path"
import { defineConfig } from "vite"

import Vue from "@vitejs/plugin-vue"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import WindiCSS from "vite-plugin-windicss"
import { NaiveUiResolver } from "unplugin-vue-components/resolvers"
import Pages from "vite-plugin-pages"
import Icons from "unplugin-icons/vite"
import IconsResolver from "unplugin-icons/resolver"

export default defineConfig({
  resolve: {
    alias: {
      "@/": `${path.resolve(__dirname, "src")}/`,
    },
  },
  publicDir: `${path.resolve(__dirname, "src")}/assets`,
  plugins: [
    Pages({
      dirs: "src/views",
      routeBlockLang: "yaml",
      extendRoute(route) {
        // route, parent
        if (route.path.startsWith("/admin")) {
          route = {
            ...route,
            meta: {
              auth: true,
            },
          }
        }

        return route
      },
    }),
    Vue({
      reactivityTransform: true,
    }),
    Components({
      dts: false,
      directoryAsNamespace: false,
      resolvers: [
        NaiveUiResolver(),
        IconsResolver({ componentPrefix: "i", enabledCollections: ["mdi"] }),
      ],
    }),
    AutoImport({
      dts: false,
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
      ],
      resolvers: [IconsResolver({ componentPrefix: "icon", enabledCollections: ["mdi"] })],
      eslintrc: {
        enabled: true,
      },
      exclude: [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/],
    }),
    WindiCSS(),
    Icons({
      autoInstall: true,
    }),
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
