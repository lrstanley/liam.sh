<script setup lang="ts">
import type { GithubEvent } from "@/lib/api"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action)
const pr = ref<Record<string, any>>(props.event.payload.pull_request)
</script>

<template>
  <div>
    <span v-if="action == 'closed' && pr.merged" class="text-purple-400">merged</span>
    <span v-else-if="['opened', 'edited', 'closed', 'reopened'].includes(action)">
      {{ action }}
    </span>
    <span v-else-if="action == 'synchronize'">updated</span>
    <span v-else-if="['assigned', 'unassigned'].includes(action)">changed assignees on</span>
    <span v-else-if="['review_requested', 'review_request_removed'].includes(action)">
      requested review on
    </span>
    <span v-else-if="['labeled', 'unlabeled'].includes(action)">edited labels on</span>

    pr

    <EventHoverItem :href="pr.html_url" :value="'#' + pr.number">
      {{ pr.title }}
    </EventHoverItem>

    <template v-if="['opened', 'edited', 'closed'].includes(action)">
      via

      <EventHoverItem :value="pr.head.ref" class="truncate text-zinc-200">
        <template #icon>
          <i-mdi-source-pull />
        </template>

        {{ pr.head.ref }}
      </EventHoverItem>
    </template>

    on
    <EventLink :href="repo.name" />

    <EventBlame>{{ pr.title }}</EventBlame>
  </div>
</template>
