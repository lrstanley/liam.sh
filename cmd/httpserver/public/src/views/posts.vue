<template>
  <LayoutDefault :error="error">
    <div class="grid gap-5 md:gap-25 mt-8">
      <div class="order-last md:order-first">
        <n-input
          v-model:value="search"
          type="text"
          placeholder="Search for a post"
          :loading="fetching"
          class="mb-8"
        >
          <template #prefix>
            <n-icon>
              <i-mdi-search />
            </n-icon>
          </template>
        </n-input>

        <CoreObjectRender v-if="data?.posts" :value="data.posts" linkable show-empty divider />
      </div>
      <div>
        <div class="text-emerald-500">Sort posts</div>
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

        <div class="text-emerald-500">Filter by label</div>
        <LabelSelect v-model="labels" :where="{ hasPosts: true }" />
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup>
import { useRouteQuery } from "@vueuse/router"
import { useGetPostsQuery } from "@/lib/api"

const orderOptions = {
  directions: ["desc", "asc"],
  fields: {
    date: "date",
    title: "title",
    view_count: "popularity",
  },
  selected: {
    direction: useRouteQuery("dir", "desc"),
    field: useRouteQuery("sort", "date"),
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

const { data, error, fetching } = useGetPostsQuery({
  variables: {
    first: 10,
    where: {
      or: [{ titleContainsFold: filterSearch }, { summaryContainsFold: filterSearch }],
      hasLabelsWith: computed(() => (labels.value.length ? { nameIn: labels.value } : null)),
    },
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
