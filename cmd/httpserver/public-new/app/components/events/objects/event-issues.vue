<script setup lang="ts">
import type { GithubEvent } from "@/utils/http/types.gen"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const issue = ref<Record<string, any>>(props.event.payload.issue as any)
</script>

<template>
  <div>
    <div class="text-cyan-400">
      <span v-if="['opened', 'edited', 'closed', 'reopened'].includes(action)">
        {{ action }}
      </span>
      <span v-else-if="['assigned', 'unassigned'].includes(action)">changed assignees on</span>
      <span v-else-if="['labeled', 'unlabeled'].includes(action)">edited labels on</span>
    </div>

    issue
    <EventHoverItem :href="issue.html_url" :value="'#' + issue.number">
      {{ issue.title }}
    </EventHoverItem>

    on
    <EventLink :href="repo.name as string" />

    <EventBlame>{{ issue.title }}</EventBlame>
  </div>
</template>
