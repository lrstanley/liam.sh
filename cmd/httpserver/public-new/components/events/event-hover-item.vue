<script setup lang="ts">
import { UPopover } from "#components"

const props = withDefaults(
  defineProps<{
    value?: string
    href?: string
    placement?: Exclude<InstanceType<typeof UPopover>["$props"]["popper"], undefined>["placement"]
  }>(),
  {
    value: "",
    href: "",
    placement: "right",
  }
)
</script>

<template>
  <UPopover
    mode="hover"
    :popper="{ placement: props.placement }"
    :ui="{ background: 'dark:bg-zinc-800', base: 'text-white', rounded: 'rounded' }"
  >
    <div v-if="props.href?.length > 0" class="align-middle" v-bind="$attrs">
      <EventLink :href="props.href">
        <slot name="icon" />
        <slot name="value">{{ props.value }}</slot>
      </EventLink>
    </div>
    <div v-else class="align-middle max-w-[24ch]" v-bind="$attrs">
      <slot name="icon" />
      <slot name="value">{{ props.value }}</slot>
    </div>

    <template #panel>
      <div class="p-[2px] text-wrap max-w-96">
        <slot />
      </div>
    </template>
  </UPopover>
</template>
