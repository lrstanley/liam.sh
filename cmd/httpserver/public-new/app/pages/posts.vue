<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router"

useHead({ title: "Posts" })
definePageMeta({
  layout: "default",
})

const pagination = usePagination<PostSortableFields>({ sort: "published_at" })
const labels = useRouteQuery<Array<string>>("label", [])
const search = useRouteQuery<string>("q", "")
const debounceSearch = refDebounced<string>(search, 300)

// TODO: make this an argument to usePagination?
resetPagination(pagination, [labels, debounceSearch])

const {
  data: posts,
  error,
  status,
} = listPosts({
  composable: "useFetch",
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
      <UInput
        v-model="search"
        :loading="status == 'pending'"
        type="search"
        placeholder="Search for a post"
        icon="mdi:search"
        class="w-full"
      >
        <template v-if="search.length" #trailing>
          <UButton
            color="neutral"
            variant="link"
            size="sm"
            icon="lucide:circle-x"
            aria-label="Clear input"
            @click="search = ''"
          />
        </template>
      </UInput>

      <UPagination
        v-show="posts?.last_page === undefined || posts?.last_page > 1"
        class="ml-auto"
        v-model:page="pagination.page.value"
        :items-per-page="pagination.perPage.value"
        :sibling-count="1"
        active-color="primary"
        active-variant="subtle"
        :show-controls="false"
        show-edges
        :total="posts?.total_count"
        :disabled="status === 'pending'"
      />
    </div>

    <CoreObjectRender
      :value="posts?.content"
      type="post"
      linkable
      show-empty
      divider
      :loading="status === 'pending'"
    />

    <template #sidebar>
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort posts</div>
        <CoreSorter
          v-if="pagination"
          v-model="pagination.sort.value"
          v-model:order="pagination.order.value"
          :sort-options="{
            published_at: { displayName: 'date', defaultOrder: 'desc' },
            title: { displayName: 'title', defaultOrder: 'asc' },
            view_count: { displayName: 'popularity', defaultOrder: 'desc' },
          }"
          class="pb-4"
        />

        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect v-model="labels" />
      </div>
    </template>
  </ContainerStickySidebar>
</template>
