<template>
  <div>
    <p class="text-center">Update tags</p>
    <n-select
      v-bind="$attrs"
      v-model:value="selectedLabels"
      clearable
      filterable
      multiple
      :loading="getLabels.fetching?.value"
      :options="labels"
      placeholder="Select labels"
      class="mb-3"
    />

    <n-input
      v-model:value="newLabelInput"
      placeholder="Create label"
      :loading="createLabel.fetching?.value"
      :status="createLabel.error.value ? 'error' : undefined"
      @keyup.enter="createNewLabel($event.target.value)"
    >
      <template #prefix>
        <n-icon><i-mdi-tag /></n-icon>
      </template>
    </n-input>
  </div>
</template>

<script setup>
import { message } from "@/lib/core/status"
import { useGetLabelsQuery, useCreateLabelMutation } from "@/lib/api"

const props = defineProps({
  modelValue: {
    type: Array,
    required: true,
    default: () => [],
  },
})
const emit = defineEmits(["update:modelValue"])

const selectedLabels = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
})

const getLabels = useGetLabelsQuery()
const labels = computed(() =>
  getLabels.data?.value?.labels.edges.map(({ node }) => {
    return {
      label: node.name,
      value: node.id,
    }
  })
)

const newLabelInput = ref("")
const createLabel = useCreateLabelMutation()
function createNewLabel(val) {
  // check if it already exists.
  for (const label of labels.value) {
    if (label.label === val) {
      // check if it's already selected.
      if (selectedLabels.value.filter((l) => l.label == val).length == 0) {
        selectedLabels.value.push(label.value)
      }
      newLabelInput.value = ""
      return
    }
  }

  createLabel.executeMutation({ input: { name: val } }).then((result) => {
    if (!result.error) {
      getLabels.executeQuery().then(() => {
        selectedLabels.value = [...(selectedLabels.value ?? []), result.data.createLabel.id]
      })
      newLabelInput.value = ""
      message.success("Created label")
    } else {
      message.error(result.error.toString())
    }
  })
}
</script>
