<script setup lang="ts">
import { UButton } from "#components"
import type { ComponentProps } from "vue-component-type-helpers"

const {
  title,
  description,
  confirmText = "Confirm",
  color = "success",
} = defineProps<{
  title: string
  description: string
  confirmText: string
  color: ComponentProps<typeof UButton>["color"]
}>()

const emit = defineEmits<{ close: [boolean] }>()
</script>

<template>
  <UModal :close="{ onClick: () => emit('close', false) }" :title="title">
    <template #body>
      <p class="text-sm text-neutral-500">{{ description }}</p>
    </template>
    <template #footer>
      <div class="ml-auto"></div>
      <UButton :label="confirmText" :color="color" variant="subtle" @click="emit('close', true)" />
      <UButton color="neutral" variant="subtle" label="Cancel" @click="emit('close', false)" />
    </template>
  </UModal>
</template>
