<script setup lang="ts">
import { ConfirmModal, UBadge } from "#components"
import type { TableColumn } from "#ui/types"

definePageMeta({
  title: "Labels",
  layout: "admin",
})

const toast = useToast()

const {
  data: labels,
  error,
  status,
  refresh: refreshLabels,
} = await getLabelsCount({ composable: "useFetch" })
if (error.value) throw error.value

const columns: TableColumn<LabelCount>[] = [
  { id: "name", header: "Name", accessorKey: "name" },
  { id: "githubrepository_count", header: "# of repos", accessorKey: "githubrepository_count" },
  { id: "post_count", header: "# of posts", accessorKey: "post_count" },
  { id: "total_count", header: "# of total associations", accessorKey: "total_count" },
  { id: "actions", header: "Actions" },
]

const overlay = useOverlay()
const labelToDelete = ref<LabelCount | null>(null)
const deleteLoading = ref(false)
const modalDelete = overlay.create(ConfirmModal)

async function promptDeleteLabel(label: LabelCount) {
  deleteLoading.value = true
  labelToDelete.value = label
  modalDelete.patch({
    title: `Delete Label ${label.name}?`,
    description: `Are you sure you want to delete "${label.name}"?`,
    confirmText: "Delete",
    color: "error",
  })
  const result = await modalDelete.open()
  if (!result) return
  await deleteLabel({
    composable: "$fetch",
    path: {
      labelID: label.id,
    },
  })
  deleteLoading.value = false
  labelToDelete.value = null
  toast.add({
    title: "Label deleted",
    description: `"${label.name}" has been deleted.`,
    color: "success",
    duration: 2000,
  })
  refreshLabels()
}
</script>

<template>
  <div class="flex flex-col">
    <UCard variant="subtle">
      <UTable
        v-if="labels"
        :data="labels"
        :columns="columns"
        :loading="status === 'pending'"
        loading-color="primary"
        loading-animation="carousel"
        class="shrink-0"
        :ui="{ td: 'p-2' }"
      >
        <template #actions-cell="{ row }">
          <div class="flex flex-row gap-2">
            <UButton
              size="xs"
              variant="subtle"
              color="error"
              icon="mdi:delete"
              @click="promptDeleteLabel(row.original)"
            >
              Delete
            </UButton>
          </div>
        </template>
      </UTable>
    </UCard>
  </div>
</template>
