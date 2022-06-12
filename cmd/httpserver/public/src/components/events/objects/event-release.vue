<template>
  <div>
    <span class="text-lime-500">{{ action }}</span>

    release from tag
    <EventHoverItem :href="release.html_url" :value="release.tag_name" class="truncate">
      {{ release.name }}
    </EventHoverItem>
    on
    <EventLink :href="repo.name" />

    <EventBlame>{{ release.name }}</EventBlame>
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
const release = ref(props.event.payload.release)
</script>
