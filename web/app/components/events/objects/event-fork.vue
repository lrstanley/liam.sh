<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["fork-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <template v-if="$props.event.repo_id === props.event.payload.forkee?.id">
      <span class="text-secondary-400">created fork</span>
    </template>
    <template v-else>
      <span class="text-secondary-400">forked</span>
      <EventLink :href="props.event.repo.name" />
      to
    </template>
    <EventLink :href="props.event.payload.forkee?.full_name ?? ''" />

    <EventBlame>{{ props.event.payload.forkee?.description ?? '' }}</EventBlame>
  </div>
</template>
