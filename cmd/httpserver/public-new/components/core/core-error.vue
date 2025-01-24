<script setup lang="ts">
import type {
  ErrorBadRequest,
  ErrorForbidden,
  ErrorInternalServerError,
  ErrorTooManyRequests,
  ErrorUnauthorized,
} from "@/utils/http/types.gen"

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
  <UAlert
    v-if="props.error"
    v-motion-fade
    icon="i-heroicons-exclamation-circle"
    title="An error occurred"
    :description="
      props.error instanceof Error
        ? props.error.message
        : props.error.error + ' (' + props.error.code + ')'
    "
    color="red"
    variant="subtle"
  />
</template>
