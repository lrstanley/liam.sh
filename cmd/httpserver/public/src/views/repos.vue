<template>
  <LayoutDefault :error="error">
    <div class="grid gap-5 md:gap-25 mt-8">
      <div class="order-last md:order-first">
        <n-input
          v-model:value="search"
          type="text"
          placeholder="Search for a repo"
          :loading="fetching"
          class="mb-8"
        >
          <template #prefix>
            <n-icon>
              <i-mdi-search />
            </n-icon>
          </template>
        </n-input>

        <ObjectRender
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
          <n-checkbox @update:checked="(v) => (where.archived = v ? null : false)">
            include archived
          </n-checkbox>
          <n-checkbox @update:checked="(v) => (where.fork = v ? null : false)">
            include forks
          </n-checkbox>
        </n-space>

        <div v-if="allLabels">
          <div class="text-emerald-500">Filter by label</div>
          <n-space size="small" inline>
            <n-tag
              v-for="label in allLabels"
              :key="label.id"
              :type="labels.includes(label.name) ? 'success' : ''"
              class="cursor-pointer"
              @click="toggleLabel(label.name)"
            >
              {{ label.name }}
            </n-tag>
          </n-space>
        </div>
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup>
import { useRouteQuery } from "@vueuse/router"
import { useGetReposQuery, useGetLabelsQuery } from "@/lib/api"

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

const search = useRouteQuery("q", "")
const filterSearch = refDebounced(search, 300)

const labels = useRouteQuery("label", [])
const filterLabels = computed(() => (labels.value.length ? { nameIn: labels.value } : null))

const labelQuery = useGetLabelsQuery({
  variables: {
    where: {
      hasGithubRepositories: true,
    },
  },
})
const allLabels = computed(() => labelQuery.data.value?.labels.edges.map(({ node }) => node))

function toggleLabel(name) {
  let active = labels.value
  if (!Array.isArray(labels.value)) {
    active = [labels.value]
  }

  if (active.includes(name)) {
    active = active.filter((l) => l !== name)
  } else {
    active = [...active, name]
  }

  labels.value = active
}

const where = ref({
  or: [
    { fullNameContainsFold: filterSearch },
    { descriptionContainsFold: filterSearch },
    { homepageContainsFold: filterSearch },
  ],
  hasLabelsWith: filterLabels,
  public: true,
  fork: false,
  archived: false,
})

const { data, error, fetching } = useGetReposQuery({
  variables: {
    first: 100,
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
