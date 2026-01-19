<script setup lang="ts">
import type { SchemaPostCreate } from '#open-fetch-schemas/api'
definePageMeta({
  title: "New Post",
  layout: "admin",
})

const { $api } = useNuxtApp()
const toast = useToast()
const router = useRouter()

const labelData = ref<number[]>([])
const createPostData = ref<SchemaPostCreate>({ published_at: new Date().toISOString(), public: false } as SchemaPostCreate)
const error = ref<string>()
const loading = ref(false)

async function invokeCreate() {
  error.value = undefined

  const data: SchemaPostCreate = {
    ...createPostData.value,
    labels: labelData.value,
  }

  $api('/posts', {
    method: 'POST',
    body: data,
  })
    .then((v) => {
      toast.add({
        title: "Post created successfully",
        color: "success",
        duration: 2000,
      })
      router.push({ name: "admin-posts-id-edit", params: { id: v.id } })
    })
    .catch((e) => (error.value = e.message))
}
</script>

<template>
  <PostCreateEdit v-model="createPostData" v-model:labels="labelData" :create="true" @save="invokeCreate"
    :loading="loading" :error="error" />
</template>
