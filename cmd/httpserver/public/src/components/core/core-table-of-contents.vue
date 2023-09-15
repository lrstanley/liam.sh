<script setup lang="ts">
import { NAnchorLink } from "naive-ui"
import { createTOC } from "@/lib/util"

const props = defineProps<{
  element: Node
}>()

const links = computed(() => createTOC(props.element, NAnchorLink))
</script>

<template>
  <div v-show="links.length" v-bind="$attrs">
    <div class="text-emerald-500">Table of Contents</div>
    <n-anchor :bound="100" ignore-gap :show-background="false">
      <component :is="item" v-for="(item, index) in links" :key="index" />
    </n-anchor>
  </div>
</template>
