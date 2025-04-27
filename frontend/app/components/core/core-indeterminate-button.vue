<script setup lang="ts">
const value = defineModel<null | boolean>()
const valueString = defineModel<string | undefined>("stringValue") // string version of the value depending on usecase.

const { intermediary = true } = defineProps<{
  intermediary?: boolean
}>()

// keep them both in sync. if one is update, update the other, but make sure not to trigger an update
// if the value is already exactly the same, otherwise infinite loops.
watch(
  value,
  (v) => {
    const newValue = v === null ? undefined : v ? "true" : "false"

    if (valueString.value !== newValue) {
      valueString.value = newValue
    }
  },
  { immediate: true }
)

watch(
  valueString,
  (v) => {
    let newValue: null | boolean = null
    if (v === "true") {
      newValue = true
    } else if (v === "false") {
      newValue = false
    }

    if (value.value !== newValue) {
      value.value = newValue
    }
  },
  { immediate: true }
)
</script>

<template>
  <UButton
    default-value="indeterminate"
    @click="value = value === null ? true : value === true ? false : intermediary ? null : true"
    :icon="
      value === null ? 'lucide:square-minus' : value === true ? 'lucide:square-check' : 'lucide:square'
    "
    :color="value === true ? 'primary' : value === false ? 'error' : 'neutral'"
    variant="subtle"
    class="cursor-pointer"
    v-bind="$attrs"
  >
    <slot />
  </UButton>
</template>
