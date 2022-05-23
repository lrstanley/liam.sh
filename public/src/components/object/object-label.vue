<template>
  <component :is="props.linkable ? 'router-link' : 'span'">
    <n-tag v-bind="$attrs">
      {{ props.value.name }}
    </n-tag>
  </component>
</template>

<script setup>
const props = defineProps({
  linkable: {
    type: Boolean,
    default: false,
  },
  value: {
    type: Object,
    required: true,
  },
})
</script>

<style scoped></style>
