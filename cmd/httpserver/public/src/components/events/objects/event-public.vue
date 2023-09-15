<script setup lang="ts">
import type { GithubEvent } from "@/lib/api"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
</script>

<template>
  <div>
    made
    <EventLink :href="repo.name" />
    public
  </div>
</template>
