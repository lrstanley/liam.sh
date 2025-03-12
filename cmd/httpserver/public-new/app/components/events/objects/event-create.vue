<script setup lang="ts">
const props = defineProps<{
  event: GithubEvent
}>()
</script>

<template>
  <div>
    <span class="text-lime-500">created</span>
    <template v-if="props.event.payload.ref">
      <EventHoverItem :value="props.event.payload.ref as string" class="truncate text-zinc-200">
        <template #icon>
          <UIcon name="mdi:source-branch" v-if="props.event.payload.ref_type == 'branch'" />
          <UIcon name="mdi:tag" v-else />
        </template>

        {{ props.event.payload.ref }}
      </EventHoverItem>
      on
    </template>
    <EventLink :href="props.event.repo.name as string" />
  </div>
</template>
