<template>
  <div v-show="element" v-bind="$attrs">
    <div class="text-emerald-500">Table of Contents</div>
    <n-anchor :bound="100" ignore-gap :show-background="false">
      <component :is="item" v-for="(item, index) in links" :key="index" />
    </n-anchor>
  </div>
</template>

<script setup lang="ts">
import { NAnchorLink } from "naive-ui"
import { createTOC } from "@/lib/toc"

const props = defineProps<{
  element: Node
}>()

const links = computed(() => createTOC(props.element, NAnchorLink))
</script>
