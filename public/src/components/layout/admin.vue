<template>
  <LayoutBase>
    <n-layout position="absolute" content-style="display: flex;flex: 1 1 auto;flex-direction: column">
      <n-layout-header bordered>
        <n-menu :options="menuOptions" mode="horizontal" class="mt-1" />
      </n-layout-header>
      <n-layout has-sider>
        <n-layout-sider
          v-model:collapsed="state.sidebarCollapsed"
          :collapsed-width="54"
          :width="270"
          bordered
          show-trigger
          collapse-mode="width"
        >
          <n-menu
            :value="$route.name"
            :root-indent="16"
            :indent="12"
            :collapsed-width="54"
            :collapsed-icon-size="22"
            :options="adminSidebarOptions"
          />
        </n-layout-sider>
        <!-- content-style="padding-top: 30px;padding-bottom: 45px" -->
        <n-layout embedded content-style="height: 100%">
          <CoreBreadcrumbs class="pt-2 pl-3" />

          <slot><router-view /></slot>
        </n-layout>
      </n-layout>
    </n-layout>
  </LayoutBase>
</template>

<script setup>
import { menuOptions, adminSidebarOptions } from "@/lib/navigation.js"

const state = useState()
</script>
