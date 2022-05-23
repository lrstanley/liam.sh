<template>
  <n-empty v-if="showEmpty && objects.length < 1" description="No items found matching filters" />
  <component
    :is="object.component"
    v-for="object in objects"
    v-else-if="objects.length > 0"
    :key="object.id"
    v-bind="$attrs"
  />
</template>

<script setup>
import { h } from "vue"
import ObjectPost from "@/components/object/object-post.vue"
import ObjectLabel from "@/components/object/object-label.vue"

const props = defineProps({
  value: {
    type: [Array, Object, null],
    required: true,
  },
  showEmpty: {
    type: Boolean,
    default: false,
  },
})

const objects = computed(() => {
  if (!props.value) {
    return []
  }

  const results = typeMapper(props.value)
  if (!Array.isArray(results)) {
    return [results]
  }

  return results
})

function typeMapper(o) {
  if (!o.__typename) {
    return []
  }

  if (Array.isArray(o)) {
    return o.map(typeMapper)
  }

  switch (o.__typename) {
    case "PostConnection":
    case "LabelConnection":
      return o.edges.map(typeMapper)

    case "PostEdge":
    case "LabelEdge":
      return typeMapper(o.node)

    case "Post":
      return { component: h(ObjectPost, { value: o }), object: o }

    case "Label":
      return { component: h(ObjectLabel, { value: o }), object: o }

    default:
      return []
  }
}
</script>
