<script setup lang="ts">
import { darkTheme } from "naive-ui"
import { VueQueryDevtools } from "@tanstack/vue-query-devtools"
</script>

<template>
  <n-config-provider :theme="darkTheme" abstract preflight-style-disabled>
    <router-view />
  </n-config-provider>
  <VueQueryDevtools />
</template>
