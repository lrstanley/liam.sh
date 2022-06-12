<template>
  <div class="text-amber-300">
    {{ props.event.payload.action == "started" ? "starred" : "unstarred" }}
    <EventLink :href="props.event.repo.name" />
  </div>
</template>

<script setup>
const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
})
</script>
