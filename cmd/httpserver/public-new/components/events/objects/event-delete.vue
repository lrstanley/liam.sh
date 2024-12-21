<script setup lang="ts">
import type { GithubEvent } from "@/utils/http/types.gen"

const props = defineProps<{
  event: GithubEvent
}>()
</script>

<template>
  <div>
    <span class="text-red-400">deleted</span>
    <template v-if="props.event.payload.ref">
      <EventHoverItem :value="props.event.payload.ref as string" class="truncate text-zinc-200">
        <template #icon>
          <Icon name="mdi:source-branch" v-if="props.event.payload.ref_type == 'branch'" />
          <Icon name="mdi:tag" v-else />
        </template>

        {{ props.event.payload.ref }}
      </EventHoverItem>
      from
    </template>
    <EventLink :href="props.event.repo.name as string" />
  </div>
</template>
