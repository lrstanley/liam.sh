<script setup lang="ts" generic="T">
import type { SchemaPagedResponse } from '#open-fetch-schemas/api'

const props = defineProps<{
  loading?: boolean
  resp?: SchemaPagedResponse | null
  pagination: ReturnType<typeof usePagination>
}>()

const page = computed({
  get: () => props.pagination.page.value || 1,
  set: (v) => {
    props.pagination.page.value = v
  },
})

const total = ref<number>(0)

watch(
  () => props.resp,
  (resp) => {
    if (resp) {
      total.value = resp.total_count
    }
  },
  { immediate: true }
)
</script>

<template>
  <motion as="div" layout :initial="{ opacity: 0 }" :animate="{ opacity: 1 }">
    <UPagination v-model:page="page" :items-per-page="pagination.perPage.value" :sibling-count="1"
      active-color="primary" active-variant="subtle" :show-controls="false" show-edges :total="total"
      :disabled="loading" v-bind="$attrs" />
  </motion>
</template>
