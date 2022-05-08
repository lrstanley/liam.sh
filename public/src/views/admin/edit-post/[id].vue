<template>
  <LayoutAdmin>
    <n-page-header :subtitle="post.title" class="px-5 mt-4 mb-8">
      <template #avatar>
        <n-icon :size="40"><Create /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">
          Editing post #{{ post.id }}
        </a>
      </template>
    </n-page-header>

    <div class="sm:container sm:mx-auto flex flex-auto flex-col flex-nowrap">
      <n-spin :show="loading">
        <n-alert v-if="error" title="Error fetching post" type="error">
          {{ error }}
        </n-alert>

        <PostCreateEdit v-if="post.title" :post="post" @update:post="updatePost" />
      </n-spin>
    </div>
  </LayoutAdmin>
</template>

<script setup>
import { useMessage } from "naive-ui"
import { query } from "@/lib/http"

const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})

const router = useRouter()
const state = useState()
const message = useMessage()
const loading = ref(true)
const error = ref(null)
const post = ref({})

function updatePost(val) {
  loading.value = true
  val.author = state.auth.id

  query.post
    .updatePost(val.id, val)
    .then(() => {
      message.success("Post updated")
      router.push({ name: "admin-posts" })
    })
    .catch((err) => {
      message.error(err)
    })
    .finally(() => {
      setTimeout(() => (loading.value = false), 200)
    })
}

onMounted(() => {
  query.post
    .readPost(props.id)
    .then((data) => {
      post.value = data
    })
    .catch((err) => {
      error.value = err
    })
    .finally(() => {
      loading.value = false
    })
})
</script>
