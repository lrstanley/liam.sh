<script setup lang="ts"></script>

<template>
  <div
    :attr="$attrs"
    class="grid grid-flow-row lg:grid-flow-col grid-cols-1 lg:grid-cols-[1fr,240px] gap-5 md:gap-10 lg:gap-[6rem]"
  >
    <div class="order-last lg:order-first"><slot /></div>
    <div class="flex flex-col flex-auto gap-4 first:lg:left-0 first:lg:top-10 first:lg:sticky">
      <slot name="sidebar" />
    </div>
  </div>
</template>

<style scoped></style>
