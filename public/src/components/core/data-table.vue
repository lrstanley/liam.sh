<template>
  <n-data-table v-bind="$props" :columns="columns" />
</template>

<script setup>
import { h } from "vue"
import { titleCase } from "@/lib/core/util"

const props = defineProps({
  headers: {
    type: Object,
    default: () => ({}),
    required: true,
  },
})

const slots = useSlots()
const columns = Object.keys(slots).map((name) => {
  return {
    title: name in props.headers ? props.headers[name] : titleCase(name),
    key: name,
    render: (row) => {
      return h(slots[name], { row: row })
    },
  }
})
</script>
