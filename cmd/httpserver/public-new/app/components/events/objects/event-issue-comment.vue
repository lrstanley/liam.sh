<script setup lang="ts">
import type { GithubEvent } from "@/utils/http/types.gen"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const comment = ref<Record<string, any>>(props.event.payload.comment as any)
const issue = ref<Record<string, any>>(props.event.payload.issue as any)
</script>

<template>
  <div>
    <span v-if="action == 'created'">added</span>
    <span v-else-if="action == 'edited'">updated</span>
    <span v-if="action == 'deleted'">removed</span>

    a
    <EventLink :href="comment.html_url" class="text-cyan-400 hover:text-cyan-500" value="comment" />
    to {{ issue.pull_request ? "pr" : "issue" }}

    <EventHoverItem :href="issue.html_url" :value="'#' + issue.number">
      {{ issue.title }}
    </EventHoverItem>

    on
    <EventLink :href="repo.name as string" />

    <EventBlame>{{ issue.title }}</EventBlame>
  </div>
</template>
