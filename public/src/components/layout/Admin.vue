<template>
    <LayoutBase>
        <n-layout position="absolute" content-style="display: flex;flex: 1 1 auto;flex-direction: column">
            <n-layout-header bordered>
                <n-menu :options="menuOptions" mode="horizontal" class="mt-1" />
            </n-layout-header>
            <n-layout has-sider>
                <n-layout-sider bordered show-trigger collapse-mode="width" :collapsed-width="54" :width="270">
                    <n-menu
                        :value="$route.name"
                        :root-indent="16"
                        :indent="12"
                        :collapsed-width="54"
                        :collapsed-icon-size="22"
                        :options="adminSidebarOptions"
                    />
                </n-layout-sider>
                <n-layout embedded content-style="padding-top: 30px;padding-bottom: 45px">
                    <slot><router-view /></slot>
                </n-layout>
            </n-layout>
        </n-layout>
    </LayoutBase>
</template>

<script setup>
import { menuOptions, adminSidebarOptions } from "@/lib/navigation.js"
</script>
