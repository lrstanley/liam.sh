<template>
  <n-space size="small" inline v-bind="$attrs">
    <n-tag
      v-for="(name, key) in props.sorter.fields"
      :key="key"
      :type="props.sorter.refs.field.value == key ? 'success' : ''"
      class="cursor-pointer"
      @click="props.sorter.toggle(key)"
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

<script setup>
const props = defineProps({
  sorter: {
    type: Object,
    required: true,
  },
})
</script>

<style scoped></style>
