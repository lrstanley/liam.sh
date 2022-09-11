<template>
  <n-config-provider :theme="darkTheme" abstract preflight-style-disabled>
    <router-view />
  </n-config-provider>
</template>

<script setup lang="ts">
import { darkTheme } from "naive-ui"
</script>
