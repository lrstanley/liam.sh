<script setup lang="ts">
definePageMeta({
  layout: "terminal",
})

const eventCount = ref<number>()
</script>

<template>
  <ContainerIde>
    <EventsRender
      class="relative overflow-x-hidden size-full grow basis-0"
      @event-count="(e) => (eventCount = e)"
    />
    <!-- <div
      style="height: 1300px"
      class="relative overflow-x-hidden overflow-y-scroll size-full grow basis-0"
    >
      TODO
    </div> -->

    <template #footer>
      <CodingStats placement="top-start" />

      <span class="ml-auto" />

      <span v-if="eventCount && eventCount > 0" class="bar-item misc">ln:{{ eventCount }}</span>
    </template>
  </ContainerIde>
</template>
