<script setup lang="ts">
import type { SchemaOutdatedRepositoryRelease } from '#open-fetch-schemas/api'
import type { TableColumn } from "#ui/types"

definePageMeta({
  title: "Repository Releases",
  layout: "admin",
})

const { data, error, status } = await useApi('/github-releases/outdated')
if (error.value) throw error.value

const now = new Date()

const columns: TableColumn<SchemaOutdatedRepositoryRelease>[] = [
  { id: "repository", header: "Repository", accessorKey: "repository" },
  { id: "lastReleaseVersion", header: "Last Release Version", accessorKey: "lastRelease" },
  { id: "lastReleaseDate", header: "Last Release", accessorKey: "lastRelease" },
]
</script>

<template>
  <UCard variant="subtle">
    <UTable v-if="data" :data="data" :columns="columns" :loading="status === 'pending'" loading-color="primary"
      loading-animation="carousel" class="shrink-0">
      <template #repository-cell="{ row }">
        <NuxtLink class="flex flex-row" :href="'https://github.com/' + row.original.repository.full_name + '/releases'"
          target="_blank">
          <img :src="row.original.repository.owner.avatar_url" class="w-6 h-6 rounded-full" />
          <span class="ml-2">{{ row.original.repository.full_name }}</span>
        </NuxtLink>
      </template>

      <template #lastReleaseVersion-cell="{ row }">
        <NuxtLink :href="row.original.release.html_url" target="_blank">
          <span>{{ row.original.release.name || row.original.release.tag_name }}</span>
        </NuxtLink>
      </template>

      <template #lastReleaseDate-cell="{ row }">
        <span :class="{
          'text-error':
            new Date(row.original.release.created_at).getTime() < now.getTime() - 90 * 86400000,
        }">
          {{ useTimeAgo(row.original.release.created_at).value }}
        </span>
      </template>
    </UTable>
  </UCard>
</template>
