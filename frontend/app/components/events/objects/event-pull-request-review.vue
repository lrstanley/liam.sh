<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["pull-request-review-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <div v-if="props.event.payload.review.state == 'approved'" class="text-success">approved</div>
    <div v-else-if="props.event.payload.review.state == 'changes_requested'">
      <span class="text-red-400">requested changes</span>
      on
    </div>
    <div v-else-if="props.event.payload.review.state == 'commented'">
      added a
      <span class="text-info-400">review comment</span>
      to
    </div>
    <div v-else>
      <span v-if="props.event.payload.action == 'created'">added</span>
      <span v-else-if="props.event.payload.action == 'edited'">updated</span>
      <span v-if="props.event.payload.action == 'deleted'">removed</span>

      a
      <EventLink :href="props.event.payload.review.html_url ?? $props.event.payload.pull_request.url"
        class="text-info-400 hover:text-info-500" :value="'review ' + props.event.payload.review.state" />
      to
    </div>

    pr

    <EventLink :href="props.event.payload.pull_request.url" :value="'#' + props.event.payload.pull_request.number" />

    on
    <EventLink :href="props.event.repo.name" />

    <EventBlame>{{ props.event.payload.pull_request.head.ref }}</EventBlame>
  </div>
</template>
