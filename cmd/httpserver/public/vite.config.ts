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
import codegen from "vite-plugin-graphql-codegen"
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
    codegen({
      enableWatcher: true,
      config: {
        errorsOnly: true,
        schema: "./../../../internal/graphql/schema/*.gql",
        documents: "./src/lib/api/*.gql",
        generates: {
          "./src/lib/api/graphql.ts": {
            plugins: ["typescript", "typescript-operations", "typescript-vue-urql"],
            config: {
              preResolveTypes: true,
              nonOptionalTypename: true,
              skipTypeNameForRoot: true,
              useTypeImports: true,
              inputMaybeValue: "T | Ref<T> | ComputedRef<T>",
            },
          },
        },
      },
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
