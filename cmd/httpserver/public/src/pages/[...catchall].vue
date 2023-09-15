<script setup lang="ts">
import type { ResultProps } from "naive-ui"

definePage({
  meta: {
    title: "Not Found",
    layout: "bare",
  },
})

const route = useRoute("/[...catchall]")
const source = computed(() => route.params.catchall)
const props = defineProps<{
  error?: Error | null
}>()

const errorCode = ref<ResultProps["status"]>()
const errorTitle = ref<string>("")
const supported = ["info", "success", "warning", "error", "404", "403", "500", "418"]

watch(
  source,
  () => {
    if (!source.value) return

    if (supported.includes(source.value)) {
      errorCode.value = source.value as ResultProps["status"]
    } else if (source.value.match(/^[45][0-9]+$/)) {
      errorCode.value = "error"
    } else if (source.value == "CombinedError") {
      errorCode.value = "error"
      errorTitle.value = "query error"
    } else {
      errorCode.value = "404"
      errorTitle.value = "not found"
    }

    if (errorTitle.value == "") {
      errorTitle.value = errorCode.value
    }
  },
  { immediate: true }
)
</script>

<template>
  <div class="flex flex-col justify-center flex-auto">
    <n-result
      :status="errorCode"
      :title="'Error code: ' + errorTitle"
      :description="props.error ? props.error.toString() : 'You know life is always ridiculous.'"
    >
      <template #footer>
        <n-button-group>
          <n-button @click="$router.back()">
            <n-icon class="mr-1"><i-mdi-undo-variant /></n-icon>
            Go back
          </n-button>
          <n-button @click="$router.push('/')">
            <n-icon class="mr-1"><i-mdi-home /></n-icon>
            Home
          </n-button>
        </n-button-group>
      </template>
    </n-result>
  </div>
</template>
