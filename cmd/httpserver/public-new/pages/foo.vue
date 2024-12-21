<script setup lang="ts">
const gh = useGithubUser()
</script>
<template>
  <h1>Foo Example Page: {{ gh?.login }}</h1>
  <n-button type="primary">Primary</n-button>

  <NuxtLink href="/">Home</NuxtLink>
</template>
