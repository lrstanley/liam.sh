<route lang="yaml">
meta:
  title: Posts
  layout: default
</route>

<template>
  <div class="mt-8 grid-sidebar">
    <div>
      <div class="flex flex-auto gap-2 mt-1 mb-8">
        <n-input
          v-model:value="search"
          :loading="fetching"
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

        <CorePagination v-model="cursor" :page-info="data?.posts?.pageInfo" />
      </div>

      <CoreObjectRender v-if="data?.posts" :value="data.posts" linkable show-empty divider />
    </div>

    <div>
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort posts</div>
        <CoreSorter :sorter="sorter" class="pb-4" />

        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect v-model="labels" :where="{ hasPosts: true }" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router"
import { useGetPostsQuery } from "@/lib/api"
import { usePagination, resetCursor, useSorter } from "@/lib/util"

const cursor = useRouteQuery<string>("cur", null)
const labels = useRouteQuery<Array<string>>("label", [])
const search = useRouteQuery<string>("q", "")
const filterSearch = refDebounced<string>(search, 300)
const direction = useRouteQuery<string>("dir", "desc")
const field = useRouteQuery<string>("sort", "date")
const sorter = useSorter(
  {
    date: "date",
    title: "title",
    view_count: "popularity",
  },
  direction,
  field
)

resetCursor(cursor, [labels, search, direction, field])

const { data, error, fetching } = await useGetPostsQuery({
  variables: {
    ...usePagination(cursor, 10),
    ...sorter.filter,
    where: {
      or: [{ titleContainsFold: filterSearch }, { summaryContainsFold: filterSearch }],
      hasLabelsWith: computed(() => (labels.value.length ? { nameIn: labels.value } : null)),
    },
  },
})

watch(error, () => {
  if (error.value) throw error.value
})
</script>
