<script setup lang="ts">
const {
  path = "~",
  prefix = "#",
  value,
} = defineProps<{
  path?: string
  prefix?: string
  value: string
}>()

const gh = useGithubUser()
</script>

<template>
  <div v-bind="$attrs">
    <span class="inline-flex mr-[10px] text-emerald-600">
      {{ gh?.name.split(" ")[0]?.toLowerCase() }}
      <span class="text-zinc-500">:</span>
      {{ path }}
      <span class="text-zinc-500">$</span>
      <span v-if="prefix" class="mr-4" />
      <span v-else class="mr-2" />
      {{ prefix }}
    </span>
    <span class="cursor-wrap">
      <span class="text-gradient bg-gradient-to-r from-blue-400 to-emerald-400 cursor">
        {{ value }}
      </span>
    </span>
  </div>
</template>

<style scoped>
@reference "@/assets/css/main.css";

.cursor-wrap {
  @apply w-auto inline-flex;
}
.cursor {
  @apply whitespace-nowrap overflow-hidden inline-flex border-r-4 border-r-emerald-500 border-solid;
  animation: typing 1s steps(40, end), blink-caret 0.75s step-end infinite;
}

@keyframes typing {
  from {
    width: 0;
  }
  to {
    width: 100%;
  }
}

@keyframes blink-caret {
  from,
  to {
    border-color: transparent;
  }
  50% {
    border-right-color: var(--color-emerald-700);
  }
}
</style>
