<script setup lang="ts">
const props = defineProps<{
  value: number
}>()
</script>

<template>
  <UBadge
    class="text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500"
    v-bind="$attrs"
    color="neutral"
    variant="outline"
    size="lg"
  >
    {{ props.value.toLocaleString() }}
    {{ props.value === 1 ? "view" : "views" }}
  </UBadge>
</template>
