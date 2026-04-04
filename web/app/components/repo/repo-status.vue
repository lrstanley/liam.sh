<script setup lang="ts">
import type { SchemaGithubRepositoryRead } from '#open-fetch-schemas/api'

const { value: repo } = defineProps<{
  value: SchemaGithubRepositoryRead
}>()

const gh = useGithubUser()
</script>

<template>
  <div class="inline-flex flex-row gap-1">
    <UBadge v-if="repo.owner.login != gh?.login" color="info" variant="outline" class="max-sm:hidden">
      maintainer
    </UBadge>
    <UBadge v-if="repo.fork" color="error" variant="outline" icon="mdi:source-fork">fork</UBadge>
    <UBadge v-if="repo.archived" color="warning" variant="outline" icon="mdi:archive-outline">
      archived
    </UBadge>
    <UBadge v-if="repo.license" variant="outline" :title="repo.license.name" icon="mdi:scale-balance">
      {{ repo.license.key }}
    </UBadge>
  </div>
</template>
