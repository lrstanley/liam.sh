<template>
  <div>
    <p class="text-center">Update tags</p>
    <LabelSelect ref="selectRef" v-model="selected" field="id" :suggest="props.suggest" class="mb-3" />

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
import { useCreateLabelMutation } from "@/lib/api"

const props = defineProps({
  modelValue: {
    type: Array,
    required: true,
    default: () => [],
  },
  suggest: {
    type: String,
    default: "",
  },
})
const emit = defineEmits(["update:modelValue"])

const selectRef = ref(null)

const selected = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
})

const newLabelInput = ref("")
const createLabel = useCreateLabelMutation()
function createNewLabel(val) {
  createLabel.executeMutation({ input: { name: val } }).then((result) => {
    if (!result.error) {
      selectRef.value.refetch().then(() => {
        selected.value = [...(selected.value ?? []), result.data.createLabel.id]
      })
      newLabelInput.value = ""
      message.success("Created label")
    } else {
      message.error(result.error.toString())
    }
  })
}
</script>
