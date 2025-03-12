<script setup lang="ts">
const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const release = ref<Record<string, any>>(props.event.payload.release as any)
</script>

<template>
  <div>
    <span class="text-lime-500">{{ action }}</span>

    release from tag
    <EventHoverItem :href="release.html_url" :value="release.tag_name" class="truncate">
      {{ release.name }}
    </EventHoverItem>
    on
    <EventLink :href="repo.name as string" />

    <EventBlame>{{ release.name }}</EventBlame>
  </div>
</template>
