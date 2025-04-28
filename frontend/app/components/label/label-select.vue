<script setup lang="ts">
import { UBadge } from "#components"

const toast = useToast()

export type LabelField = keyof Pick<Label, "id" | "name">

const {
  modelValue: value = [],
  field = "name",
  suggest = "",
  allowCreate,
} = defineProps<{
  modelValue: Label[LabelField] | Label[LabelField][]
  field?: LabelField
  suggest?: string
  allowCreate?: boolean
}>()

const emit = defineEmits<{
  "update:modelValue": [value: Label[LabelField][]]
}>()

const { data: labels, status, refresh } = await getLabelsCount({ composable: "useFetch" })

defineExpose({ refetch: refresh })

const label2rendered = (label: LabelCount): RenderedLabel => {
  return {
    label: label.name,
    value: label[field],
    popularity: label.total_count,
    data: label,
  }
}

const selected = computed<RenderedLabel[]>({
  get: () => {
    const values = Array.isArray(value) ? value : [value]
    return labels.value?.filter((label) => values.includes(label[field])).map(label2rendered) || []
  },
  set: (val) => {
    emit(
      "update:modelValue",
      val.map((label) => label.data[field])
    )
  },
})

type RenderedLabel = {
  label: string
  value: Label[LabelField]
  data: LabelCount
  popularity: number
}

const options = computed<RenderedLabel[]>(() => {
  return labels.value?.map(label2rendered) || []
})

// const suggestions = ref<RenderedLabel[]>([])
// const suggestion = computed(() => suggest)
// watchDebounced(suggestion, makeSuggestions, { debounce: 300, maxWait: 700, immediate: true })
// watchDebounced(selected, makeSuggestions, { debounce: 300, maxWait: 700, immediate: true })

// TODO: any?
// function makeSuggestions(val: any) {
//   if (!val || !options.value) return
//   const newSuggestions: RenderedLabel[] = []

//   for (const option of options.value) {
//     if (selected.value.includes(option.data[field])) continue

//     if (suggestion.value.match(new RegExp(`(^|\\W)${option.data.name}(\\W|$)`, "ig"))) {
//       newSuggestions.push(option)
//     }
//   }

//   suggestions.value = newSuggestions
// }

// function addSuggestion(option: RenderedLabel) {
//   selected.value.push(option.data[field])
//   makeSuggestions(suggestion.value)
// }

function clear(e: any) {
  e.stopPropagation()
  selected.value = []
}

async function onCreate(name: string) {
  try {
    const label = await createLabel({
      composable: "$fetch",
      body: { name },
    })

    toast.add({
      title: "Label created",
      description: `Label ${label.name} created`,
      color: "success",
      duration: 3000,
    })
    await refresh()
  } catch (e) {
    toast.add({
      title: "Failed to create label",
      description: (e as any).message,
      color: "error",
      duration: 5000,
    })
  }
}
</script>

<template>
  <USelectMenu
    v-model="selected"
    :items="options"
    :loading="status === 'pending'"
    icon="lucide:tag"
    multiple
    :create-item="allowCreate"
    @create="onCreate"
    placeholder="Select labels"
    class="w-full lg:max-w-[15vw]"
  >
    <template #leading="{ modelValue, ui }">
      <span v-if="modelValue && modelValue.length > 0" class="group">
        <UBadge
          :color="modelValue.length >= 3 ? 'error' : 'success'"
          size="sm"
          variant="soft"
          class="inline-block rounded-full group-hover:hidden"
        >
          {{ modelValue.length }}
        </UBadge>
        <UBadge
          color="neutral"
          size="sm"
          variant="subtle"
          class="hidden rounded-full cursor-pointer group-hover:inline-block"
          @click="clear"
        >
          x
        </UBadge>
      </span>
    </template>
    <template #item-leading="{ item }">
      <UBadge
        :color="
          item.popularity > 25
            ? 'primary'
            : item.popularity > 10
              ? 'info'
              : item.popularity > 0
                ? 'neutral'
                : 'error'
        "
        variant="soft"
        class="rounded-full"
        size="sm"
      >
        {{ item.popularity }}
      </UBadge>
    </template>
  </USelectMenu>
  <!-- <div v-if="suggestions.length > 0" v-motion-fade class="flex flex-col mb-2">
    <span class="text-(--ui-primary)">Suggestions</span>
    <div class="inline-flex flex-row flex-wrap gap-1">
      <UBadge
        v-for="label in suggestions"
        :key="label.data.id"
        class="cursor-pointer hover:bg-(--ui-color-primary-700)"
        @click="addSuggestion(label)"
      >
        <UBadge color="neutral" variant="outline" class="mr-0">{{ label.popularity }}</UBadge>
        {{ label.data.name }}
      </UBadge>
    </div>
  </div> -->
</template>
