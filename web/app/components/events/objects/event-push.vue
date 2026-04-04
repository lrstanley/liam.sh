<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["push-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <span class="text-purple-400">pushed</span>

    <EventHoverItem :href="props.event.repo.name + '/commit/' + props.event.payload.head"
      :value="props.event.payload.head.slice(0, 7)">
      <UIcon name="mdi:source-commit" />
      {{ props.event.payload.head.slice(0, 7) }}
    </EventHoverItem>
    to
    <EventLink :href="props.event.repo.name" />

    <EventBlame v-if="props.event.payload.ref">
      {{ props.event.payload.ref }}
    </EventBlame>
  </div>
</template>
