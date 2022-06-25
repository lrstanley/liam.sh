<template>
  <a
    :href="props.href.includes('://') ? props.href : 'https://github.com/' + props.href"
    target="_blank"
    :class="{ truncate: props.value ? props.value.length > 8 : props.href.length > 8 }"
    class="flex-3"
    :title="value"
  >
    <slot>{{ value }}</slot>
  </a>
</template>

<script setup>
const props = defineProps({
  href: {
    type: String,
    required: true,
  },
  value: {
    type: String,
    required: false,
    default: "",
  },
})

const value = computed(() => (props.value ? props.value : props.href))
</script>
