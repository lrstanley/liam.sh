<template>
  <n-select
    v-bind="$attrs"
    v-model:value="selected"
    :options="options"
    :loading="labels.fetching?.value"
    clearable
    filterable
    multiple
    placeholder="Select labels"
  />
</template>

<script setup>
import { h } from "vue"
import { NBadge } from "naive-ui"
import { useGetLabelsQuery } from "@/lib/api"

const props = defineProps({
  modelValue: {
    type: [Array, String, Number],
    required: true,
    default: () => [],
  },
  field: {
    type: String,
    default: "name",
  },
  where: {
    type: Object,
    default: () => null,
  },
})
const emit = defineEmits(["update:modelValue"])

const selected = computed({
  get: () => (Array.isArray(props.modelValue) ? props.modelValue : [props.modelValue]),
  set: (val) => emit("update:modelValue", val),
})

const labels = useGetLabelsQuery({ variables: { where: props.where } })
const options = computed(() =>
  labels.data?.value?.labels.edges
    .map(({ node }) => ({
      label: (option) => {
        return [
          h(NBadge, {
            "show-zero": true,
            color: "grey",
            style: { "margin-right": "1ch" },
            value: option.popularity,
          }),
          option.data.name,
        ]
      },
      value: node[props.field],
      popularity: node.githubRepositories.totalCount + node.posts.totalCount,
      data: node,
    }))
    .sort((a, b) => b.popularity - a.popularity)
)
</script>
