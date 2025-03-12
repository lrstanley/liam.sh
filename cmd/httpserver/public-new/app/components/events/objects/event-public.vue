<script setup lang="ts">
const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
</script>

<template>
  <div>
    made
    <EventLink :href="repo.name as string" />
    public
  </div>
</template>
