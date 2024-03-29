<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router"
import { useGetReposQuery } from "@/lib/api"
import { usePagination, resetCursor, useSorter } from "@/lib/util"

definePage({
  meta: {
    title: "Repos",
    layout: "default",
  },
})

const cursor = useRouteQuery<string>("cur", null)
const labels = useRouteQuery<Array<string>>("label", [])
const archived = useRouteQuery<string>("archived", "true")
const forks = useRouteQuery<string>("forks", "false")
const search = useRouteQuery<string>("q", "")
const filterSearch = refDebounced<string>(search, 300)
const direction = useRouteQuery<string>("dir", "desc")
const field = useRouteQuery<string>("sort", "pushed_at")
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
  fork: computed(() => (forks.value == "true" ? null : false)),
  archived: computed(() => (archived.value == "false" ? false : null)),
})

const { data, error, fetching } = await useGetReposQuery({
  variables: {
    ...usePagination(cursor, 10),
    ...sorter.filter,
    where: where,
  },
})

watch(error, () => {
  if (error.value) throw error.value
})
</script>

<template>
  <div class="mt-8 grid-sidebar">
    <div>
      <div class="flex flex-auto gap-2 mt-1 mb-8">
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

        <CorePagination v-model="cursor" :page-info="data?.githubRepositories?.pageInfo" />
      </div>

      <CoreObjectRender
        v-if="data?.githubRepositories"
        :value="data.githubRepositories"
        linkable
        show-empty
        divider
      />
    </div>

    <div>
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort repos</div>
        <CoreSorter :sorter="sorter" class="pb-4" />

        <div class="text-emerald-500">Filter attributes</div>
        <n-space size="small" class="pb-4" inline>
          <n-checkbox
            :checked="archived == 'true'"
            @update:checked="(v) => (archived = v ? 'true' : 'false')"
          >
            include archived
          </n-checkbox>
          <n-checkbox :checked="forks == 'true'" @update:checked="(v) => (forks = v ? 'true' : 'false')">
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
            },
          }"
        />
      </div>
    </div>
  </div>
</template>
