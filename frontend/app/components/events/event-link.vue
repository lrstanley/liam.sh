<script setup lang="ts">
const props = defineProps<{
  href: string
  value?: string
}>()

const value = computed(() => (props.value ? props.value : props.href))
</script>

<template>
  <a :href="props.href.includes('://') ? props.href : 'https://github.com/' + props.href" target="_blank"
    :class="{ truncate: props.value ? props.value.length > 8 : props.href.length > 8 }" :title="value">
    <slot>{{ value }}</slot>
  </a>
</template>
