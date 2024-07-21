<script setup lang="ts">
import type { SortOrder } from "@/lib/http"

const props = defineProps<{
  // id -> display name
  sortOptions: Record<string, string>
}>()

const sort = defineModel<string>()
const order = defineModel<SortOrder>("order")

function toggle(field: string) {
  if (sort.value == field) {
    order.value = order.value == "asc" ? "desc" : "asc"
  } else {
    sort.value = field
    order.value = "desc"
  }
}
</script>

<template>
  <n-space size="small" inline v-bind="$attrs">
    <n-tag
      v-for="(displayName, id) in props.sortOptions"
      :key="id"
      :type="sort == id ? 'success' : 'default'"
      class="cursor-pointer"
      @click="toggle(id)"
    >
      <template #avatar>
        <n-icon v-if="sort == id" class="">
          <i-mdi-arrow-up v-if="order == 'asc'" />
          <i-mdi-arrow-down v-else />
        </n-icon>
      </template>

      {{ displayName }}
    </n-tag>
  </n-space>
</template>

<style scoped></style>
