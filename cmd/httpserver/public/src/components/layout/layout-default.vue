<template>
  <!-- TODO: less flex? -->
  <div class="z-1 relative flex flex-auto flex-col">
    <div
      class="sm:container flex flex-auto flex-col mt-30px lg:mt-70px mb-60px max-w-100vw px-4 md:px-0 xl:px-200px sm:mx-auto"
    >
      <CoreNavigation />

      <n-alert
        v-if="props.error"
        v-motion-fade
        title="An error occurred"
        type="error"
        class="m-2 md:m-6"
      >
        {{ props.error }}
      </n-alert>
      <template v-else-if="props.loading">
        <n-spin class="flex-auto">
          <template #description> Loading... </template>
        </n-spin>
      </template>
      <template v-else>
        <slot><router-view /></slot>
      </template>
    </div>

    <span class="px-2 py-2 text-center">
      Made with
      <n-icon class="text-green-600 align-middle">
        <i-mdi-heart />
      </n-icon>
      by <a :href="state.base.githubUser.htmlurl" target="_blank">{{ state.base.githubUser.login }}</a>
    </span>
  </div>
</template>

<script setup>
const props = defineProps({
  loading: {
    type: Boolean,
    default: false,
  },
  error: {
    type: Error,
    default: null,
  },
})

const state = useState()
</script>
