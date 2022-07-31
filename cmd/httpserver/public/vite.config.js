/**
 * @type {import('vite').UserConfig}
 */
import { defineConfig } from "vite"

import path from "path"
import Vue from "@vitejs/plugin-vue"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import { NaiveUiResolver, VueUseComponentsResolver } from "unplugin-vue-components/resolvers"
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
      dirs: "src/pages",
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
      dts: true,
      directives: true,
      directoryAsNamespace: false,
      resolvers: [
        VueUseComponentsResolver(),
        NaiveUiResolver(),
        IconsResolver({ componentPrefix: "i", enabledCollections: ["mdi"] }),
      ],
    }),
    AutoImport({
      dts: true,
      imports: [
        "vue",
        "vue/macros",
        "vue-router",
        "@vueuse/core",
        {
          pinia: ["storeToRefs"],
        },
        {
          "@/lib/core/state.ts": ["useState"],
        },
      ],
      resolvers: [IconsResolver({ componentPrefix: "icon", enabledCollections: ["mdi"] })],
      eslintrc: {
        enabled: true,
      },
    }),
    Icons({
      autoInstall: true,
      defaultClass: "icon",
    }),
  ],
  base: "/",
  build: {
    sourcemap: true,
    emptyOutDir: true,
    mode: "production",
  },
  preview: {
    port: 8081,
    mode: "production",
  },
  server: {
    base: "/",
    mode: "development",
    port: 8081,
    strictPort: true,
    proxy: {
      "^/(-|security\\.txt|robots\\.txt)(/.*|$)": {
        target: "http://localhost:8080",
        xfwd: true,
      },
    },
  },
})
