<script setup lang="ts">
const valueString = defineModel<string>({ required: true })
const props = defineProps<{
  intermediary?: boolean
}>()

const result = computed(() => {
  if (valueString.value === "true") {
    return true
  } else if (valueString.value === "false") {
    return false
  } else if (valueString.value === "") {
    if (props.intermediary) {
      return null
    } else {
      return true
    }
  }
  return null
})

function click() {
  if (result.value === null) {
    valueString.value = "true"
  } else if (result.value === true) {
    valueString.value = "false"
  } else if (result.value === false) {
    if (props.intermediary) {
      valueString.value = ""
    } else {
      valueString.value = "true"
    }
  } else {
    valueString.value = props.intermediary ? "" : "true"
  }
}
</script>

<template>
  <UButton default-value="indeterminate" @click="click()" :icon="result === null ? 'lucide:square-minus' : result === true ? 'lucide:square-check' : 'lucide:square'
    " :color="result === true ? 'primary' : result === false ? 'error' : 'neutral'" variant="subtle"
    class="cursor-pointer" v-bind="$attrs">
    <slot />
  </UButton>
</template>
