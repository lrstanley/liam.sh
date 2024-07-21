<script setup lang="ts">
const value = defineModel<string>()

function intermediary(v: string): string {
  if (v == "true") return "false"
  if (v == "false") return ""
  return "true"
}
</script>

<template>
  <n-checkbox
    :checked="value == 'true'"
    :indeterminate="value == ''"
    @update:checked="(v) => (value = intermediary(value))"
  >
    <slot />
  </n-checkbox>
</template>
