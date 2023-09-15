<script setup lang="ts">
const state = useState()
const route = useRoute()
const error = ref<Error>(null)

onErrorCaptured((err: Error) => {
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

      <main class="w-full h-full">
        <CoreError v-if="error" :error="error" />
        <router-view v-else v-slot="{ Component, route: r }">
          <transition name="fade" mode="out-in" appear>
            <Suspense>
              <component :is="Component" :key="r.path" />

              <template #fallback>
                <n-spin class="flex flex-auto h-full gap-4 mx-auto place-content-center">
                  <template #description>Loading...</template>
                </n-spin>
              </template>
            </Suspense>
          </transition>
        </router-view>
      </main>
    </div>

    <span class="p-2 text-center">
      Made with
      <n-icon class="align-middle text-emerald-600">
        <i-mdi-heart />
      </n-icon>
      by
      <a :href="state.base.githubUser.htmlurl" target="_blank">{{ state.base.githubUser.login }}</a>
    </span>
  </div>
</template>
