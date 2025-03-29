<script setup lang="ts"></script>

<template>
  <div
    :attr="$attrs"
    class="grid grid-flow-row lg:grid-flow-col grid-cols-1 lg:grid-cols-[1fr,240px] gap-5 md:gap-10 lg:gap-[4rem]"
  >
    <div class="order-last lg:order-first"><slot /></div>
    <UPageAside class="lg:w-[240px] pe-0! ps-0! py-0!">
      <div class="flex flex-col gap-4 text-center md:text-left">
        <slot name="sidebar" />
      </div>
    </UPageAside>
  </div>
</template>

<style scoped></style>
