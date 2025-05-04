<script setup lang="ts"></script>

<template>
  <div :attr="$attrs" class="grid md:grid-cols-12 gap-5 md:gap-10 lg:gap-[4rem]">
    <div class="md:col-span-9 order-last lg:order-first overflow-x-auto overflow-y-hidden p-1">
      <slot />
    </div>
    <div
      class="md:col-span-3 md:sticky md:top-10 flex flex-col gap-4 text-center md:text-left max-h-screen p-1"
    >
      <slot name="sidebar" />
    </div>
  </div>
</template>
