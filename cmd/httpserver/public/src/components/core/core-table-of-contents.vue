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

<script setup>
import { h } from "vue"
import { NAnchorLink } from "naive-ui"

const props = defineProps({
  element: {
    type: [HTMLElement, null],
    required: true,
  },
})

const validHeaders = ["H1", "H2", "H3", "H4", "H5", "H6"]

function allDescendants(node, descendants = []) {
  // Collect all headers, and their descendants.
  node.childNodes.forEach((child) => {
    if (validHeaders.includes(child.nodeName)) {
      descendants.push({
        href: `#${child.id}`,
        title: child.textContent,
        type: child.nodeName,
        children: allDescendants(child),
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
      descendants[i - 1].children = [...descendants[i - 1].children, descendants[i]]
      descendants.splice(i, 1)
      i--
    }
  }

  return descendants
}

// wraps all elements in a custom type, and if it has children, it will recursively
// wrap them as well.
function elementize(elementType, elements) {
  return elements.map((el) => {
    return h(
      elementType,
      {
        href: el.href,
        title: el.title,
      },
      () => elementize(elementType, el.children)
    )
  })
}

const links = computed({
  get() {
    if (!props.element) {
      return []
    }

    return elementize(NAnchorLink, allDescendants(props.element))
  },
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
