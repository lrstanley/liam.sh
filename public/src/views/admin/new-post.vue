<template>
  <LayoutAdmin>
    <n-page-header class="px-5 mt-4 mb-8">
      <template #avatar>
        <n-icon :size="40"><Create /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit"> Create new post </a>
      </template>
    </n-page-header>

    <div class="sm:container sm:mx-auto flex flex-auto flex-col flex-nowrap">
      <PostCreateEdit :post="post?.data.value" @update:post="createPost" />
    </div>
  </LayoutAdmin>
</template>

<script setup>
import { useMessage } from "naive-ui"
import { useCreatePostMutation } from "@/lib/api"

const router = useRouter()
const message = useMessage()
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
