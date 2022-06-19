<template>
  <ul
    class="flex flex-wrap justify-center md:justify-start md:gap-x-2 md:gap-y-4 px-4 md:px-0 pb-5 md:pb-0"
  >
    <li v-for="link in menuOptions" :key="link.name">
      <router-link :to="link.to" active-class="active">
        {{ link.alias }}
      </router-link>
    </li>
    <li><a :href="state.base.githubUser.htmlurl">github</a></li>
    <li><a href="/chat" target="_blank">discord</a></li>
    <!-- <li>
      <router-link :to="{ name: 'admin' }" class="filter blur-1">sudo</router-link>
    </li> -->
  </ul>
</template>

<script setup>
import { menuOptions } from "@/lib/navigation"

const state = useState()
</script>

<style scoped>
ul :deep(a) {
  position: relative;
  display: inline-flex;
  padding-left: 5ch;
  padding-right: 4ch;
  @apply !text-violet-400 transition duration-100;
}

ul :deep(a):hover,
ul :deep(a).active {
  padding-right: 4ch;
  @apply !text-violet-500;
}

ul :deep(a)::before {
  content: "func";
  @apply absolute left-0 text-red-400;
}

ul :deep(a)::after {
  content: "()";
  @apply absolute right-2ch text-yellow-300 transition duration-100;
}

ul :deep(a):hover::after {
  content: "(ctx)";
  @apply absolute right-0 text-red-400;
}

ul :deep(a):hover::after {
  content: "(go)";
  @apply absolute right-0 text-gray-400;
}
</style>
