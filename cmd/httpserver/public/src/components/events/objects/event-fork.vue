<script setup lang="ts">
import type { GithubEvent } from "@/lib/http/types.gen"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const fork = ref<Record<string, any>>(props.event.payload.forkee)
</script>

<template>
  <div>
    <span class="text-purple-400">forked</span>
    <EventLink :href="repo.name" />
    to
    <EventLink :href="fork.full_name" />

    <EventBlame>{{ fork.description }}</EventBlame>
  </div>
</template>
