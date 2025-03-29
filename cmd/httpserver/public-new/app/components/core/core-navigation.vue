<script setup lang="ts">
const gh = useGithubUser()
const self = useSelf()
</script>

<template>
  <ul class="flex flex-wrap justify-center md:justify-start gap-x-2">
    <li v-for="link in menuOptions" :key="link.label">
      <router-link v-if="link.to" :to="link.to" active-class="active">
        {{ link.functionAlias || link.label?.toLowerCase() }}
      </router-link>
    </li>
    <li><a :href="gh?.html_url">github</a></li>
    <li><a href="/chat" target="_blank">discord</a></li>
    <li v-if="self">
      <router-link to="/admin">sudo</router-link>
    </li>
  </ul>
</template>

<style scoped>
@reference "@/assets/css/main.css";

ul :deep(a) {
  position: relative;
  display: inline-flex;
  padding-left: 5ch;
  padding-right: 4ch;
  @apply py-3 !text-(--ui-color-secondary-400) transition duration-100 text-sm;
}

ul :deep(a):not(.active):hover {
  padding-right: 4ch;
  @apply !text-(--ui-color-secondary-500);
}

ul :deep(a.active) {
  padding-right: 7ch;
  @apply !text-(--ui-color-secondary-500);
}

ul :deep(a)::before {
  content: "func";
  @apply absolute left-0 text-(--ui-color-error-400);
}

ul :deep(a):not(.active)::after {
  content: "()";
  @apply absolute right-[2ch] text-(--ui-color-warning-300) transition duration-100;
}

ul :deep(a):not(.active):hover::after {
  content: "(go)";
  @apply absolute right-0 text-(--ui-text-muted);
}

ul :deep(a.active)::after {
  content: "(ctx)";
  @apply absolute right-[2ch] text-(--ui-color-error-300) transition duration-100;
}

ul :deep(a.active):hover::after {
  content: "(ctx)";
  @apply absolute text-(--ui-color-error-400);
}
</style>
