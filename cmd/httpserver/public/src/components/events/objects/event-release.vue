<script setup lang="ts">
import type { GithubEvent } from "@/lib/api"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action)
const release = ref<Record<string, any>>(props.event.payload.release)
</script>

<template>
  <div>
    <span class="text-lime-500">{{ action }}</span>

    release from tag
    <EventHoverItem :href="release.html_url" :value="release.tag_name" class="truncate">
      {{ release.name }}
    </EventHoverItem>
    on
    <EventLink :href="repo.name" />

    <EventBlame>{{ release.name }}</EventBlame>
  </div>
</template>
