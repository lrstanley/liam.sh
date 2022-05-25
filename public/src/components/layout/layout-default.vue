<template>
  <LayoutBase>
    <!-- TODO: less flex? -->
    <div class="z-1 relative flex flex-auto flex-col">
      <div
        class="sm:container flex flex-auto flex-col mt-70px mb-45px max-w-100vw px-4 xl:px-200px sm:mx-auto"
      >
        <CoreNavigation />

        <n-alert
          v-if="props.error"
          v-motion-fade-visible
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

      <span class="flex flex-auto justify-center px-2 py-2">
        Made with
        <n-icon class="text-green-600 mx-1">
          <i-mdi-heart />
        </n-icon>
        by <a :href="state.base.githubUser.htmlurl" target="_blank">{{ state.base.githubUser.login }}</a>
      </span>
    </div>
  </LayoutBase>
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
