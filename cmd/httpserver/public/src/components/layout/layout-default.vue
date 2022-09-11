<template>
  <!-- TODO: less flex? -->
  <div class="z-[1] relative flex flex-auto flex-col">
    <div
      class="sm:container flex flex-auto flex-col pt-[15px] lg:pt-[70px] pb-[60px] max-w-[100vw] px-4 md:px-0 xl:px-[200px] sm:mx-auto"
    >
      <CoreNavigation />

      <template v-if="props.loading">
        <n-spin class="flex-auto">
          <template #description> Loading... </template>
        </n-spin>
      </template>
      <n-alert
        v-else-if="props.error"
        v-motion-fade
        title="An error occurred"
        type="error"
        class="m-2 md:m-6"
      >
        {{ props.error }}
      </n-alert>
      <template v-else>
        <slot><router-view /></slot>
      </template>
    </div>

    <span class="p-2 text-center">
      Made with
      <n-icon class="align-middle text-emerald-600">
        <i-mdi-heart />
      </n-icon>
      by <a :href="state.base.githubUser.htmlurl" target="_blank">{{ state.base.githubUser.login }}</a>
    </span>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  loading?: boolean
  error?: Error | string
}>()

const state = useState()
</script>
