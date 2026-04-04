<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["watch-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div class="text-amber-300">
    {{ props.event.payload.action == "started" ? "starred" : "unstarred" }}
    <EventLink :href="props.event.repo.name" />
  </div>
</template>
