<template>
  <LayoutDefault :error="error">
    <div class="grid gap-5 md:gap-25 mt-8">
      <div class="order-last md:order-first">
        <div class="mb-8 flex flex-auto gap-2">
          <n-input
            v-model:value="search"
            :loading="fetching"
            type="text"
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
      <div>
        <div class="text-emerald-500">Sort repos</div>
        <n-space size="small" class="pb-4" inline>
          <n-tag
            v-for="(name, key) in orderOptions.fields"
            :key="key"
            :type="orderOptions.selected.field.value == key ? 'success' : ''"
            class="cursor-pointer"
            @click="changeOrder(key)"
          >
            <template #avatar>
              <n-icon v-if="orderOptions.selected.field.value == key" class="">
                <i-mdi-arrow-up v-if="orderOptions.selected.direction.value == 'asc'" />
                <i-mdi-arrow-down v-else />
              </n-icon>
            </template>

            {{ name }}
          </n-tag>
        </n-space>

        <div class="text-emerald-500">Filter attributes</div>
        <n-space size="small" class="pb-4" inline>
          <n-checkbox :checked="archived == 'true'" @update:checked="(v) => (archived = !!v)">
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
import { usePagination } from "@/lib/pagination"

const orderOptions = {
  directions: ["desc", "asc"],
  fields: {
    star_count: "stars",
    pushed_at: "updated",
    created_at: "created",
    name: "name",
  },
  selected: {
    direction: useRouteQuery("dir", "desc"),
    field: useRouteQuery("sort", "pushed_at"),
  },
  computed: {
    direction: computed(() => {
      return orderOptions.directions.includes(orderOptions.selected.direction.value)
        ? orderOptions.selected.direction.value.toUpperCase()
        : "DESC"
    }),
    field: computed(() => {
      return Object.keys(orderOptions.fields).includes(orderOptions.selected.field.value)
        ? orderOptions.selected.field.value.toUpperCase()
        : "DATE"
    }),
  },
}

function changeOrder(key) {
  if (key === orderOptions.selected.field.value) {
    orderOptions.selected.direction.value =
      orderOptions.selected.direction.value === "desc" ? "asc" : "desc"
    return
  }

  // https://github.com/vueuse/vueuse/issues/901
  orderOptions.selected.field.value = key
  setTimeout(() => {
    nextTick(() => {
      orderOptions.selected.direction.value = "desc"
    })
  }, 1)
}

const cursor = useRouteQuery("cur", null)
const labels = useRouteQuery("label", [])
const archived = useRouteQuery("archived", false)
const forks = useRouteQuery("forks", false)
const search = useRouteQuery("q", "")
const filterSearch = refDebounced(search, 300)

const where = ref({
  or: [
    { fullNameContainsFold: filterSearch },
    { descriptionContainsFold: filterSearch },
    { homepageContainsFold: filterSearch },
  ],
  hasLabelsWith: computed(() => (labels.value.length ? { nameIn: labels.value } : null)),
  public: true,
  fork: computed(() => (forks.value ? null : false)),
  archived: computed(() => (archived.value ? null : false)),
})

const { data, error, fetching } = useGetReposQuery({
  variables: {
    ...usePagination(cursor, 10),
    where: where,
    order: orderOptions.computed.direction,
    orderBy: orderOptions.computed.field,
  },
})
</script>

<style scoped>
.grid {
  @apply grid-cols-1 md:grid-cols-[1fr,240px];
}
</style>
