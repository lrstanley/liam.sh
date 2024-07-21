<script setup lang="ts">
import type { GithubEvent } from "@/lib/http/types.gen"

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
