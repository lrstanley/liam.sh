<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["pull-request-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <div class="text-fuchsia-400">
      <span v-if="['opened', 'edited', 'closed', 'reopened'].includes(props.event.payload.action)">
        {{ props.event.payload.action }}
      </span>
      <span v-else-if="props.event.payload.action == 'synchronize'">updated</span>
      <span v-else-if="['assigned', 'unassigned'].includes(props.event.payload.action)">changed assignees for</span>
      <span v-else-if="['review_requested', 'review_request_removed'].includes(props.event.payload.action)">
        requested review for
      </span>
      <span v-else-if="['labeled', 'unlabeled'].includes(props.event.payload.action)">edited labels for</span>
    </div>

    pr

    <EventLink :href="props.event.payload.pull_request.url" :value="'#' + props.event.payload.pull_request.number" />

    <template v-if="['opened', 'edited', 'closed'].includes(props.event.payload.action)">
      via

      <EventHoverItem :value="props.event.payload.pull_request.head.ref" class="truncate text-default block">
        <template #icon>
          <UIcon name="mdi:source-pull" />
        </template>

        {{ props.event.payload.pull_request.head.ref }}
      </EventHoverItem>
    </template>

    on
    <EventLink :href="props.event.repo.name" />

    <EventBlame>{{ props.event.payload.pull_request.head.ref }}</EventBlame>
  </div>
</template>
