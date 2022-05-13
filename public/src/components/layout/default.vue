<template>
  <LayoutBase>
    <div class="z-1 relative flex flex-auto flex-col">
      <n-layout-header bordered class="absolute pt-1 h-45px z-10">
        <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions" />
      </n-layout-header>

      <n-layout embedded content-style="margin-top: 70px;margin-bottom: 45px">
        <div class="sm:container sm:mx-auto flex flex-auto">
          <n-card>
            <slot><router-view /></slot>
          </n-card>
        </div>
      </n-layout>

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

const activeKey = ref(null)
</script>
