<script setup lang="ts">
const route = useRoute()
const error = ref<Error | null>(null)

const githubUser = useGithubUser()

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
  <div class="z-[1] relative flex flex-auto flex-col">
    <div
      class="sm:container flex flex-auto flex-col pt-[15px] lg:pt-[70px] pb-[60px] max-w-[100vw] px-4 md:px-0 xl:px-[200px] sm:mx-auto"
    >
      <CoreNavigation />

      <main class="size-full">
        <slot />
      </main>
    </div>

    <motion
      as="span"
      :initial="{ opacity: 0 }"
      :animate="{ opacity: 1 }"
      :transition="{ delay: 2 }"
      class="p-2 text-center"
    >
      Made with
      <UIcon name="mdi:heart" class="align-middle text-(--ui-color-primary-600)" />
      by
      <a :href="githubUser?.html_url" target="_blank">{{ githubUser?.login }}</a>
    </motion>
  </div>
</template>
