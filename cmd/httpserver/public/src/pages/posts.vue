<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router"
import { listPosts } from "@/lib/http/services.gen"
import type { PostSortableFields } from "@/lib/http/types.gen"

definePage({
  meta: {
    title: "Posts",
    layout: "default",
  },
})

const sortOptions = {
  published_at: "date",
  title: "title",
  view_count: "popularity",
}

const pagination = usePagination<PostSortableFields>({ sort: "published_at" })
const labels = useRouteQuery<Array<string>>("label", [])
const search = useRouteQuery<string>("q", "")
const debounceSearch = refDebounced<string>(search, 300)

resetPagination(pagination, [labels, debounceSearch])

const {
  data: posts,
  error,
  isPending,
} = useQuery({
  queryKey: ["posts", "admin", pagination, labels, debounceSearch],
  queryFn: () =>
    unwrapErrors(
      listPosts({
        query: {
          page: pagination.page.value,
          per_page: pagination.perPage.value,
          sort: pagination.sort.value,
          order: pagination.order.value,
          "contentHTML.ihas": debounceSearch.value,
          "label.name.in": labels.value.length > 0 ? labels.value : null,
          "public.eq": true,
        },
      })
    ),
  placeholderData: keepPreviousData,
})
</script>

<template>
  <div class="mt-8 grid-sidebar">
    <div>
      <div class="flex flex-auto gap-2 mt-1 mb-8">
        <n-input
          v-model:value="search"
          :loading="isPending"
          type="text"
          clearable
          placeholder="Search for a post"
        >
          <template #prefix>
            <n-icon>
              <i-mdi-search />
            </n-icon>
          </template>
        </n-input>
      </div>

      <CoreError v-if="error" :error="error" />
      <CoreObjectRender
        v-else
        :value="posts?.content"
        type="post"
        linkable
        show-empty
        divider
        :loading="isPending"
      />

      <div class="flex">
        <n-pagination
          v-model:page="pagination.page.value"
          v-model:page-size="pagination.perPage.value"
          :item-count="posts?.total_count"
          :page-sizes="pagination.sizes"
          show-size-picker
          class="mt-2 ml-auto"
        />
      </div>
    </div>

    <div>
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort posts</div>
        <CoreSorter
          v-if="pagination"
          v-model="pagination.sort.value"
          v-model:order="pagination.order.value"
          :sort-options="sortOptions"
          class="pb-4"
        />

        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect v-model="labels" :where="{ hasPosts: true }" />
      </div>
    </div>
  </div>
</template>
