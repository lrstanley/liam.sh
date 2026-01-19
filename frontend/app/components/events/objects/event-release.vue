<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["release-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <span class="text-success">{{ props.event.payload.action }}</span>

    release from tag
    <EventHoverItem :href="props.event.payload.release.html_url" :value="props.event.payload.release.tag_name"
      class="truncate block">
      {{ props.event.payload.release.name }}
    </EventHoverItem>
    on
    <EventLink :href="props.event.repo.name" />

    <EventBlame>{{ props.event.payload.release.name }}</EventBlame>
  </div>
</template>
