<template>
  <n-anchor
    v-if="element"
    :top="90"
    style="z-index: 1"
    :bound="100"
    ignore-gap
    class="invisible xl:visible fixed right-[10px] w-[250px]"
  >
    <component :is="item" v-for="(item, index) in links" :key="index" />
  </n-anchor>
</template>

<script setup lang="ts">
import { NAnchorLink } from "naive-ui"
import { createTOC } from "@/lib/toc"

const props = defineProps<{
  element: Node
}>()

const links = computed(() => createTOC(props.element, NAnchorLink))
</script>

<style scoped>
.n-anchor :deep(a) {
  white-space: pre-wrap !important;
  line-height: 1.2em;
  padding-bottom: 4px;
}
.n-anchor :deep(a)::before {
  content: "·";
  position: relative;
  right: 3px;
  font-weight: bold;
  font-size: 1.5em;
  top: 2px;
  @apply text-pink-500;
}
</style>
