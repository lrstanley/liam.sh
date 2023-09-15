<template>
  <div>
    marked a
    <EventLink :href="thread.html_url" class="text-cyan-400 hover:text-cyan-500" value="thread" />

    as

    <span :class="{ 'text-lime-500': action == 'resolved', 'text-red-400': action == 'unresolved' }">
      {{ action }}
    </span>

    for pr

    <EventHoverItem :href="pr.html_url" :value="'#' + pr.number">
      {{ pr.title }}
    </EventHoverItem>

    on
    <EventLink :href="repo.name" />

    <EventBlame>{{ pr.title }}</EventBlame>
  </div>
</template>

<script setup lang="ts">
import type { GithubEvent } from "@/lib/api"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action)
const thread = ref<Record<string, any>>(props.event.payload.thread)
const pr = ref<Record<string, any>>(props.event.payload.pull_request)
</script>
