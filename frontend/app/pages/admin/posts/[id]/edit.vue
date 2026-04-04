<script setup lang="ts">
import type { SchemaPostRead, SchemaPostUpdate } from '#open-fetch-schemas/api'

definePageMeta({
  title: "Edit Post",
  layout: "admin",
  validate: async (route) => {
    return typeof route.params.id === 'string' && /^\d+$/.test(route.params.id)
  }
})

const { $api } = useNuxtApp()
const toast = useToast()
const postID = +(useRoute("admin-posts-id-edit").params.id as string)
const labelData = ref<number[]>([])
const post = ref<SchemaPostRead>()
const error = ref<string>()
const loading = ref(false)

async function fetchPost() {
  loading.value = true
  $api('/posts/{postID}', { path: { postID: postID } })
    .then((data) => {
      post.value = data
      labelData.value = data.edges.labels?.map((label) => label.id) ?? []
    })
    .catch((e) => (error.value = "message" in e ? e.message : e.error))
    .finally(() => (loading.value = false))
}

async function invokeUpdate() {
  error.value = undefined

  const data: SchemaPostUpdate = {
    slug: post.value?.slug,
    title: post.value?.title,
    content: post.value?.content,
    published_at: post.value?.published_at,
    public: post.value?.public ?? false,
    labels: labelData.value,
  }

  loading.value = true
  $api('/posts/{postID}', { method: 'PATCH', path: { postID: postID }, body: data })
    .then(() => {
      toast.add({
        title: "Post updated successfully",
        color: "success",
        duration: 2000,
      })
    })
    .catch((e) => (error.value = e.message))
    .finally(() => (loading.value = false))
}

await fetchPost()
</script>

<template>
  <PostCreateEdit v-if="post" v-model="post" v-model:labels="labelData" :create="false" @save="invokeUpdate"
    :loading="loading" :error="error" />
</template>
