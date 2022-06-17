<template>
  <n-empty v-if="showEmpty && objects.length < 1" description="No items found matching filters" />
  <TransitionGroup v-else-if="objects.length > 0" appear>
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
import ObjectPost from "@/components/object/object-post.vue"
import ObjectLabel from "@/components/object/object-label.vue"
import ObjectRepo from "@/components/object/object-repo.vue"

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
      return { component: h(ObjectPost, { value: o }), object: o }

    case "Label":
      return { component: h(ObjectLabel, { value: o }), object: o }

    case "GithubRepository":
      return { component: h(ObjectRepo, { value: o }), object: o }

    default:
      return []
  }
}
</script>

<style scoped>
.v-move {
  transition: opacity 0.2s linear, transform 0.2s ease-in;
}

.v-enter-active {
  @apply transform transition-all duration-100;
  transition-delay: calc(min(0.1s, 0.1s * (var(--total) - var(--i))));
}

.v-leave-active {
  @apply transform transition-all duration-100 absolute max-w-90vw md:max-w-40vw;
}

.v-enter-from,
.v-leave-to {
  @apply opacity-0 -translate-x-50px;
}
</style>
