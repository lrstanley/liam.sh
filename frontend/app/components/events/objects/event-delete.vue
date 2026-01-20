<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["delete-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <span class="text-red-400">deleted</span>
    <template v-if="props.event.payload.ref">
      <EventHoverItem :value="props.event.payload.ref" class="truncate text-default block">
        <template #icon>
          <UIcon name="mdi:source-branch" v-if="props.event.payload.ref_type == 'branch'" />
          <UIcon name="mdi:tag" v-else />
        </template>

        {{ props.event.payload.ref }}
      </EventHoverItem>
      from
    </template>
    <EventLink :href="props.event.repo.name" />
  </div>
</template>
