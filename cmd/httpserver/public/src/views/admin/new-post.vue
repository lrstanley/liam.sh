<template>
  <LayoutAdmin>
    <n-page-header class="px-5 mt-4 mb-8 hidden md:block">
      <template #avatar>
        <n-icon :size="40"><i-mdi-pencil-outline /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit"> Create new post </a>
      </template>
    </n-page-header>

    <div class="sm:container sm:mx-auto p-4 lg:p-0">
      <PostCreateEdit create :post="post?.data.value" @update:post="createPost" />
    </div>
  </LayoutAdmin>
</template>

<script setup>
import { message } from "@/lib/core/status"
import { useCreatePostMutation } from "@/lib/api"

const router = useRouter()
const post = useCreatePostMutation()

function createPost(val) {
  post.executeMutation({ input: val }).then((result) => {
    if (!result.error) {
      message.success("Post created successfully")
      router.push({ name: "admin-posts" })
    } else {
      message.error(result.error.toString())
    }
  })
}
</script>
