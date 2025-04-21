<script setup lang="ts">
definePageMeta({
  title: "Edit Post",
  layout: "admin",
})

const toast = useToast()

const postID = +useRoute("admin-posts-id-edit").params.id
const labelData = ref<number[]>([])
const post = ref<PostRead>()
const error = ref<string>()
const loading = ref(false)

// TODO: can this use useFetch without breaking everything?
async function fetchPost() {
  loading.value = true
  getPost({ composable: "$fetch", path: { postID: postID } })
    .then((data) => {
      post.value = data
      labelData.value = data.edges.labels?.map((label) => label.id) ?? []
    })
    .catch((e) => (error.value = "message" in e ? e.message : e.error))
    .finally(() => (loading.value = false))
}

async function invokeUpdate() {
  error.value = undefined

  const data: PostUpdate = {
    slug: post.value?.slug,
    title: post.value?.title,
    content: post.value?.content,
    published_at: post.value?.published_at,
    public: post.value?.public,
    labels: labelData.value,
  }

  loading.value = true
  updatePost({ composable: "$fetch", path: { postID: postID }, body: data })
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
  <PostCreateEdit
    v-if="post"
    v-model="post"
    v-model:labels="labelData"
    :create="false"
    @save="invokeUpdate"
    :loading="loading"
    :error="error"
  />
</template>
