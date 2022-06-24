<template>
  <n-popover trigger="hover" style="padding: 2px 6px" :to="false">
    <template #trigger>
      <span v-if="props.href?.length > 0" class="align-middle" v-bind="$attrs">
        <EventLink :href="props.href">
          <slot name="icon"></slot>
          <slot name="value">{{ props.value }}</slot>
        </EventLink>
      </span>
      <span v-else class="align-middle max-w-24ch" v-bind="$attrs">
        <slot name="icon"></slot>
        <slot name="value">{{ props.value }}</slot>
      </span>
    </template>

    <slot></slot>
  </n-popover>
</template>

<script setup>
const props = defineProps({
  value: {
    type: String,
    required: false,
    default: "",
  },
  href: {
    type: String,
    required: false,
    default: "",
  },
})
</script>
