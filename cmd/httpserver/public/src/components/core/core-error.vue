<script setup lang="ts">
import type {
  ErrorBadRequest,
  ErrorForbidden,
  ErrorInternalServerError,
  ErrorTooManyRequests,
  ErrorUnauthorized,
} from "@/lib/http/types.gen"

const props = defineProps<{
  error:
    | Error
    | (ErrorBadRequest &
        ErrorUnauthorized &
        ErrorForbidden &
        ErrorTooManyRequests &
        ErrorInternalServerError)
}>()
</script>

<template>
  <n-alert v-if="props.error" v-motion-fade title="An error occurred" type="error" class="m-2 md:m-6">
    {{
      props.error instanceof Error
        ? props.error.message
        : props.error.error + " (" + props.error.code + ")"
    }}
  </n-alert>
</template>
