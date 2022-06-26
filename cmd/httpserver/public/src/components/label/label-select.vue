<template>
  <n-select
    v-bind="$attrs"
    v-model:value="selected"
    :options="options"
    :render-label="renderLabel"
    :loading="labels.fetching?.value"
    clearable
    filterable
    multiple
    placeholder="Select labels"
  />
  <div v-if="suggestions.length > 0" v-motion-fade class="flex flex-col mb-2">
    <p class="text-center">Suggestions</p>
    <div class="inline-flex flex-row flex-wrap gap-1">
      <n-tag
        v-for="label in suggestions"
        :key="label.data.id"
        class="hover:bg-emerald-700 cursor-pointer"
        @click="addSuggestion(label)"
      >
        <n-badge show-zero color="grey" class="mr-[-2ch]" :value="label.popularity" />
        {{ label.data.name }}
      </n-tag>
    </div>
  </div>
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
  suggest: {
    type: String,
    default: "",
  },
})
const emit = defineEmits(["update:modelValue"])

const selected = computed({
  get: () => (Array.isArray(props.modelValue) ? props.modelValue : [props.modelValue]),
  set: (val) => emit("update:modelValue", val),
})

const labels = useGetLabelsQuery({ variables: { where: props.where } })
defineExpose({ refetch: labels.executeQuery })

const options = computed(() =>
  labels.data?.value?.labels.edges
    .map(({ node }) => ({
      label: node.name,
      value: node[props.field],
      popularity: node.githubRepositories.totalCount + node.posts.totalCount,
      data: node,
    }))
    .sort((a, b) => b.popularity - a.popularity)
)

function renderLabel(option) {
  return [
    h(NBadge, {
      "show-zero": true,
      color: "grey",
      style: { "margin-right": "1ch" },
      value: option.popularity,
    }),
    option.data?.name,
  ]
}

const suggestions = ref([])
const suggest = computed(() => props.suggest)
watchDebounced(suggest, makeSuggestions, { debounce: 300, maxWait: 700, immediate: true })
watchDebounced(selected, makeSuggestions, { debounce: 300, maxWait: 700, immediate: true })

function makeSuggestions(val) {
  if (!val || !options.value) return
  let newSuggestions = []

  for (const option of options.value) {
    if (selected.value.includes(option.data[props.field])) continue

    if (suggest.value.match(new RegExp(`(^|\\W)${option.data.name}(\\W|$)`, "ig"))) {
      newSuggestions.push(option)
    }
  }

  suggestions.value = newSuggestions
}

function addSuggestion(option) {
  selected.value.push(option.data[props.field])
  makeSuggestions(suggest.value)
}
</script>
