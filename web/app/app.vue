<script setup lang="ts">
const loading = useLoadingIndicator()
const nuxtApp = useNuxtApp()

nuxtApp.hook('openFetch:onRequest', (ctx) => {
  if (loading.isLoading.value) return
  loading.start({ force: true })
})
nuxtApp.hook('openFetch:onResponse', (ctx) => ctx.response.status < 300 ? loading.finish() : loading.finish({ error: true }))

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
