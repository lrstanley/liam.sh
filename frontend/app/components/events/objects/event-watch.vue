<script setup lang="ts">
const props = defineProps<{
  event: GithubEvent
}>()
</script>

<template>
  <div class="text-amber-300">
    {{ props.event.payload.action == "started" ? "starred" : "unstarred" }}
    <EventLink :href="props.event.repo.name as string" />
  </div>
</template>
