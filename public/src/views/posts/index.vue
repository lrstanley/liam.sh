<template>
  <LayoutDefault :error="error">
    <div class="grid gap-4 mt-8">
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
        <div class="col-auto">
          <ObjectRender v-if="data?.posts" :value="data.posts" linkable show-empty />
        </div>
      </div>
      <div class="col-auto">
        <n-button type="primary" class="mb-8" @click="labels = [...labels, 'test']"> New post </n-button>
        <n-card> test2 </n-card>
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup>
import { useRouteQuery } from "@vueuse/router"
import { useGetPostsQuery } from "@/lib/api"

const search = useRouteQuery("q", "")
const labels = useRouteQuery("label", [])
const filterSearch = refDebounced(search, 300)

const { data, error, fetching } = useGetPostsQuery({
  variables: {
    first: 10,
    where: {
      or: [{ titleContainsFold: filterSearch }, { summaryContainsFold: filterSearch }],
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
