<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router"
import { listGithubRepositories } from "@/lib/http/services.gen"
import type { GithubRepositorySortableFields } from "@/lib/http/types.gen"
import CoreIntermediaryCheckbox from "@/components/core/core-intermediary-checkbox.vue"

definePage({
  meta: {
    title: "Repos",
    layout: "default",
  },
})

const sortOptions = {
  pushed_at: "updated",
  star_count: "stars",
  created_at: "created",
  name: "name",
}

const pagination = usePagination<GithubRepositorySortableFields>({ sort: "pushed_at" })
const labels = useRouteQuery<Array<string>>("label", [])
const archived = useRouteQuery<string>("archived", "")
const forks = useRouteQuery<string>("forks", "")
const search = useRouteQuery<string>("q", "")
const debounceSearch = refDebounced<string>(search, 300)

resetPagination(pagination, [labels, archived, forks, debounceSearch])

const {
  data: repos,
  error,
  isPending,
} = useQuery({
  queryKey: ["posts", "admin", pagination, labels, archived, forks, debounceSearch],
  queryFn: () =>
    unwrapErrors(
      listGithubRepositories({
        query: {
          page: pagination.page.value,
          per_page: pagination.perPage.value,
          sort: pagination.sort.value,
          order: pagination.order.value,
          "archived.eq": archived.value == "true" ? true : archived.value == "false" ? false : null,
          "fork.eq": forks.value == "true" ? true : forks.value == "false" ? false : null,
          "fullName.ihas": debounceSearch.value,
          "label.name.in": labels.value.length > 0 ? labels.value : null,
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
          placeholder="Search for a repo"
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
        :value="repos?.content"
        type="repo"
        linkable
        show-empty
        divider
        :loading="isPending"
      />

      <div class="flex">
        <n-pagination
          v-model:page="pagination.page.value"
          v-model:page-size="pagination.perPage.value"
          :item-count="repos?.total_count"
          :page-sizes="pagination.sizes"
          show-size-picker
          class="mt-2 ml-auto"
        />
      </div>
    </div>

    <div>
      <div class="text-center md:text-left">
        <div class="text-emerald-500">Sort repos</div>
        <CoreSorter
          v-if="pagination"
          v-model="pagination.sort.value"
          v-model:order="pagination.order.value"
          :sort-options="sortOptions"
          class="pb-4"
        />

        <div class="text-emerald-500">Filter attributes</div>
        <n-space size="small" class="pb-4" inline>
          <CoreIntermediaryCheckbox v-model="archived">archived</CoreIntermediaryCheckbox>
          <CoreIntermediaryCheckbox v-model="forks">forks</CoreIntermediaryCheckbox>
        </n-space>

        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect v-model="labels" />
      </div>
    </div>
  </div>
</template>
