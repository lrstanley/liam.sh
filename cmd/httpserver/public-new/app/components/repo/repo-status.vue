<script setup lang="ts">
const { value: repo } = defineProps<{
  value: GithubRepositoryRead
}>()

const gh = useGithubUser()
</script>

<template>
  <div class="inline-flex flex-row gap-1">
    <UBadge v-if="repo.owner.login != gh?.login" color="info" variant="outline">maintainer</UBadge>
    <UBadge v-if="repo.fork" color="error" variant="outline" icon="mdi:source-fork">fork</UBadge>
    <UBadge v-if="repo.archived" color="info" variant="outline" icon="mdi:archive-outline">
      archived
    </UBadge>
    <UBadge v-if="repo.license" variant="outline" :title="repo.license.name" icon="mdi:scale-balance">
      {{ repo.license.key }}
    </UBadge>
  </div class="inline-flex flex-row gap-1">
</template>
