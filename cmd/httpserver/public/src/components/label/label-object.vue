<template>
  <component
    :is="props.linkable ? 'router-link' : 'span'"
    :to="{ name: props.route, query: { [props.query]: label.name } }"
  >
    <n-tag v-bind="$attrs" class="cursor-pointer hover:bg-emerald-700">
      {{ label.name }}
    </n-tag>
  </component>
</template>

<script setup>
const props = defineProps({
  linkable: {
    type: Boolean,
    default: false,
  },
  value: {
    type: Object,
    required: true,
  },
  route: {
    type: String,
    default: "posts",
  },
  query: {
    type: String,
    default: "label",
  },
})

const label = ref(props.value)
</script>
