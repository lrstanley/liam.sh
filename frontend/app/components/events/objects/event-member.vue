<script setup lang="ts">
import type { SchemaGithubEvent } from "#open-fetch-schemas/api";
import type { components as Github } from "./events"

const props = defineProps<{
  event: Exclude<SchemaGithubEvent, "payload" | "repo"> & { payload: Github["schemas"]["member-event"], repo: Github["schemas"]["event"]["repo"] }
}>()
</script>

<template>
  <div>
    <span v-if="props.event.payload.action == 'edited'">updated collaborator permissions on</span>
    <span v-else>
      {{ props.event.payload.action }}

      <EventLink :href="props.event.payload.member.html_url" :title="props.event.payload.member.login" />

      as collaborator on
    </span>

    <EventLink :href="props.event.repo.name" />
  </div>
</template>
