<template>
  <div>
    <span v-if="review.state == 'approved'" class="text-lime-500"> approved </span>
    <span v-else-if="review.state == 'changes_requested'">
      <span class="text-red-400">requested changes</span> on
    </span>
    <span v-else-if="review.state == 'commented'">
      added a
      <span class="text-cyan-400">review comment</span>
      to
    </span>
    <span v-else>
      <span v-if="action == 'created'"> added </span>
      <span v-else-if="action == 'edited'"> updated </span>
      <span v-if="action == 'deleted'"> removed </span>

      a
      <EventLink
        :href="review.html_url"
        class="text-cyan-400 hover:text-cyan-500"
        :value="'review ' + review.state"
      />
      to
    </span>

    pr
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
const review = ref(props.event.payload.review)
const pr = ref(props.event.payload.pull_request)
</script>
