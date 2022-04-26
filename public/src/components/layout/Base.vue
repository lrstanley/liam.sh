<template>
    <slot></slot>
</template>

<script setup>
import { useLoadingBar } from "naive-ui"

const loadingBar = useLoadingBar()
const { loading } = storeToRefs(useState())

loading ? loadingBar.start() : loadingBar.finish()
watch(loading, (isLoading) => {
    isLoading ? loadingBar.start() : loadingBar.finish()
})
</script>
