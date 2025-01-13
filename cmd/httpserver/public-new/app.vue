<script setup lang="ts">
import { darkTheme } from "naive-ui"
import { client } from "./utils/http/sdk.gen"

useHead({
  titleTemplate: (chunk) => (chunk ? `${chunk} · Liam Stanley` : "Liam Stanley"),
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
  title: "Personal Website & Blog - Liam Stanley",
  ogTitle: "Personal Website & Blog - Liam Stanley",
  description: "Personal website, including blog posts, Git repositories, and more",
  ogDescription: "Personal website, including blog posts, Git repositories, and more",
  robots: "index, follow",
  viewport: {
    width: "device-width",
    initialScale: "1.0",
  },
})

const runtime = useRuntimeConfig()

let url: string
if (runtime.API_URL) {
  url = runtime.API_URL as string
} else if (runtime.public.API_URL) {
  url = runtime.public.API_URL as string
} else {
  url = `${useRequestURL().origin}/-`
}

client.setConfig({ baseURL: url })
</script>

<template>
  <n-config-provider :theme="darkTheme" abstract preflight-style-disabled inline-theme-disabled>
    <NuxtRouteAnnouncer />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </n-config-provider>
</template>
