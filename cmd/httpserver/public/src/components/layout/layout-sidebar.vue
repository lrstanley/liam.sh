<template>
  <LayoutDefault :error="props.error" :loading="props.loading">
    <div class="grid gap-5 md:gap-10 lg:gap-[6rem] mt-8">
      <div v-bind="$attrs">
        <slot />
      </div>

      <div>
        <div
          class="flex flex-col flex-auto gap-4"
          :class="{ 'md:sticky md:left-0 md:top-10': props.affix }"
        >
          <slot name="sidebar" />
        </div>
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup lang="ts">
const props = defineProps<{
  loading?: boolean
  error?: Error | string
  affix?: boolean
}>()
</script>

<style scoped>
.grid {
  @apply grid-cols-1 md:grid-cols-[1fr,240px];
}
</style>
