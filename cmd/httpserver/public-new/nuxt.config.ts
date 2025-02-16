import tailwindcss from "@tailwindcss/vite"

export default defineNuxtConfig({
  modules: [
    "@vueuse/nuxt",
    "@formkit/auto-animate",
    "@nuxt/ui-pro",
    // "@nuxtjs/html-validator"
    // "@nuxtjs/fontaine",
  ],
  future: {
    compatibilityVersion: 4,
  },
  // features: {
  //   inlineStyles: false,
  // },
  compatibilityDate: "2025-01-13",
  devtools: {
    viteInspect: false,
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
  ui: {
    primary: "green",
    grey: "neutral",
  },
  uiPro: {},
  build: {
    // TODO: https://github.com/hey-api/openapi-ts/issues/1660
    transpile: ["@hey-api/client-nuxt"],
  },
})
