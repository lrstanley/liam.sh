<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["public-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    made
    <EventLink :href="props.event.repo.name" />
    <span class="text-amber-300">public</span>
  </div>
</template>
