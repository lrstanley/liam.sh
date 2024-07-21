<script setup lang="ts">
import { message } from "@/lib/core/status"
import { getPost, updatePost } from "@/lib/http/services.gen"
import type { PostRead, PostUpdate } from "@/lib/http/types.gen"

definePage({
  meta: {
    layout: "admin",
  },
})

const route = useRoute("/admin/edit-post/[id]")
const router = useRouter()
const queryClient = useQueryClient()

const { data: post, error } = await getPost({ path: { postID: parseInt(route.params.id) } })
if (error) throw error

const { mutate, isPending } = useMutation({
  mutationFn: (data: PostUpdate) =>
    unwrapErrors(updatePost({ path: { postID: parseInt(route.params.id) }, body: data })),
  onError: (error) => {
    message.error("error updating post: " + error.message)
  },
  onSuccess: () => {
    queryClient.invalidateQueries({ queryKey: ["posts"] })
    message.success("Post updated")
    router.push({ name: "/admin/posts" })
  },
})

function invokeUpdate(data: PostRead, labelIDs: number[]) {
  const addedLabels = labelIDs.filter((id) => !post.edges.labels?.map((label) => label.id).includes(id))
  const removedLabels = post.edges.labels
    ?.map((label) => label.id)
    .filter((id) => !labelIDs.includes(id))

  mutate({
    slug: data.slug,
    title: data.title,
    content: data.content,
    published_at: data.published_at,
    public: data.public,
    add_labels: addedLabels,
    remove_labels: removedLabels,
  })
}
</script>

<template>
  <div class="p-4 sm:container sm:mx-auto lg:p-0">
    <n-page-header :subtitle="post.title" class="hidden mt-4 mb-8 lg:block">
      <template #avatar>
        <n-icon :size="40"><i-mdi-pencil-outline /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">
          Editing post #{{ post.id }}
        </a>
      </template>
    </n-page-header>

    <n-spin :show="isPending">
      <CoreError v-if="error" :error="error" />

      <PostCreateEdit v-if="post" :post="post" @update:post="invokeUpdate" />
    </n-spin>
  </div>
</template>
