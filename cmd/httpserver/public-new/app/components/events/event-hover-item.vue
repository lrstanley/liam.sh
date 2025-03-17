<script setup lang="ts">
import { UPopover } from "#components"

const {
  value = "",
  href = "",
  placement = "right",
} = defineProps<{
  value?: string
  href?: string
  placement?: Exclude<InstanceType<typeof UPopover>["$props"]["content"], undefined>["side"]
}>()
</script>

<template>
  <UPopover mode="hover" :open-delay="150" :content="{ side: placement }" :ui="{ content: 'px-1' }">
    <div v-if="href?.length > 0" class="align-middle" v-bind="$attrs">
      <EventLink :href="href">
        <slot name="icon" />
        <slot name="value">{{ value }}</slot>
      </EventLink>
    </div>
    <div v-else class="align-middle max-w-[24ch]" v-bind="$attrs">
      <slot name="icon" />
      <slot name="value">{{ value }}</slot>
    </div>

    <template #content>
      <div class="p-[2px] text-wrap max-w-96">
        <slot />
      </div>
    </template>
  </UPopover>
</template>
