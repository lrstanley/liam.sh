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
      <PostCreateEdit :post="post" @update:post="createPost" />
    </div>
  </LayoutAdmin>
</template>

<script setup>
import { useMessage } from "naive-ui"
import { query } from "@/lib/http"

const router = useRouter()
const state = useState()
const message = useMessage()
const post = ref({})

function createPost(val) {
  val.author = state.auth.id

  query.post
    .createPost(val)
    .then(() => {
      message.success("Post created")
      router.push({ name: "admin-posts" })
    })
    .catch((err) => {
      message.error(err)
    })
}
</script>
