<script setup lang="ts">
const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const review = ref<Record<string, any>>(props.event.payload.review as any)
const pr = ref<Record<string, any>>(props.event.payload.pull_request as any)
</script>

<template>
  <div>
    <div v-if="review.state == 'approved'" class="text-lime-500">approved</div>
    <div v-else-if="review.state == 'changes_requested'">
      <span class="text-red-400">requested changes</span>
      on
    </div>
    <div v-else-if="review.state == 'commented'">
      added a
      <span class="text-cyan-400">review comment</span>
      to
    </div>
    <div v-else>
      <span v-if="action == 'created'">added</span>
      <span v-else-if="action == 'edited'">updated</span>
      <span v-if="action == 'deleted'">removed</span>

      a
      <EventLink
        :href="review.html_url"
        class="text-cyan-400 hover:text-cyan-500"
        :value="'review ' + review.state"
      />
      to
    </div>

    pr
    <EventHoverItem :href="pr.html_url" :value="'#' + pr.number">
      {{ pr.title }}
    </EventHoverItem>

    on
    <EventLink :href="repo.name as string" />

    <EventBlame>{{ pr.title }}</EventBlame>
  </div>
</template>
