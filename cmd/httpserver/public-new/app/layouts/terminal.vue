<script setup lang="ts">
const route = useRoute()
const error = ref<Error | null>(null)

onErrorCaptured((err) => {
  error.value = err
  return false
})

watch(
  () => route.path,
  () => {
    error.value = null
  }
)
</script>

<template>
  <div class="flex items-stretch justify-center flex-auto size-full lg:items-center">
    <div
      class="flex size-full shrink grow-0 basis-auto flex-col items-stretch pt-[15px] md:items-center lg:max-h-[28rem] lg:max-w-3xl"
    >
      <CoreNavigation />
      <CoreTerminal
        class="mb-4 flex flex-auto justify-center text-[38px] md:text-[45px]"
        path="~"
        prefix="#"
        value="Hello!"
      />

      <slot />
    </div>
  </div>
</template>
