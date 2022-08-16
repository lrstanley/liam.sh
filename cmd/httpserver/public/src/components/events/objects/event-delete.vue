<template>
  <div>
    <span class="text-red-400">deleted</span>
    <template v-if="props.event.payload.ref">
      <EventHoverItem :value="props.event.payload.ref" class="truncate text-zinc-200">
        <template #icon>
          <i-mdi-source-branch v-if="props.event.payload.ref_type == 'branch'" />
          <i-mdi-tag v-else />
        </template>

        {{ props.event.payload.ref }}
      </EventHoverItem>
      from
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
