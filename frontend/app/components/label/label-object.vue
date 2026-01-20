<script setup lang="ts">
import { NuxtLink } from "#components"
import type { SchemaLabel } from '#open-fetch-schemas/api'
import type { RouteNamedMap } from "vue-router/auto-routes"

const {
  linkable,
  value: label,
  route = "/posts",
  query = "label",
} = defineProps<{
  linkable?: boolean
  value: SchemaLabel
  route?: RouteNamedMap[keyof RouteNamedMap]["path"]
  query?: string
}>()
const { width } = useWindowSize()
</script>

<template>
  <component :is="linkable ? NuxtLink : 'span'" :to="{ path: route, query: { [query]: label.name } }">
    <UButton color="neutral" variant="outline" :size="width <= 640 ? 'md' : 'xs'" v-bind="$attrs"
      class="cursor-pointer">
      {{ label.name }}
    </UButton>
  </component>
</template>
