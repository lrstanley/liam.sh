<template>
  <LayoutDefault :error="error">
    <div class="grid gap-5 md:gap-10 lg:gap-[6rem] mt-8">
      <div class="order-last md:order-first">
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
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort posts</div>
        <CoreSorter :sorter="sorter" class="pb-4" />

        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect v-model="labels" :where="{ hasPosts: true }" />
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup>
import { useRouteQuery } from "@vueuse/router"
import { useGetPostsQuery } from "@/lib/api"
import { usePagination, resetCursor } from "@/lib/pagination"
import { useSorter } from "@/lib/sorter"

const cursor = useRouteQuery("cur", null)
const labels = useRouteQuery("label", [])
const search = useRouteQuery("q", "")
const filterSearch = refDebounced(search, 300)
const direction = useRouteQuery("dir", "desc")
const field = useRouteQuery("sort", "date")
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

const { data, error, fetching } = useGetPostsQuery({
  variables: {
    ...usePagination(cursor, 10),
    ...sorter.filter,
    where: {
      or: [{ titleContainsFold: filterSearch }, { summaryContainsFold: filterSearch }],
      hasLabelsWith: computed(() => (labels.value.length ? { nameIn: labels.value } : null)),
    },
  },
})
</script>

<style scoped>
.grid {
  @apply grid-cols-1 md:grid-cols-[1fr,240px];
}
</style>
