<route lang="yaml">
meta:
  layout: admin
</route>

<template>
  <div>
    <n-page-header class="hidden px-5 mt-4 mb-8 md:block">
      <template #avatar>
        <n-icon :size="40"><i-mdi-pencil-outline /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit"> Create new post </a>
      </template>
    </n-page-header>

    <div class="p-4 sm:container sm:mx-auto lg:p-0">
      <PostCreateEdit create @update:post="createPost" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { message } from "@/lib/core/status"
import { useCreatePostMutation, type CreatePostInput } from "@/lib/api"

const router = useRouter()
const post = useCreatePostMutation()

function createPost(val: CreatePostInput) {
  post.executeMutation({ input: val }).then((result) => {
    if (!result.error) {
      message.success("Post created successfully")
      router.push({ name: "/admin/posts" })
    } else {
      message.error(result.error.toString())
    }
  })
}
</script>
