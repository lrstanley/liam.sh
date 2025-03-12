<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import { listGithubGists } from "@/utils/http/sdk.gen"

useHead({ title: "Gists" })
definePageMeta({
  layout: "terminal",
})

const { data, error } = await listGithubGists({
  composable: "useFetch",
  query: {
    page: 1,
    per_page: 250,
    "public.eq": true,
    sort: "name",
  },
})
if (error.value) throw error.value
const gists = computed(() => data.value?.content ?? [])
</script>

<template>
  <ContainerIde>
    <div class="relative overflow-x-hidden size-full grow basis-0" v-auto-animate>
      <div
        v-for="(gist, i) in gists"
        :key="gist.id"
        class="flex flex-row items-center flex-auto px-1 text-sm transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-zinc-400 border-b-gray-100"
      >
        <a :href="gist.owner.html_url" target="_blank">
          <UAvatar
            size="3xs"
            :src="gist.owner.avatar_url + '&s=40'"
            class="mr-1 align-middle"
            :alt="gist.owner.login"
          />
        </a>

        <div class="flex items-center gap-2 truncate grow">
          <span class="text-zinc-300">{{ gist.name }}</span>
          <span class="text-violet-500">{{ gist.language || "Unknown" }}</span>
          <a target="_blank" class="text-emerald-500" :href="'/-/gh/g/' + gist.name">[open]</a>
          <a target="_blank" class="text-emerald-500" :href="gist.html_url">[gh]</a>
          <a target="_blank" class="text-amber-500 hover:text-amber-600" :href="gist.raw_url">[raw]</a>

          <span class="flex-1 lowercase truncate select-none text-zinc-600" :title="gist.description">
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
      </div>
    </div>
  </ContainerIde>
</template>

<style scoped>
@reference "@/assets/css/main.css";
</style>
