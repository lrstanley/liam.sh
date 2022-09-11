<template>
  <div>
    <span v-if="action == 'created'"> added </span>
    <span v-else-if="action == 'edited'"> updated </span>
    <span v-if="action == 'deleted'"> removed </span>

    a
    <EventLink :href="comment.html_url" class="text-cyan-400 hover:text-cyan-500" value="comment" />
    to {{ issue.pull_request ? "pr" : "issue" }}

    <EventHoverItem :href="issue.html_url" :value="'#' + issue.number">
      {{ issue.title }}
    </EventHoverItem>

    on <EventLink :href="repo.name" />

    <EventBlame>{{ issue.title }}</EventBlame>
  </div>
</template>

<script setup lang="ts">
import type { GithubEvent } from "@/lib/api"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action)
const comment = ref<Record<string, any>>(props.event.payload.comment)
const issue = ref<Record<string, any>>(props.event.payload.issue)
</script>
