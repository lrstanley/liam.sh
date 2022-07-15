<template>
  <n-empty v-if="showEmpty && objects.length < 1" description="No items found matching filters" />
  <TransitionGroup v-else-if="objects.length > 0" appear name="fade">
    <div
      v-for="(object, i) in objects"
      :key="object.object.id"
      :style="{ '--i': i, '--total': objects.length }"
    >
      <component :is="object.component" v-bind="$attrs" />
      <n-divider v-if="divider && i != objects.length - 1" />
    </div>
  </TransitionGroup>
</template>

<script setup>
import { h } from "vue"
import PostObject from "@/components/post/post-object.vue"
import LabelObject from "@/components/label/label-object.vue"
import RepoObject from "@/components/repo/repo-object.vue"

const props = defineProps({
  value: {
    type: [Array, Object, null],
    required: true,
  },
  showEmpty: {
    type: Boolean,
    default: false,
  },
  divider: {
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
  if (Array.isArray(o)) {
    return o.map(typeMapper)
  }

  if (!o.__typename) {
    return []
  }

  if (o.__typename.endsWith("Connection")) {
    return o.edges.map(typeMapper)
  }

  if (o.__typename.endsWith("Edge")) {
    return typeMapper(o.node)
  }

  switch (o.__typename) {
    case "Post":
      return { component: h(PostObject, { value: o }), object: o }

    case "Label":
      return { component: h(LabelObject, { value: o }), object: o }

    case "GithubRepository":
      return { component: h(RepoObject, { value: o }), object: o }

    default:
      return []
  }
}
</script>

<style scoped>
.fade-move,
.fade-enter-active,
.fade-leave-active {
  transition: all 0.1s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
