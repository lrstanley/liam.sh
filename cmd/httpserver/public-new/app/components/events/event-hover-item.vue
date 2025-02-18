<script setup lang="ts">
import { UPopover } from "#components"

const props = withDefaults(
  defineProps<{
    value?: string
    href?: string
    placement?: Exclude<InstanceType<typeof UPopover>["$props"]["content"], undefined>["side"]
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
    :open-delay="150"
    :content="{ side: props.placement }"
    :ui="{ content: 'px-1' }"
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

    <template #content>
      <div class="p-[2px] text-wrap max-w-96">
        <slot />
      </div>
    </template>
  </UPopover>
</template>
