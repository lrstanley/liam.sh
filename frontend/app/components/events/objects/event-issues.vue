<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["issues-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <div class="text-info-400">
      <span v-if="['opened', 'edited', 'closed', 'reopened'].includes(props.event.payload.action)">
        {{ props.event.payload.action }}
      </span>
      <span v-else-if="['assigned', 'unassigned'].includes(props.event.payload.action)">changed assignees for</span>
      <span v-else-if="['labeled', 'unlabeled'].includes(props.event.payload.action)">edited labels for</span>
    </div>

    issue
    <EventHoverItem :href="props.event.payload.issue.html_url" :value="'#' + props.event.payload.issue.number">
      {{ props.event.payload.issue.title }}
    </EventHoverItem>

    on
    <EventLink :href="props.event.repo.name" />

    <EventBlame>{{ props.event.payload.issue.title }}</EventBlame>
  </div>
</template>
