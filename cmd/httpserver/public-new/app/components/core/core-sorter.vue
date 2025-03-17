<script setup lang="ts">
export type SortConfig = {
  displayName?: string // uses ID if not provided.
  defaultOrder: "asc" | "desc"
}

const { sortOptions } = defineProps<{
  sortOptions: Record<string, SortConfig>
}>()

const field = defineModel<string>()
const order = defineModel<"asc" | "desc">("order")

function toggle(id: string) {
  if (field.value == id) {
    order.value = order.value == "asc" ? "desc" : "asc"
  } else {
    field.value = id
    order.value = sortOptions[id]?.defaultOrder || "desc"
  }
}
</script>

<template>
  <div class="inline-flex flex-wrap gap-1">
    <UButton
      v-for="(config, id) in sortOptions"
      :key="id"
      :color="field == id ? 'primary' : 'neutral'"
      variant="outline"
      :icon="field == id ? (order == 'asc' ? 'mdi:arrow-up' : 'mdi:arrow-down') : undefined"
      @click="toggle(id)"
      class="cursor-pointer"
    >
      {{ config.displayName || id }}
    </UButton>
  </div>
</template>
