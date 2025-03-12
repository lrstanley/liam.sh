import tailwindcss from "@tailwindcss/vite"

export default defineNuxtConfig({
  modules: [
    "@hey-api/nuxt",
    "@vueuse/nuxt",
    "@formkit/auto-animate",
    "@nuxt/ui-pro",
    // "@nuxtjs/html-validator"
  ],
  future: {
    compatibilityVersion: 4,
  },
  compatibilityDate: "2025-01-13",
  devtools: {
    viteInspect: false,
    componentInspector: false,
  },
  ssr: true,
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
  css: ["~/assets/css/main.css"],
  devServer: {
    port: 8081,
  },
  experimental: {
    typedPages: true,
  },
  telemetry: false,
  routeRules: {
    "/admin/**": {
      ssr: false,
    },
  },
  vite: {
    server: {
      open: false,
      strictPort: true,
    },
    plugins: [tailwindcss()],
  },
  heyApi: {
    config: {
      experimentalParser: true,
      input: {
        path: "../../../internal/database/ent/rest/openapi.json",
      },
      output: {
        path: "app/utils/http",
        clean: true,
        indexFile: true,
      },
      plugins: [
        "@hey-api/client-nuxt",
        "@hey-api/schemas",
        {
          name: "@hey-api/sdk",
          transformer: true,
        },
        {
          enums: "javascript",
          name: "@hey-api/typescript",
        },
        {
          name: "@hey-api/transformers",
          dates: true,
        },
      ],
    },
  },
})
