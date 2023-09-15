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

        <router-view v-slot="{ Component, route }">
          <transition name="fade" mode="out-in" appear>
            <Suspense>
              <component :is="Component" :key="route.path" />

              <template #fallback>
                <n-spin class="flex-auto">
                  <template #description>Loading...</template>
                </n-spin>
              </template>
            </Suspense>
          </transition>
        </router-view>
      </n-layout>
    </n-layout>
  </n-layout>
</template>

<script setup lang="ts">
import { adminSidebarOptions } from "@/lib/core/navigation"

const state = useState()
</script>
