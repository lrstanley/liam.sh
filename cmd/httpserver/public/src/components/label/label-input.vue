<template>
  <div>
    <span class="text-emerald-500">Update tags</span>
    <LabelSelect ref="selectRef" v-model="selected" field="id" :suggest="props.suggest" class="mb-3" />

    <n-input
      v-model:value="newLabelInput"
      placeholder="Create label"
      :loading="createLabel.fetching?.value"
      :status="createLabel.error.value ? 'error' : undefined"
      @keyup.enter="createNewLabel(newLabelInput)"
    >
      <template #prefix>
        <n-icon><i-mdi-tag /></n-icon>
      </template>
    </n-input>
  </div>
</template>

<script setup lang="ts">
import { message } from "@/lib/core/status"
import { useCreateLabelMutation } from "@/lib/api"

const props = defineProps<{
  modelValue: string[]
  suggest?: string
}>()

const emit = defineEmits(["update:modelValue"])

const selectRef = ref(null)

const selected = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
})

const newLabelInput = ref("")
const createLabel = useCreateLabelMutation()

function createNewLabel(val: string) {
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
