import Vue from "@vitejs/plugin-vue"
import path from "path"
import { visualizer } from "rollup-plugin-visualizer"
import AutoImport from "unplugin-auto-import/vite"
import IconsResolver from "unplugin-icons/resolver"
import Icons from "unplugin-icons/vite"
import { NaiveUiResolver, VueUseComponentsResolver } from "unplugin-vue-components/resolvers"
import Components from "unplugin-vue-components/vite"
import { VueRouterAutoImports } from "unplugin-vue-router"
import VueRouter from "unplugin-vue-router/vite"
import { defineConfig } from "vite"
import Layouts from "vite-plugin-vue-layouts"

const icons = IconsResolver({
  componentPrefix: "i",
  enabledCollections: ["mdi", "logos"],
})

export default defineConfig({
  resolve: {
    alias: {
      "@/": `${path.resolve(__dirname, "src")}/`,
    },
  },
  publicDir: `${path.resolve(__dirname, "src")}/assets`,
  plugins: [
    visualizer({
      filename: "./dist/stats.html",
    }),
    VueRouter({
      routesFolder: "src/pages",
    }),
    Vue({ include: [/\.vue$/] }),
    Layouts({
      layoutsDirs: "src/layouts",
      defaultLayout: "default",
    }),
    Components({
      dts: true,
      directives: true,
      directoryAsNamespace: false,
      resolvers: [VueUseComponentsResolver(), NaiveUiResolver(), icons],
    }),
    AutoImport({
      dts: true,
      imports: [
        "vue",
        "@vueuse/core",
        {
          "@/lib/core/state": ["useState"],
          "@/lib/util/request": ["usePagination", "resetPagination", "unwrapErrors"],
          "@tanstack/vue-query": ["useQuery", "useMutation", "useQueryClient", "keepPreviousData"],
        },
        VueRouterAutoImports,
      ],
      resolvers: [icons],
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
    sourcemap: "hidden",
    emptyOutDir: true,
    rollupOptions: {
      output: {
        chunkFileNames: "chunk-[hash].js",
        entryFileNames: "entry-[hash].js",
        assetFileNames: "assets/[hash][extname]",
      },
    },
  },
  preview: {
    port: 8081,
    open: false,
  },
  server: {
    port: 8081,
    open: false,
    strictPort: true,
    proxy: {
      "^/(-|security\\.txt|robots\\.txt)(/.*|$)": {
        target: "http://localhost:8080",
        xfwd: true,
        ws: true,
      },
    },
  },
})
