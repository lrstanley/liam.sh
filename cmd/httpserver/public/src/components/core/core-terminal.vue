<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    path?: string
    prefix?: string
    value: string
  }>(),
  {
    path: "~",
    prefix: "#",
  }
)

const state = useState()
</script>

<template>
  <div v-bind="$attrs">
    <span class="inline-flex mr-[10px] text-emerald-600">
      {{ state.githubUser?.name.split(" ")[0].toLowerCase() }}
      <span class="text-zinc-500">:</span>
      {{ props.path }}
      <span class="text-zinc-500">$</span>
      <span v-if="props.prefix" class="mr-4" />
      <span v-else class="mr-2" />
      {{ props.prefix }}
    </span>
    <span class="cursor-wrap">
      <span class="text-gradient bg-gradient-to-r from-blue-400 to-emerald-400 cursor">
        {{ props.value }}
      </span>
    </span>
  </div>
</template>

<style scoped>
.cursor-wrap {
  @apply w-auto inline-flex;
}
.cursor {
  @apply whitespace-nowrap overflow-hidden inline-flex border-r-4 border-r-emerald-500 border-solid;
  animation:
    typing 1s steps(40, end),
    blink-caret 0.75s step-end infinite;
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
    @apply border-transparent;
  }
  50% {
    @apply border-r-emerald-700;
  }
}
</style>
