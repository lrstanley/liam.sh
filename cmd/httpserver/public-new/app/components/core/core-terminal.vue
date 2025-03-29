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
    <span class="inline-flex text-emerald-600">
      {{ gh?.name.split(" ")[0]?.toLowerCase() }}
      <span class="text-zinc-500">:</span>
      {{ path }}
      <span class="text-zinc-500">$</span>
      <span v-if="prefix" class="mr-4" />
      <span v-else class="mr-2" />
      {{ prefix }}
    </span>
    <span class="inline-flex w-auto">
      <span
        class="inline-flex overflow-hidden border-solid border-r-3 text-gradient bg-gradient-to-r from-blue-400 to-emerald-400 cursor whitespace-nowrap border-r-emerald-500"
      >
        {{ value }}
      </span>
    </span>
  </div>
</template>

<style scoped>
.cursor {
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
