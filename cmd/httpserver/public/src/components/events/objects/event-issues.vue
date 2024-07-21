<script setup lang="ts">
import type { GithubEvent } from "@/lib/http/types.gen"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const issue = ref<Record<string, any>>(props.event.payload.issue)
</script>

<template>
  <div>
    <span class="text-cyan-400">
      <span v-if="['opened', 'edited', 'closed', 'reopened'].includes(action)">
        {{ action }}
      </span>
      <span v-else-if="['assigned', 'unassigned'].includes(action)">changed assignees on</span>
      <span v-else-if="['labeled', 'unlabeled'].includes(action)">edited labels on</span>
    </span>

    issue
    <EventHoverItem :href="issue.html_url" :value="'#' + issue.number">
      {{ issue.title }}
    </EventHoverItem>

    on
    <EventLink :href="repo.name" />

    <EventBlame>{{ issue.title }}</EventBlame>
  </div>
</template>
