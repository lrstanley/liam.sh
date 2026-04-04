<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"

definePageMeta({
  title: "Gists",
  layout: "terminal",
})

const { data, error } = await useApi('/github-gists', {
  query: {
    page: 1,
    per_page: 100,
    "public.eq": true,
    sort: "name",
  }
})
if (error.value) throw error.value
const gists = computed(() => data.value?.content ?? [])
</script>

<template>
  <ContainerIde>
    <div class="relative overflow-x-hidden size-full grow basis-0">
      <motion as="div" :initial="{ opacity: 0 }" :animate="{ opacity: 1 }" :exit="{ opacity: 0 }"
        :transition="{ delay: ((i % 100) + 1) * 0.015 }" v-for="(gist, i) in gists" :key="gist.id"
        class="flex flex-row items-center flex-auto px-1 text-sm transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-muted border-b-gray-100">
        <NuxtLink :href="gist.owner.html_url" target="_blank">
          <UAvatar size="3xs" :src="gist.owner.avatar_url + '&s=40'" class="mr-1 align-middle"
            :alt="gist.owner.login" />
        </NuxtLink>

        <div class="flex items-center gap-2 truncate grow">
          <span class="text-default">{{ gist.name }}</span>
          <span class="text-secondary">{{ gist.language || "Unknown" }}</span>
          <NuxtLink target="_blank" class="text-primary" :href="'/-/gh/g/' + gist.name">[open]</NuxtLink>
          <NuxtLink target="_blank" class="text-primary" :href="gist.html_url">[gh]</NuxtLink>
          <NuxtLink target="_blank" class="text-warning hover:text-warning-600" :href="gist.raw_url">
            [raw]
          </NuxtLink>

          <span class="flex-1 lowercase truncate select-none text-(--ui-text)/30" :title="gist.description">
            {{ gist.description || "No description set" }}
          </span>
        </div>

        <div class="flex-none">
          <EventHoverItem placement="left">
            <template #value>
              <UIcon name="mdi:clock-time-two-outline" class="timestamp" />
            </template>

            {{ useTimeAgo(gist.updated_at).value }}
          </EventHoverItem>
        </div>
      </motion>
    </div>
  </ContainerIde>
</template>
