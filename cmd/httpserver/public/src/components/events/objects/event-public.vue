<template>
  <div>made <EventLink :href="repo.name" /> public</div>
</template>

<script setup>
const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
})

const repo = ref(props.event.repo)
</script>
