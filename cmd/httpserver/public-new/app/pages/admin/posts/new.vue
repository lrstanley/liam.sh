<script setup lang="ts">
definePageMeta({
  title: "New Post",
  layout: "admin",
})

const toast = useToast()
const router = useRouter()

const labelData = ref<number[]>([])
const createPostData = ref<PostCreate>({ published_at: new Date().toISOString() } as PostCreate)
const error = ref<string>()
const loading = ref(false)

async function invokeCreate() {
  error.value = undefined

  const data: PostCreate = {
    ...createPostData.value,
    labels: labelData.value,
  }

  createPost({ composable: "$fetch", body: data })
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
  <PostCreateEdit
    v-model="createPostData"
    v-model:labels="labelData"
    :create="true"
    @save="invokeCreate"
    :loading="loading"
    :error="error"
  />
</template>
