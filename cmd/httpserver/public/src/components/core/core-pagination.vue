<template>
  <n-button-group>
    <n-button ghost :disabled="!page.hasPreviousPage" @click="cursor = 'b.' + page.startCursor">
      <template #icon>
        <n-icon><i-mdi-chevron-left /></n-icon>
      </template>
    </n-button>
    <n-button ghost :disabled="!page.hasNextPage" @click="cursor = 'a.' + page.endCursor">
      <template #icon>
        <n-icon><i-mdi-chevron-right /></n-icon>
      </template>
    </n-button>
  </n-button-group>
</template>

<script setup>
const props = defineProps({
  modelValue: {
    type: [String, null],
    required: true,
  },
  pageInfo: {
    type: [Object, null],
    required: false,
    default: null,
  },
})

const emit = defineEmits(["update:modelValue"])

const cursor = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
})

const page = computed(() => props.pageInfo || {})
</script>
