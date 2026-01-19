<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router"
import type { SchemaPostSortableFields } from '#open-fetch-schemas/api'

definePageMeta({
  title: "Posts",
  layout: "default",
})

const labels = useRouteQuery<Array<string>>("label", [])
const search = useRouteQuery<string>("q", "")
const debounceSearch = refDebounced<string>(search, 300)
const pagination = usePagination<SchemaPostSortableFields>({
  sort: "published_at",
  resetChanged: [labels, debounceSearch],
})

const {
  data: posts,
  error,
  status,
} = await useApi('/posts', {
  query: computed(() => ({
    page: pagination.page.value,
    per_page: pagination.perPage.value,
    sort: pagination.sort.value,
    order: pagination.order.value,
    "search.ihas": debounceSearch.value,
    "label.name.in": labels.value.length > 0 ? labels.value : undefined,
    "public.eq": true,
  })),
})
if (error.value) throw error.value
</script>

<template>
  <ContainerStickySidebar class="mt-8">
    <div class="flex flex-auto gap-2 mt-1 mb-4">
      <UInput v-model="search" :loading="status == 'pending'" type="search" placeholder="Search for a post"
        icon="mdi:search" class="w-full">
        <template v-if="search.length" #trailing>
          <UButton color="neutral" variant="link" size="sm" icon="lucide:circle-x" aria-label="Clear input"
            @click="search = ''" />
        </template>
      </UInput>

      <CorePagination class="ml-auto" :loading="status === 'pending'" :resp="posts" :pagination="pagination" />
    </div>

    <CoreObjectRender :value="posts?.content" type="post" show-empty divider :loading="status === 'pending'" />

    <template #sidebar>
      <div>
        <div class="text-primary">Sort posts</div>
        <CoreSorter v-if="pagination" v-model="pagination.sort.value" v-model:order="pagination.order.value"
          :sort-options="{
            published_at: { displayName: 'date', defaultOrder: 'desc' },
            title: { displayName: 'title', defaultOrder: 'asc' },
            view_count: { displayName: 'popularity', defaultOrder: 'desc' },
          }" />
      </div>

      <div>
        <div class="text-primary">Filter by label</div>
        <LabelSelect v-model="labels" />
      </div>
    </template>
  </ContainerStickySidebar>
</template>
