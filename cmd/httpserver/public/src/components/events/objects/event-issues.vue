<template>
  <div>
    <span class="text-cyan-400">
      <span v-if="['opened', 'edited', 'closed', 'reopened'].includes(action)">
        {{ action }}
      </span>
      <span v-else-if="['assigned', 'unassigned'].includes(action)"> changed assignees on </span>
      <span v-else-if="['labeled', 'unlabeled'].includes(action)"> edited labels on </span>
    </span>

    issue
    <EventHoverItem :href="issue.html_url" :value="'#' + issue.number" class="!text-purple-500">
      {{ issue.title }}
    </EventHoverItem>

    on <EventLink :href="repo.name" />

    <EventBlame>{{ issue.title }}</EventBlame>
  </div>
</template>

<script setup>
const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
})

const repo = ref(props.event.repo)
const action = ref(props.event.payload.action)
const issue = ref(props.event.payload.issue)
</script>
