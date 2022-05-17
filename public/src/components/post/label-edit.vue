<template>
  <div>
    <p class="text-center">Update tags</p>
    <n-select
      v-bind="$props"
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
import { useMessage } from "naive-ui"
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

const message = useMessage()
const newLabelInput = ref("")
const createLabel = useCreateLabelMutation()
function createNewLabel(val) {
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
