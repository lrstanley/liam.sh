<route lang="yaml">
meta:
  title: Home
  layout: terminal
</route>

<template>
  <ContainerIde>
    <EventsRender
      class="relative w-full h-full overflow-x-hidden grow basis-0"
      @event-count="(e) => (eventCount = e)"
    />

    <template #footer>
      <CodingStats placement="top-start" />

      <span class="ml-auto" />

      <span v-if="eventCount > 0" class="bar-item misc">ln:{{ eventCount }}</span>
    </template>
  </ContainerIde>
</template>

<script setup lang="ts">
const eventCount = ref<number>()
</script>
