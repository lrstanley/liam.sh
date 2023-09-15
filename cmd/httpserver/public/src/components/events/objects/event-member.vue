<script setup lang="ts">
import type { GithubEvent } from "@/lib/api"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action)
const member = ref<Record<string, any>>(props.event.payload.member)
</script>

<template>
  <div>
    <span v-if="action == 'edited'">updated collaborator permissions on</span>
    <span v-else>
      {{ action }}

      <EventLink :href="member.html_url" :title="member.login" />

      as collaborator on
    </span>

    <EventLink :href="repo.name" />
  </div>
</template>
