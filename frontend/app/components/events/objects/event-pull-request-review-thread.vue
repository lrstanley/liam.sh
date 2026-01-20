<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["webhook-pull-request-review-thread-resolved"] | Github["schemas"]["webhook-pull-request-review-thread-unresolved"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    marked a
    <EventLink :href="props.event.payload.pull_request.url" class="text-info-400 hover:text-info-500" value="thread" />

    as

    <span
      :class="{ 'text-success': props.event.payload.action == 'resolved', 'text-red-400': props.event.payload.action == 'unresolved' }">
      {{ props.event.payload.action }}
    </span>

    for pr

    <EventHoverItem :href="props.event.payload.pull_request.url" :value="'#' + props.event.payload.pull_request.number">
      {{ props.event.payload.pull_request.title }}
    </EventHoverItem>

    on
    <EventLink :href="props.event.repo.name" />

    <EventBlame>{{ props.event.payload.pull_request.title }}</EventBlame>
  </div>
</template>
