<template>
  <div>
    <span class="text-lime-500">created</span>
    <template v-if="props.event.payload.ref">
      <EventHoverItem :value="props.event.payload.ref" class="text-zinc-200 truncate">
        <template #icon>
          <i-mdi-source-branch v-if="props.event.payload.ref_type == 'branch'" />
          <i-mdi-tag v-else />
        </template>

        {{ props.event.payload.ref }}
      </EventHoverItem>
      on
    </template>
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
