<script setup lang="ts">
useHead({
  titleTemplate: (chunk) => (chunk ? `${chunk} · Liam Stanley` : "Liam Stanley"),
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
  title: "Personal Website & Blog",
  ogTitle: "Personal Website & Blog",
  description: "Personal website, including blog posts, Git repositories, and more",
  ogDescription: "Personal website, including blog posts, Git repositories, and more",
  robots: "index, follow",
  viewport: {
    width: "device-width",
    initialScale: "1.0",
  },
})

setHTTPClientBaseURL()
</script>

<template>
  <NuxtRouteAnnouncer />
  <NuxtLoadingIndicator />
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
