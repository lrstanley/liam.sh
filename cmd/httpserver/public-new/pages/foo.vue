<script setup lang="ts">
import { getGithubUser } from "~/utils/http/sdk.gen"

const gh = useGithubUser()
const test = await getGithubUser({ composable: "$fetch" })
</script>
<template>
  <h1>Foo Example Page: {{ test.name }}</h1>
  <n-button type="primary">Primary</n-button>

  <NuxtLink href="/">Home</NuxtLink>
</template>
