<script setup lang="ts">
const state = useState()
</script>

<template>
  <n-breadcrumb>
    <n-breadcrumb-item v-for="(item, index) in state.history" :key="item.timestamp">
      <router-link :to="item.path">
        <n-tag
          :disabled="index == state.history.length - 1"
          size="small"
          class="cursor-pointer hover:bg-emerald-700"
        >
          {{ item.title }}
        </n-tag>
      </router-link>
    </n-breadcrumb-item>
  </n-breadcrumb>
</template>
