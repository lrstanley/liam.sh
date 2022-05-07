<template>
    <LayoutBase>
        <div class="flex flex-auto flex-col justify-center">
            <n-result
                :status="errorCode"
                :title="'Error code: ' + $route.params.pathMatch"
                description="You know life is always ridiculous."
            >
                <template #footer>
                    <n-button-group>
                        <n-button @click="$router.back()">
                            <n-icon class="mr-1"><ArrowUndoOutline /></n-icon>
                            Go back
                        </n-button>
                        <n-button @click="$router.push('/')">
                            <n-icon class="mr-1"><Home /></n-icon>
                            Home
                        </n-button>
                    </n-button-group>
                </template>
            </n-result>
        </div>
    </LayoutBase>
</template>

<script setup>
const $route = useRoute()
const errorCode = ref("0")

const supported = ["info", "success", "warning", "error", "404", "403", "500", "418"]

onMounted(() => {
    if (supported.includes($route.params.pathMatch)) {
        errorCode.value = $route.params.pathMatch
    } else if ($route.params.pathMatch.match(/^[45][0-9]+$/)) {
        errorCode.value = "error"
    } else {
        errorCode.value = $route.params.pathMatch
    }
})
</script>
