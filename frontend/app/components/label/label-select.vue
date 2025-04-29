<script setup lang="ts">
import { UBadge } from "#components"

const toast = useToast()

export type LabelField = keyof Pick<Label, "id" | "name">

const {
  modelValue: value = [],
  field = "name",
  suggestText = "",
  allowCreate,
  allowSuggest,
} = defineProps<{
  modelValue: Label[LabelField] | Label[LabelField][]
  field?: LabelField
  suggestText?: string
  allowCreate?: boolean
  allowSuggest?: boolean
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

function clear(e: any) {
  e.stopPropagation()
  selected.value = []
}

async function addLabel(name: string) {
  if (!labels.value?.some((l) => l.name === name)) {
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
      return
    }
  }

  const newLabel = labels.value?.find((l) => l.name === name)
  if (newLabel) {
    selected.value = [...selected.value, label2rendered(newLabel)]
  }
}

const suggestions = ref<SuggestedLabel[]>([])
const suggestionsFiltered = computed(() => {
  const filtered = suggestions.value.filter((s) => !selected.value.some((l) => l.data.name === s.name))

  filtered.sort((a, b) => {
    // sort existing to the start.
    if (!a.existing_id && b.existing_id) return -1
    if (a.existing_id && !b.existing_id) return 1

    // if both are existing, sort by popularity.
    const existingA = labels.value?.find((l) => l.id === a.existing_id)
    const existingB = labels.value?.find((l) => l.id === b.existing_id)
    if (existingA && existingB) return existingB.total_count - existingA.total_count

    // if both are not existing, sort by name.
    return a.name.localeCompare(b.name)
  })

  if (filtered.length > 20) {
    filtered.splice(20)
  }

  return filtered
})
const hasMadeSuggestions = ref(false)
const loadingSuggestions = ref(false)

async function suggest() {
  if (!allowSuggest || !suggestText) return

  loadingSuggestions.value = true
  hasMadeSuggestions.value = true

  try {
    suggestions.value = await suggestLabels({
      composable: "$fetch",
      body: { content: suggestText },
    })
  } catch (e) {
    toast.add({
      title: "Failed to suggest labels",
      description: (e as any).message,
      color: "error",
      duration: 5000,
    })
  } finally {
    loadingSuggestions.value = false
  }
}
</script>

<template>
  <div class="flex flex-col gap-2">
    <div class="flex flex-row gap-1">
      <UButton
        v-if="allowSuggest && suggestText"
        size="md"
        variant="link"
        icon="lucide:wand-sparkles"
        :loading="loadingSuggestions"
        class="group transition-all duration-200 ease-in-out cursor-pointer"
        :ui="{
          leadingIcon:
            '!text-gradient bg-gradient-to-r group-hover:bg-gradient-to-br from-indigo-400 via-fuchsia-400 to-cyan-400 group-hover:scale-110 opacity-90 group-hover:opacity-100 group-hover:-rotate-6 transition-all duration-200 ease-in-out',
        }"
        @click="suggest()"
      />

      <USelectMenu
        v-model="selected"
        :items="options"
        :loading="status === 'pending'"
        icon="lucide:tag"
        multiple
        :create-item="allowCreate"
        @create="addLabel"
        placeholder="Select labels"
        class="max-w-full min-w-0 grid"
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
    </div>

    <div
      v-show="allowSuggest && hasMadeSuggestions && !loadingSuggestions"
      class="flex flex-row flex-wrap gap-1"
      :v-auto-animate="{}"
    >
      <UBadge
        v-if="suggestionsFiltered.length === 0"
        color="secondary"
        variant="subtle"
        class="block w-full text-center"
      >
        No suggestions found
      </UBadge>

      <button
        v-if="suggestionsFiltered.length > 0"
        v-for="suggestion in suggestionsFiltered"
        :key="suggestion.name"
        @click="addLabel(suggestion.name)"
        class="relative inline-flex h-7 overflow-hidden rounded-sm p-[1px] focus:outline-none focus:ring-0 group *:transition-opacity duration-200 ease-in-out"
      >
        <span
          v-if="suggestion.existing_id"
          class="absolute inset-[-1000%] animate-[spin_3s_linear_infinite] bg-[conic-gradient(from_90deg_at_50%_50%,#171717_0%,#18d67d_50%,#171717_100%)] opacity-75 group-hover:opacity-100"
        />
        <span
          v-else
          class="absolute inset-[-1000%] animate-[spin_3s_linear_infinite] bg-[conic-gradient(from_90deg_at_50%_50%,#E2CBFF_0%,#494beb_50%,#E2CBFF_100%)] opacity-75 group-hover:opacity-100"
        />
        <span
          class="inline-flex h-full w-full cursor-pointer items-center justify-center rounded-sm bg-neutral-900 px-1 py-1 text-sm font-medium text-(--ui-text) group-hover:text-white backdrop-blur-3xl"
        >
          <UIcon
            :name="suggestion.existing_id ? 'lucide:plus' : 'lucide:wand-sparkles'"
            class="h-4 w-4 text-(--ui-text) mr-1"
            :class="{
              '!text-gradient bg-gradient-to-r group-hover:bg-gradient-to-br from-indigo-400 via-fuchsia-400 to-cyan-400 opacity-80 group-hover:opacity-100 transition-all duration-200 ease-in-out':
                !suggestion.existing_id,
              'text-(--ui-text-muted)': suggestion.existing_id,
            }"
          />
          {{ suggestion.name }}
        </span>
      </button>
    </div>
  </div>
</template>
