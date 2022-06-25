<template>
  <LayoutDefault :error="error">
    <div class="grid gap-5 md:gap-10 lg:gap-25 mt-8">
      <div class="order-last md:order-first">
        <div class="mt-1 mb-8 flex flex-auto gap-2">
          <n-input
            v-model:value="search"
            :loading="fetching"
            type="text"
            clearable
            placeholder="Search for a repo"
          >
            <template #prefix>
              <n-icon>
                <i-mdi-search />
              </n-icon>
            </template>
          </n-input>

          <CorePagination v-model="cursor" :page-info="data?.githubrepositories?.pageInfo" />
        </div>

        <CoreObjectRender
          v-if="data?.githubrepositories"
          :value="data.githubrepositories"
          linkable
          show-empty
          divider
        />
      </div>
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort repos</div>
        <CoreSorter :sorter="sorter" class="pb-4" />

        <div class="text-emerald-500">Filter attributes</div>
        <n-space size="small" class="pb-4" inline>
          <n-checkbox :checked="archived" @update:checked="(v) => (archived = !!v)">
            include archived
          </n-checkbox>
          <n-checkbox :checked="forks == 'true'" @update:checked="(v) => (forks = !!v)">
            include forks
          </n-checkbox>
        </n-space>

        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect
          v-model="labels"
          :where="{
            hasGithubRepositoriesWith: {
              fork: where.fork,
              archived: where.archived,
              public: where.public,
            },
          }"
        />
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup>
import { useRouteQuery } from "@vueuse/router"
import { useGetReposQuery } from "@/lib/api"
import { usePagination, resetCursor } from "@/lib/pagination"
import { useSorter } from "@/lib/sorter"

const cursor = useRouteQuery("cur", null)
const labels = useRouteQuery("label", [])
const archived = useRouteQuery("archived", true)
const forks = useRouteQuery("forks", false)
const search = useRouteQuery("q", "")
const filterSearch = refDebounced(search, 300)
const direction = useRouteQuery("dir", "desc")
const field = useRouteQuery("sort", "pushed_at")
const sorter = useSorter(
  {
    pushed_at: "updated",
    star_count: "stars",
    created_at: "created",
    name: "name",
  },
  direction,
  field
)

resetCursor(cursor, [labels, archived, forks, search, direction, field])

const where = ref({
  or: [
    { fullNameContainsFold: filterSearch },
    { descriptionContainsFold: filterSearch },
    { homepageContainsFold: filterSearch },
  ],
  hasLabelsWith: computed(() => (labels.value.length ? { nameIn: labels.value } : null)),
  fork: computed(() => (forks.value ? null : false)),
  archived: computed(() => (archived.value == "false" ? false : null)),
})

const { data, error, fetching } = useGetReposQuery({
  variables: {
    ...usePagination(cursor, 10),
    ...sorter.filter,
    where: where,
  },
})
</script>

<style scoped>
.grid {
  @apply grid-cols-1 md:grid-cols-[1fr,240px];
}
</style>
