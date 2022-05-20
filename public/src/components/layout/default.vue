<template>
  <LayoutBase>
    <div class="z-1 relative flex flex-auto flex-col">
      <n-layout-header bordered class="absolute pt-1 h-45px z-10">
        <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions" />
      </n-layout-header>

      <n-alert
        v-if="props.error"
        v-motion-fade-visible
        title="An error occurred"
        type="error"
        class="m-2 md:m-6"
      >
        {{ props.error }}
      </n-alert>
      <template v-else-if="props.loading">
        <n-spin class="flex-auto">
          <template #description> Loading... </template>
        </n-spin>
      </template>
      <template v-else>
        <div class="sm:container sm:mx-auto flex flex-auto flex-col mt-70px mb-45px">
          <slot><router-view /></slot>
        </div>
      </template>

      <n-layout-footer bordered>
        <span class="flex flex-auto justify-end px-2">
          Made with
          <n-icon class="text-green-600 mx-1">
            <i-mdi-heart />
          </n-icon>
          by Liam
        </span>
      </n-layout-footer>
    </div>
  </LayoutBase>
</template>

<script setup>
import { menuOptions } from "@/lib/navigation.js"

const props = defineProps({
  loading: {
    type: Boolean,
    default: false,
  },
  error: {
    type: Error,
    default: null,
  },
})

const activeKey = ref(null)
</script>
