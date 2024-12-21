<script setup lang="ts">
import type { GithubEvent } from "@/utils/http/types.gen"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const action = ref<string>(props.event.payload.action as string)
const member = ref<Record<string, any>>(props.event.payload.member as any)
</script>

<template>
  <div>
    <span v-if="action == 'edited'">updated collaborator permissions on</span>
    <span v-else>
      {{ action }}

      <EventLink :href="member.html_url" :title="member.login" />

      as collaborator on
    </span>

    <EventLink :href="repo.name as string" />
  </div>
</template>
