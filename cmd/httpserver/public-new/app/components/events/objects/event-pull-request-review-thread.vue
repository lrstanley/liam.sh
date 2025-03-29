<script setup lang="ts">
const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const thread = ref<Record<string, any>>(props.event.payload.thread as any)
const pr = ref<Record<string, any>>(props.event.payload.pull_request as any)
</script>

<template>
  <div>
    marked a
    <EventLink
      :href="thread.html_url"
      class="text-(--ui-color-info-400) hover:text-(--ui-color-info-500)"
      value="thread"
    />

    as

    <span
      :class="{ 'text-(--ui-success)': action == 'resolved', 'text-red-400': action == 'unresolved' }"
    >
      {{ action }}
    </span>

    for pr

    <EventHoverItem :href="pr.html_url" :value="'#' + pr.number">
      {{ pr.title }}
    </EventHoverItem>

    on
    <EventLink :href="repo.name as string" />

    <EventBlame>{{ pr.title }}</EventBlame>
  </div>
</template>
