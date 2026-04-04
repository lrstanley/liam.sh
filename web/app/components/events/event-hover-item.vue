<script setup lang="ts">
import { UPopover } from "#components"
import type { ComponentProps } from "vue-component-type-helpers"

const {
  value = "",
  href = "",
  placement = "right",
} = defineProps<{
  value?: string
  href?: string
  placement?: Exclude<ComponentProps<typeof UPopover>["content"], undefined>["side"]
}>()

const open = ref(false)
const target = useTemplateRef<HTMLDivElement>("target")

// UPopover doesn't do well when the mouse stays in its original place, but the user scrolls the element
// below the mouse, so it assumes it is still hovered. this should force it to close when the scrollable-parent
// is scrolled.
const { isScrolling } = getComputedScrollParent(target)

watch(isScrolling, () => open.value && (open.value = false))
</script>

<template>
  <UPopover v-model:open="open" mode="hover" :open-delay="100" :close-delay="0" :content="{ side: placement }"
    :ui="{ content: 'px-1' }">
    <div ref="target">
      <span v-if="href?.length > 0" class="align-middle" v-bind="$attrs">
        <EventLink :href="href">
          <slot name="icon" />
          <slot name="value">{{ value }}</slot>
        </EventLink>
      </span>
      <span v-else class="align-middle max-w-[24ch]" v-bind="$attrs">
        <slot name="icon" />
        <slot name="value">{{ value }}</slot>
      </span>
    </div>

    <template #content>
      <div class="p-[2px] text-wrap max-w-96">
        <slot />
      </div>
    </template>
  </UPopover>
</template>
