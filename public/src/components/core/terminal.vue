<template>
  <div v-bind="$attrs">
    <span class="inline-flex mr-10px text-green-600">
      {{ state.base.githubUser.name.split(" ")[0].toLowerCase() }}
      <span class="text-gray-500">:</span>
      {{ props.path }}
      <span class="text-gray-500">$</span>
      <span v-if="props.prefix" class="mr-4" />
      <span v-else class="mr-2" />
      {{ props.prefix }}
    </span>
    <span class="text-gradient from-teal-500 to-violet-600">
      {{ props.value }}
    </span>
    <span class="text-size-0.9em text-gray-600 animate-pulse">▌</span>
  </div>
</template>

<script setup>
const props = defineProps({
  path: {
    type: String,
    default: "~",
  },
  prefix: {
    type: String,
    default: "#",
  },
  value: {
    type: String,
    default: "",
  },
})

const state = useState()
</script>
