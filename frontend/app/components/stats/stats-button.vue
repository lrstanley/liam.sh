<template>
  <UBadge color="neutral" variant="outline" size="lg" v-bind="$attrs">
    <span class="text-gradient bg-linear-to-br from-pink-500 via-red-500 to-yellow-500">
      <slot />
    </span>
  </UBadge>
</template>
