<template>
  <n-anchor
    v-if="element"
    :top="90"
    style="z-index: 1"
    :bound="100"
    ignore-gap
    class="invisible xl:visible fixed right-[10px] w-[250px]"
  >
    <component :is="item" v-for="(item, index) in links" :key="index" />
  </n-anchor>
</template>

<script setup lang="ts">
import { h } from "vue"
import { NAnchorLink } from "naive-ui"
import type { VNode } from "vue"

interface Descendant {
  href: string
  title: string
  type: Node["nodeName"]
  descendants: Descendant[]
}

const props = defineProps<{
  element: Node
}>()

function allDescendants(node: Node, descendants: Descendant[] = []): Descendant[] {
  // Collect all headers, and their descendants.
  node.childNodes.forEach((child) => {
    if (child instanceof HTMLHeadingElement) {
      descendants.push({
        href: `#${child.id}`,
        title: child.textContent,
        type: child.nodeName,
        descendants: allDescendants(child),
      })
    }
  })

  for (let i = 0; i < descendants.length; i++) {
    if (i == 0) {
      continue
    }

    // If the previous header is larger, then take the current header in the loop
    // and append it as a child of the previous header.
    if (descendants[i].type > descendants[i - 1].type) {
      descendants[i - 1].descendants = [...descendants[i - 1].descendants, descendants[i]]
      descendants.splice(i, 1)
      i--
    }
  }

  return descendants
}

// wraps all elements in a custom type, and if it has children, it will recursively
// wrap them as well.
function elementize<T>(elementType: T, elements: Descendant[]): VNode[] {
  if (!elements) return []

  return elements.map((el) => {
    return h(
      elementType,
      {
        href: el.href,
        title: el.title,
      },
      () => elementize(elementType, el.descendants)
    )
  })
}

const links = computed(() => {
  if (!props.element) {
    return []
  }

  return elementize(NAnchorLink, allDescendants(props.element))
})
</script>

<style scoped>
.n-anchor :deep(a) {
  white-space: pre-wrap !important;
  line-height: 1.2em;
  padding-bottom: 4px;
}
.n-anchor :deep(a)::before {
  content: "·";
  position: relative;
  right: 3px;
  font-weight: bold;
  font-size: 1.5em;
  top: 2px;
  @apply text-pink-500;
}
</style>
