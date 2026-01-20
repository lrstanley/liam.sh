export default defineNuxtConfig({
  app: {
    layoutTransition: { name: "layout", mode: "out-in" },
  },
  modules: ["@vueuse/nuxt", "motion-v/nuxt", "@nuxt/ui", "@nuxtjs/seo", "nuxt-open-fetch"],
  compatibilityDate: "2026-01-19",
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
    clientBundle: { scan: true, sizeLimitKb: 0, includeCustomCollections: true },
  },
  devServer: { port: 8081 },
  experimental: { typedPages: true },
  telemetry: false,
  vite: {
    server: {
      open: false,
      strictPort: true,
    },
    plugins: [],
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
  openFetch: {
    disableNitroPlugin: true,
    disableNuxtPlugin: true,
    openAPITS: {
      rootTypes: true,
    },
    clients: {
      api: {
        schema: "../internal/database/ent/rest/openapi.json",
      }
    }
  },
})
