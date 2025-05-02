import tailwindcss from "@tailwindcss/vite"

export default defineNuxtConfig({
  app: {
    layoutTransition: { name: "layout", mode: "out-in" },
  },
  modules: [
    "@hey-api/nuxt",
    "@vueuse/nuxt",
    "@formkit/auto-animate/nuxt", // TODO: remove if motion-v ends up working.
    "@nuxt/ui-pro",
    "nuxt-typed-router",
    "@nuxtjs/seo",
  ],
  future: {
    compatibilityVersion: 4,
  },
  compatibilityDate: "2025-01-13",
  devtools: {
    enabled: false,
    viteInspect: false,
    componentInspector: false,
  },
  ssr: true,
  runtimeConfig: {
    API_URL: "http://localhost:8081/-", // use $NUXT_API_URL
    CHAT_URL: "", // use $NUXT_CHAT_URL
    public: {
      API_URL: "/-", // use $NUXT_PUBLIC_API_URL
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
      proxy: {
        "^/-.*": {
          target: "http://localhost:8080",
          xfwd: true,
          ws: true,
        },
      },
    },
    plugins: [tailwindcss()],
  },
  routeRules: {
    "/security.txt": { redirect: "/-/security.txt" },
    "/chat": { redirect: process.env.NUXT_CHAT_URL },
  },
  robots: {
    disallow: "/admin/*",
  },
  sitemap: {
    exclude: ["/admin/**"],
    sources: ["/api/__sitemap__/urls"],
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
