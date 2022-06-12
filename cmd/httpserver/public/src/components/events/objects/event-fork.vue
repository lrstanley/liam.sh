<template>
  <div>
    <span class="text-purple-400">forked</span>
    <EventLink :href="repo.name" />
    to
    <EventLink :href="fork.full_name" />

    <EventBlame>{{ fork.description }}</EventBlame>
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
const fork = ref(props.event.payload.forkee)
</script>
