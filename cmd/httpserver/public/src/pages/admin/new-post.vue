<script setup lang="ts">
import { message } from "@/lib/core/status"
import { createPost } from "@/lib/http/services.gen"
import type { PostCreate } from "@/lib/http/types.gen"

definePage({
  meta: {
    layout: "admin",
  },
})

const router = useRouter()
const state = useState()
const queryClient = useQueryClient()

const { mutate, isPending } = useMutation({
  mutationFn: (data: PostCreate) => unwrapErrors(createPost({ body: data })),
  onError: (error) => {
    message.error("error creating post: " + error.message)
  },
  onSuccess: () => {
    queryClient.invalidateQueries({ queryKey: ["posts"] })
    message.success("post created successfully")
    router.push({ name: "/admin/posts" })
  },
})

function invokeCreate(data: PostCreate, labelIDs: number[]) {
  mutate({
    slug: data.slug,
    title: data.title,
    content: data.content,
    published_at: data.published_at,
    public: data.public,
    labels: labelIDs,
    author: state.user.id,
  })
}
</script>

<template>
  <div>
    <n-page-header class="hidden px-5 mt-4 mb-8 md:block">
      <template #avatar>
        <n-icon :size="40"><i-mdi-pencil-outline /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">Create new post</a>
      </template>
    </n-page-header>

    <div class="p-4 sm:container sm:mx-auto lg:p-0">
      <n-spin :show="isPending">
        <PostCreateEdit create @update:post="invokeCreate" />
      </n-spin>
    </div>
  </div>
</template>
