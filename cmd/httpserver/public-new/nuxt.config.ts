import { NaiveUiResolver } from "unplugin-vue-components/resolvers"
import Components from "unplugin-vue-components/vite"

export default defineNuxtConfig({
  $production: {},
  compatibilityDate: "2024-04-03",
  devtools: {
    enabled: true,
    // NaiveUI causes 10K LOC to be spammed in the console.
    componentInspector: false,
  },
  ssr: true,
  appConfig: {
    title: "liam.sh",
    icon: {
      class: "icon",
    },
  },
  runtimeConfig: {
    API_URL: "", // use $NUXT_API_URL
    public: {
      API_URL: "", // use $NUXT_PUBLIC_API_URL
    },
  },
  components: [
    {
      path: "~/components",
      pathPrefix: false,
    },
  ],
  app: {},
  css: ["~/assets/css/main.css"],
  devServer: {
    port: 8081,
  },
  experimental: {
    typedPages: true,
  },
  telemetry: false,
  vite: {
    server: {
      open: false,
      strictPort: true,
      proxy: {
        // TODO: can this be removed with the usage of nuxt request logic?
        "^/(-|security\\.txt|robots\\.txt)(/.*|$)": {
          target: "http://localhost:8080",
          xfwd: true,
          ws: true,
        },
      },
    },
    plugins: [
      Components({
        resolvers: [NaiveUiResolver()],
      }),
    ],
  },
  tailwindcss: {
    cssPath: "~/assets/css/main.css",
    config: {
      corePlugins: {
        preflight: false,
      },
      safelist: ["flex flex-row w-6 h-6 rounded-full ml-2 text-red-500"], // /admin/repo-needs-release
    },
  },
  imports: {
    presets: [
      {
        from: "naiveui",
        imports: ["useDialog", "useMessage", "useNotification", "useLoadingBar"],
      },
    ],
  },
  modules: [
    "@nuxtjs/tailwindcss",
    "@vueuse/nuxt",
    "@nuxt/icon",
    "nuxtjs-naive-ui",
    "@formkit/auto-animate",
  ],
})
