<script setup lang="ts">
import type { GithubEvent } from "@/utils/http/types.gen"

const props = defineProps<{
  event: GithubEvent
}>()

const repo = ref(props.event.repo)
const commits = ref<Record<string, any>[]>(props.event.payload.commits as Record<string, any>[])
</script>

<template>
  <div>
    <span style="color: #d086ff">pushed</span>

    <span v-if="commits.length > 1">
      <EventHoverItem
        :href="repo.name + '/compare/' + commits[0].sha + '...' + props.event.payload.head"
        :value="commits.length + ' commits'"
        style="max-width: 350px"
      >
        <p v-for="commit in commits" :key="commit.sha" class="truncate" :title="commit.message">
          <UIcon name="mdi:source-commit" />
          {{ commit.sha.slice(0, 7) }}:
          {{ commit.message.split("\n")[0] }}
        </p>
      </EventHoverItem>
    </span>
    <span v-else>
      <EventHoverItem
        :href="repo.name + '/commit/' + commits[0].sha"
        :value="commits[0].sha.slice(0, 7)"
      >
        <UIcon name="mdi:source-commit" />
        {{ commits[0].sha.slice(0, 7) }}:
        {{ commits[0].message.split("\n")[0] }}
      </EventHoverItem>
    </span>
    to
    <EventLink :href="repo.name as string" />

    <EventBlame>
      {{ commits[0].message.split("\n")[0] }}
    </EventBlame>
  </div>
</template>
