<script setup lang="ts">
const route = useRoute()
const error = ref<Error>(null)

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

      <router-view v-slot="{ Component, route: r }">
        <transition name="fade" mode="out-in" appear>
          <CoreError v-if="error" :error="error" />
          <Suspense v-else>
            <component :is="Component" :key="r.path" />

            <template #fallback>
              <n-spin class="flex flex-auto h-full gap-4 mx-auto place-content-center">
                <template #description>Loading...</template>
              </n-spin>
            </template>
          </Suspense>
        </transition>
      </router-view>
    </div>
  </div>
</template>
