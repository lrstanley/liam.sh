<script setup lang="ts">
import { menuOptions } from "@/lib/core/navigation"

const state = useState()
</script>

<template>
  <ul class="flex flex-wrap justify-center md:justify-start gap-x-2">
    <li v-for="link in menuOptions" :key="link.name">
      <router-link :to="link.to" active-class="active">
        {{ link.alias }}
      </router-link>
    </li>
    <li><a :href="state.base.githubUser.htmlurl">github</a></li>
    <li><a href="/chat" target="_blank">discord</a></li>
    <li v-if="state.base?.self">
      <router-link to="/admin/posts">sudo</router-link>
    </li>
  </ul>
</template>

<style scoped>
ul :deep(a) {
  position: relative;
  display: inline-flex;
  padding-left: 5ch;
  padding-right: 4ch;
  @apply py-3;
  @apply !text-violet-400 transition duration-100;
}

ul :deep(a):not(.active):hover {
  padding-right: 4ch;
  @apply !text-violet-500;
}

ul :deep(a.active) {
  padding-right: 7ch;
  @apply !text-violet-500;
}

ul :deep(a)::before {
  content: "func";
  @apply absolute left-0 text-red-400;
}

ul :deep(a):not(.active)::after {
  content: "()";
  @apply absolute right-[2ch] text-yellow-300 transition duration-100;
}

ul :deep(a):not(.active):hover::after {
  content: "(go)";
  @apply absolute right-0 text-zinc-400;
}

ul :deep(a.active)::after {
  content: "(ctx)";
  @apply absolute right-[2ch] text-red-300 transition duration-100;
}

ul :deep(a.active):hover::after {
  content: "(ctx)";
  @apply absolute text-red-400;
}
</style>
