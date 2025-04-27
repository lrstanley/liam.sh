<script setup lang="ts">
const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const comment = ref<Record<string, any>>(props.event.payload.comment as any)
const pr = ref<Record<string, any>>(props.event.payload.pull_request as any)
</script>

<template>
  <div>
    <span v-if="action == 'created'">added</span>
    <span v-else-if="action == 'edited'">updated</span>
    <span v-if="action == 'deleted'">removed</span>

    a
    <EventLink
      :href="comment.html_url"
      class="text-(--ui-color-info-400) hover:text-(--ui-color-info-500)"
      value="review comment"
    />
    to pr

    <EventHoverItem :href="pr.html_url" :value="'#' + pr.number">
      {{ pr.title }}
    </EventHoverItem>

    on
    <EventLink :href="repo.name as string" />

    <EventBlame>{{ pr.title }}</EventBlame>
  </div>
</template>
