<template>
  <n-layout position="absolute" content-style="display: flex;flex: 1 1 auto;flex-direction: column">
    <n-layout has-sider>
      <n-layout-sider
        v-model:collapsed="state.sidebarCollapsed"
        :collapsed-width="54"
        :width="220"
        bordered
        show-trigger
        collapse-mode="width"
      >
        <n-menu
          :options="adminSidebarOptions"
          :value="$route.name.toString()"
          :root-indent="16"
          :indent="12"
          :collapsed-width="54"
          :collapsed-icon-size="22"
        />
      </n-layout-sider>
      <n-layout
        embedded
        content-style="max-height: 100%;display: flex;flex: 1 1 auto;flex-direction: column"
      >
        <CoreBreadcrumbs class="hidden pt-2 pl-3 md:block" />

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
          <slot><router-view /></slot>
        </template>
      </n-layout>
    </n-layout>
  </n-layout>
</template>

<script setup lang="ts">
import { adminSidebarOptions } from "@/lib/util"

const state = useState()
const props = defineProps<{
  loading?: boolean
  error?: Error | string
}>()
</script>
