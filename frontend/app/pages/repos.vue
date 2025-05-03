<script setup lang="ts">
import { useRouteQuery } from "@vueuse/router"

definePageMeta({
  title: "Repos",
  layout: "default",
})

const labels = useRouteQuery<Array<string>>("label", [])
const archived = useRouteQuery<string>("archived", "")
const forks = useRouteQuery<string>("forks", "")
const search = useRouteQuery<string>("q", "")
const debounceSearch = refDebounced<string>(search, 300)
const pagination = usePagination<GithubRepositorySortableFields>({
  sort: "pushed_at",
  resetChanged: [labels, archived, forks, debounceSearch],
})

const {
  data: repos,
  error,
  status,
} = await listGithubRepositories({
  composable: "useFetch",
  query: computed(() => ({
    page: pagination.page.value,
    per_page: pagination.perPage.value,
    sort: pagination.sort.value,
    order: pagination.order.value,
    "archived.eq": archived.value == "true" ? true : archived.value == "false" ? false : undefined,
    "fork.eq": forks.value == "true" ? true : forks.value == "false" ? false : undefined,
    "fullName.ihas": debounceSearch.value.length > 0 ? debounceSearch.value : undefined,
    "label.name.in": labels.value.length > 0 ? labels.value : undefined, // TODO: in vs has
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
        placeholder="Search for a repo"
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

      <CorePagination
        class="ml-auto"
        :loading="status === 'pending'"
        :resp="repos"
        :pagination="pagination"
      />
    </div>

    <CoreObjectRender
      :value="repos?.content"
      type="repo"
      linkable
      show-empty
      divider
      :loading="status === 'pending'"
    />

    <template #sidebar>
      <div>
        <div class="text-(--ui-primary)">Sort repos</div>
        <CoreSorter
          v-if="pagination"
          v-model="pagination.sort.value"
          v-model:order="pagination.order.value"
          :sort-options="{
            pushed_at: { displayName: 'updated', defaultOrder: 'desc' },
            star_count: { displayName: 'stars', defaultOrder: 'desc' },
            created_at: { displayName: 'created', defaultOrder: 'desc' },
            name: { displayName: 'name', defaultOrder: 'asc' },
          }"
        />
      </div>

      <div>
        <div class="text-(--ui-primary)">Filter attributes</div>
        <div class="inline-flex flex-row gap-1">
          <CoreIndeterminateButton v-model:string-value="archived" size="xs" label="archived" />
          <CoreIndeterminateButton v-model:string-value="forks" size="xs" label="forks" />
        </div>
      </div>

      <div>
        <div class="text-(--ui-primary)">Filter by label</div>
        <LabelSelect v-model="labels" />
      </div>
    </template>
  </ContainerStickySidebar>
</template>
