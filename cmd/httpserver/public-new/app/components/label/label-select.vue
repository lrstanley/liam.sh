<script setup lang="ts">
import { h } from "vue"
import { NBadge } from "naive-ui"
import type { ComponentProps } from "@/lib/util/vueprops"

const {
  modelValue = () => [],
  field = "name",
  suggest = "",
} = defineProps<{
  modelValue: Label[keyof Label] | Label[keyof Label][]
  field?: keyof Label
  suggest?: string
}>()

const emit = defineEmits(["update:modelValue"])

const selected = computed<Label[keyof Label][]>({
  get: () => (Array.isArray(modelValue) ? modelValue : [modelValue]),
  set: (val) => emit("update:modelValue", val),
})

const {
  data: labels,
  isFetching,
  refetch,
} = useQuery({
  queryKey: ["labels", "count"],
  queryFn: () => unwrapErrors(getLabelsCount()),
})

defineExpose({ refetch: refetch })

type RenderedLabel = {
  label: string
  value: ComponentProps<typeof NBadge>["value"]
  data: LabelCount
  popularity: number
}

const options = computed<RenderedLabel[]>(() => {
  return labels.value?.map((label) => ({
    label: label.name,
    value: label[field],
    popularity: label.total_count,
    data: label,
  }))
})

function renderLabel(option: RenderedLabel) {
  return [
    h(NBadge, {
      "show-zero": true,
      color: "grey",
      style: { "margin-right": "1ch" },
      value: option.popularity,
    }),
    option.label,
  ]
}

const suggestions = ref<RenderedLabel[]>([])
const suggestion = computed(() => suggest)
watchDebounced(suggestion, makeSuggestions, { debounce: 300, maxWait: 700, immediate: true })
watchDebounced(selected, makeSuggestions, { debounce: 300, maxWait: 700, immediate: true })

function makeSuggestions(val) {
  if (!val || !options.value) return
  const newSuggestions: RenderedLabel[] = []

  for (const option of options.value) {
    if (selected.value.includes(option.data[field])) continue

    if (suggestion.value.match(new RegExp(`(^|\\W)${option.data.name}(\\W|$)`, "ig"))) {
      newSuggestions.push(option)
    }
  }

  suggestions.value = newSuggestions
}

function addSuggestion(option: RenderedLabel) {
  selected.value.push(option.data[field])
  makeSuggestions(suggestion.value)
}
</script>

<template>
  <n-select
    v-bind="$attrs"
    v-model:value="selected"
    :options="options"
    :render-label="renderLabel"
    :loading="isFetching"
    clearable
    filterable
    multiple
    placeholder="Select labels"
  />
  <div v-if="suggestions.length > 0" v-motion-fade class="flex flex-col mb-2">
    <span class="text-emerald-500">Suggestions</span>
    <div class="inline-flex flex-row flex-wrap gap-1">
      <n-tag
        v-for="label in suggestions"
        :key="label.data.id"
        class="cursor-pointer hover:bg-emerald-700"
        @click="addSuggestion(label)"
      >
        <n-badge show-zero color="grey" class="mr-0" :value="label.popularity" />
        {{ label.data.name }}
      </n-tag>
    </div>
  </div>
</template>
