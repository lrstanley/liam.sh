<template>
  <n-tag class="text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500" v-bind="$attrs">
    {{ props.value.toLocaleString() }}
    {{ props.value === 1 ? "view" : "views" }}
  </n-tag>
</template>

<script setup lang="ts">
const props = defineProps<{
  value: number
}>()
</script>
