<script setup lang="ts">
import { useScrollspy } from "#imports"
const props = defineProps<{
  links: HeadingTree[]
  element: HTMLElement | null
}>()

const { activeHeadings, updateHeadings } = useScrollspy()

watchEffect(() => {
  if (!props.element) return
  updateHeadings([
    ...props.element.querySelectorAll("h1"),
    ...props.element.querySelectorAll("h2"),
    ...props.element.querySelectorAll("h3"),
    ...props.element.querySelectorAll("h4"),
  ])
})

// ul > li > a (recursively generate)
function generate(links: HeadingTree[], depth: number): VNode[] {
  return links.map((link) => {
    return h(
      "li",
      {
        class: ["flex flex-col gap-2 text-xs/3", depth > 0 ? "ml-4 list-none" : "list-none"].join(" "),
      },
      [
        h(
          "a",
          {
            href: `#${link.id}`,
            class: [
              "transition-all duration-100 hover:text-(--ui-color-secondary-400)",
              activeHeadings.value.includes(link.id)
                ? "text-(--ui-color-secondary-400)"
                : "text-(--ui-text)",
            ].join(" "),
          },
          link.title
        ),
        generate(link.children, depth + 1),
      ]
    )
  })
}

const elements = computed(() => generate(props.links, 0))
</script>

<template>
  <div v-show="links.length" v-bind="$attrs" class="text-sm">
    <div class="mb-1 font-bold text-(--ui-primary)">Table of Contents</div>

    <div class="text-left">
      <component :is="item" v-for="(item, index) in elements" :key="index" />
    </div>
  </div>
</template>
