<script setup lang="ts">
useHead({ title: "Home" })
definePageMeta({
  layout: "terminal",
})

const eventCount = ref<number>()
</script>

<template>
  <ContainerIde>
    <EventsRender
      class="relative overflow-x-hidden overflow-y-scroll size-full grow basis-0"
      @event-count="(e) => (eventCount = e)"
    />

    <template #footer>
      <CodingStats placement="top-start" />

      <div class="ml-auto" />

      <div v-show="eventCount && eventCount > 0" class="bar-item misc">ln:{{ eventCount }}</div>
    </template>
  </ContainerIde>
</template>
