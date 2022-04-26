<template>
    <LayoutBase>
        <div class="flex flex-col h-full justify-center content-center">
            <n-result
                :status="errorCode"
                :title="'Error code: ' + (errorCode == 'error' ? 'unknown' : errorCode)"
                description="You know life is always ridiculous."
            >
                <template #footer>
                    <n-button-group>
                        <n-button @click="$router.back()">
                            <n-icon class="mr-1"><ArrowHookDownLeft24Filled /></n-icon>
                            Go back
                        </n-button>
                        <n-button @click="$router.push('/')">
                            <n-icon class="mr-1"><Home24Regular /></n-icon>
                            Home
                        </n-button>
                    </n-button-group>
                </template>
            </n-result>
        </div>
    </LayoutBase>
</template>

<script setup>
import { ArrowHookDownLeft24Filled, Home24Regular } from "@vicons/fluent"

const $route = useRoute()
const errorCode = ref("0")

onMounted(() => {
    if ($route.params.pathMatch.match(/^[0-9]+$/)) {
        errorCode.value = $route.params.pathMatch
    } else {
        errorCode.value = "404"
    }
})
</script>
