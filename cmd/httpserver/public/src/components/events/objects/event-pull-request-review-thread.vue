<template>
  <div>
    marked a
    <EventLink :href="thread.html_url" class="text-cyan-400 hover:text-cyan-500" value="thread" />

    as

    <span :class="{ 'text-lime-500': action == 'resolved', 'text-red-400': action == 'unresolved' }">
      {{ action }}
    </span>

    for pr

    <EventHoverItem :href="pr.html_url" :value="'#' + pr.number">
      {{ pr.title }}
    </EventHoverItem>

    on <EventLink :href="repo.name" />

    <EventBlame>{{ pr.title }}</EventBlame>
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
const thread = ref(props.event.payload.thread)
const pr = ref(props.event.payload.pull_request)
</script>
