<template>
  <LayoutDefault :error="props.error" :loading="props.loading">
    <div class="grid gap-5 md:gap-10 lg:gap-[6rem] mt-8">
      <div v-bind="$attrs">
        <slot />
      </div>

      <div class="hidden" :class="{ 'md:inline-flex flex-col flex-wrap gap-6': affix }">
        <n-affix
          :top="props.affixTop ?? props.affixTrigger ?? 100"
          :trigger-top="props.affixTrigger ?? props.affixTop ?? 100"
          class="flex flex-col flex-wrap gap-6 max-w-[240px]"
        >
          <slot name="sidebar" />
        </n-affix>
      </div>
      <div :class="{ 'inline-flex md:hidden flex-col flex-wrap gap-6': affix }">
        <slot name="sidebar" />
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup lang="ts">
const props = defineProps<{
  loading?: boolean
  error?: Error | string
  affix?: boolean
  affixTop?: number
  affixTrigger?: number
}>()
</script>

<style scoped>
.grid {
  @apply grid-cols-1 md:grid-cols-[1fr,240px];
}
</style>
