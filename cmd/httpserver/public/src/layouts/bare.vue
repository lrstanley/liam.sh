<template>
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
</template>

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
