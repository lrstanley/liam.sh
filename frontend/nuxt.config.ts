import tailwindcss from "@tailwindcss/vite"

export default defineNuxtConfig({
  app: {
    layoutTransition: { name: "layout", mode: "out-in" },
  },
  modules: ["@hey-api/nuxt", "@vueuse/nuxt", "motion-v/nuxt", "@nuxt/ui-pro", "@nuxtjs/seo"],
  future: { compatibilityVersion: 4 },
  compatibilityDate: "2025-05-03",
  devtools: {
    enabled: false,
    viteInspect: false,
    componentInspector: false,
  },
  ssr: true,
  runtimeConfig: {
    API_URL: "http://localhost:8080/-", // use $NUXT_API_URL
    CHAT_URL: "", // use $NUXT_CHAT_URL
    public: {
      API_URL: "http://localhost:8080/-", // use $NUXT_PUBLIC_API_URL
    },
  },
  components: [
    {
      path: "~/components",
      pathPrefix: false,
    },
  ],
  css: ["~/assets/main.css"],
  icon: {
    fallbackToApi: false,
    serverBundle: "local",
    clientBundle: { scan: true },
  },
  devServer: { port: 8081 },
  experimental: { typedPages: true },
  telemetry: false,
  vite: {
    server: {
      open: false,
      strictPort: true,
    },
    plugins: [tailwindcss()],
  },
  routeRules: {
    "/security.txt": { redirect: "/-/security.txt" },
    "/chat": { redirect: process.env.NUXT_CHAT_URL },
  },
  robots: { disallow: "/admin/*" },
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
