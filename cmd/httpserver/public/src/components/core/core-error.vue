<script setup lang="ts">
const props = defineProps<{
  error: Error
}>()
</script>

<template>
  <n-alert v-if="props.error" v-motion-fade title="An error occurred" type="error" class="m-2 md:m-6">
    {{ props.error.message || props.error }}
  </n-alert>
</template>
