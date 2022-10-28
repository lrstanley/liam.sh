<template>
  <component
    :is="props.linkable ? 'router-link' : 'span'"
    :to="{ name: props.route, query: { [props.query]: label.name } }"
  >
    <n-tag v-bind="$attrs" class="cursor-pointer hover:bg-emerald-700">
      {{ label.name }}
    </n-tag>
  </component>
</template>

<script setup lang="ts">
import type { Label } from "@/lib/api"

const props = withDefaults(
  defineProps<{
    linkable?: boolean
    value: Label
    route?: string
    query?: string
  }>(),
  {
    route: "/posts",
    query: "label",
  }
)

const label = ref(props.value)
</script>
