import tailwindcss from "@tailwindcss/vite"

export default defineNuxtConfig({
  modules: [
    "@hey-api/nuxt",
    "@vueuse/nuxt",
    "@formkit/auto-animate/nuxt", // TODO: remove if motion-v ends up working.
    "@nuxt/ui-pro",
    "nuxt-typed-router",
  ],
  future: {
    compatibilityVersion: 4,
  },
  compatibilityDate: "2025-01-13",
  devtools: {
    // viteInspect: false,
    // componentInspector: false,
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
  css: ["~/assets/main.css"],
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
    },
    plugins: [tailwindcss()],
  },
  heyApi: {
    config: {
      input: {
        path: "../internal/database/ent/rest/openapi.json",
      },
      plugins: [
        "@hey-api/client-nuxt",
        "@hey-api/schemas",
        "@hey-api/sdk",
        "zod",
        {
          name: "@hey-api/typescript",
          enums: "javascript",
        },
      ],
    },
  },
})
