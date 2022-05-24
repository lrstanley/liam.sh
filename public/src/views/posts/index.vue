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

        <ObjectRender v-if="data?.posts" :value="data.posts" linkable show-empty />
      </div>
      <div v-if="allLabels">
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
  </LayoutDefault>
</template>

<script setup>
import { useRouteQuery } from "@vueuse/router"
import { useGetPostsQuery, useGetLabelsQuery } from "@/lib/api"

const search = useRouteQuery("q", "")
const filterSearch = refDebounced(search, 300)

const labels = useRouteQuery("label", [])
const filterLabels = computed(() => (labels.value.length ? { nameIn: labels.value } : null))

const labelQuery = useGetLabelsQuery({
  variables: {
    where: {
      hasPosts: true,
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

const { data, error, fetching } = useGetPostsQuery({
  variables: {
    first: 10,
    where: {
      or: [{ titleContainsFold: filterSearch }, { summaryContainsFold: filterSearch }],
      hasLabelsWith: filterLabels,
    },
  },
})
</script>

<style scoped>
.grid {
  @apply grid-cols-1 md:grid-cols-[1fr,240px];
}

.post-title {
  @apply text-size-1.5em;
}
</style>
