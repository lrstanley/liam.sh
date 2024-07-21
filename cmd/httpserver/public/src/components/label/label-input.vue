<script setup lang="ts">
import { message } from "@/lib/core/status"
import { createLabel } from "@/lib/http/services.gen"
import type { LabelCreate } from "@/lib/http/types.gen"

const props = defineProps<{
  suggest?: string
}>()

const selected = defineModel<number[]>()
const selectRef = ref(null)
const newLabelInput = ref("")

const { mutate, isPending, error } = useMutation({
  mutationFn: (data: LabelCreate) => unwrapErrors(createLabel({ body: data })),
  onError: (error) => {
    message.error("error creating label: " + error.message)
  },
  onSuccess: (resp) => {
    message.success(`created label ${resp.name}`)
    newLabelInput.value = ""
    selectRef.value.refetch().then(() => {
      selected.value = [...(selected.value ?? []), resp.id]
    })
  },
})
</script>

<template>
  <div>
    <span class="text-emerald-500">Update tags</span>
    <LabelSelect ref="selectRef" v-model="selected" field="id" :suggest="props.suggest" class="mb-3" />

    <n-input
      v-model:value="newLabelInput"
      placeholder="Create label"
      :loading="isPending"
      :status="error ? 'error' : undefined"
      @keyup.enter="mutate({ name: newLabelInput })"
    >
      <template #prefix>
        <n-icon><i-mdi-tag /></n-icon>
      </template>
    </n-input>
  </div>
</template>
