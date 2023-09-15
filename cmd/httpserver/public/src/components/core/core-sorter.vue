<script setup lang="ts">
import type { Sorter } from "@/lib/util"

const props = defineProps<{
  sorter: Sorter
}>()
</script>

<template>
  <n-space size="small" inline v-bind="$attrs">
    <n-tag
      v-for="(name, key) in props.sorter.fields"
      :key="key"
      :type="props.sorter.refs.field.value == key ? 'success' : 'default'"
      class="cursor-pointer"
      @click="props.sorter.toggle(key.toString())"
    >
      <template #avatar>
        <n-icon v-if="props.sorter.refs.field.value == key" class="">
          <i-mdi-arrow-up v-if="props.sorter.refs.direction.value == 'asc'" />
          <i-mdi-arrow-down v-else />
        </n-icon>
      </template>

      {{ name }}
    </n-tag>
  </n-space>
</template>

<style scoped></style>
