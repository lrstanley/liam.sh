<script setup lang="ts">
import { client } from "#hey-api/client.gen"

const runtime = useRuntimeConfig()
const loading = useLoadingIndicator()

client.setConfig({
  baseURL: runtime.API_URL
    ? runtime.API_URL.replace(/\/$/, "")
    : runtime.public.API_URL.replace(/\/$/, ""),
  credentials: "include",
  retry: 3,
  retryDelay: 1000,
  headers: useRequestHeaders(["cookie"]), // Allows pass-through of cookies from origin browser, to API requests (mainly helpful when in SSR).
  onRequest: () => loading.start(),
  onResponse: (ctx) => (ctx.response.status < 300 ? loading.finish() : loading.finish({ error: true })),
})

const route = useRoute()

const title = computed(() =>
  route.meta.title ? `${route.meta.title} · Liam Stanley` : "Personal Website & Blog"
)

useHead({
  title: title,
  htmlAttrs: { lang: "en" },
  meta: [
    { name: "revisit-after", content: "10" },
    { name: "language", content: "en" },
    { name: "copyright", content: "© Liam Stanley" },
    { name: "darkreader-lock" },
  ],
  link: [{ rel: "icon", href: "/favicon.svg", type: "image/svg+xml" }],
})

useSeoMeta({
  author: "Liam Stanley",
  title: title,
  ogTitle: title,
  description: "Personal website, including blog posts, Git repositories, and more",
  ogDescription: "Personal website, including blog posts, Git repositories, and more",
  robots: "index, follow",
  viewport: {
    width: "device-width",
    initialScale: "1.0",
  },
})
</script>

<template>
  <NuxtRouteAnnouncer />
  <NuxtLoadingIndicator />
  <UApp>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </UApp>
</template>
