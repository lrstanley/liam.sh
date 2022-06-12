<template>
  <div>
    <span v-if="action == 'created'"> added </span>
    <span v-else-if="action == 'edited'"> updated </span>
    <span v-if="action == 'deleted'"> removed </span>

    a
    <EventLink
      :href="comment.html_url"
      class="text-cyan-400 hover:text-cyan-500"
      value="review comment"
    />
    to pr

    <EventHoverItem :href="pr.html_url" :value="'#' + pr.number" class="!text-purple-500">
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
const comment = ref(props.event.payload.comment)
const pr = ref(props.event.payload.pull_request)
</script>
