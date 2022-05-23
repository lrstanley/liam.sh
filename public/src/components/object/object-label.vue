<template>
  <component
    :is="props.linkable ? 'router-link' : 'span'"
    :to="{ name: 'posts', query: { label: props.value.name } }"
  >
    <n-tag v-bind="$attrs" class="hover:bg-emerald-700 cursor-pointer">
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
