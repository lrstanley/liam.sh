<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import type { SchemaPostRead } from '#open-fetch-schemas/api'

const props = defineProps<{
  value: SchemaPostRead
  linkable?: boolean
}>()

const post = ref(props.value)
const drawerActive = ref(false)
</script>

<template>
  <div class="flex flex-row gap-4 py-4">
    <UAvatar :src="post.edges.author.avatar_url + '&s=40'" class="mt-1" />
    <div class="flex flex-col w-full gap-1">
      <div class="flex flex-row">
        <NuxtLink
          class="text-[1.4em] md:text-[1.5em] text-gradient bg-linear-to-br from-pink-500 via-red-500 to-yellow-500"
          :to="{ name: 'p-slug', params: { slug: post.slug } }">
          {{ post.title }}
        </NuxtLink>
      </div>

      <div class="flex flex-row gap-2 text-sm">
        <div class="inline-flex flex-row items-center">
          <UIcon name="mdi:update" class="mr-1 text-secondary-400" />
          <span class="italic text-muted">
            Published
            {{ useTimeAgo(post.published_at).value }}
            by
            <NuxtLink :href="post.edges.author.html_url" target="_blank">{{ post.edges.author.name }}</NuxtLink>
          </span>
        </div>
      </div>

      <div v-html="post.summary" class="my-2" />

      <div class="flex flex-row gap-2">
        <div v-if="post.edges.labels" class="flex flex-row flex-wrap gap-1">
          <LabelObject v-for="label in post.edges.labels" :key="label.id" :value="label" route="/posts" linkable
            class="hidden md:flex" />

          <UDrawer v-model:open="drawerActive" direction="bottom">
            <UButton label="Post labels" color="neutral" variant="subtle" size="xs" trailing-icon="lucide:chevron-up"
              class="flex md:hidden" />

            <template #content>
              <div class="flex flex-wrap flex-auto gap-1 p-4">
                <LabelObject v-for="label in post.edges.labels" :key="label.id" :value="label" route="/posts" linkable
                  @click="drawerActive = false" />
              </div>
            </template>
          </UDrawer>
        </div>

        <!-- views -->
        <div class="flex flex-col ml-auto shrink-0">
          <StatsButton>
            {{ post.view_count.toLocaleString() }}
            {{ post.view_count === 1 ? "view" : "views" }}
          </StatsButton>

          <!-- TODO: show draft & edit badges/buttons -->
        </div>
      </div>
    </div>
  </div>
</template>
