<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["pull-request-review-comment-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <span v-if="props.event.payload.action == 'created'">added</span>
    <span v-else-if="props.event.payload.action == 'edited'">updated</span>
    <span v-if="props.event.payload.action == 'deleted'">removed</span>

    a
    <EventLink :href="props.event.payload.comment.html_url" class="text-info-400 hover:text-info-500"
      value="review comment" />
    to pr
    <EventLink :href="props.event.payload.pull_request.url" :value="'#' + props.event.payload.pull_request.number" />
    on
    <EventLink :href="props.event.repo.name" />

    <EventBlame>
      {{ props.event.payload.pull_request.head.ref }}
    </EventBlame>
  </div>
</template>
