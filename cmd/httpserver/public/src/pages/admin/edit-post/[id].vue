<script setup lang="ts">
import { message } from "@/lib/core/status"
import { useGetPostQuery, useUpdatePostMutation, type UpdatePostInput } from "@/lib/api"
import type { Post } from "@/lib/api"

definePage({
  meta: {
    layout: "admin",
  },
})

const route = useRoute("/admin/edit-post/[id]")

const router = useRouter()
const { data, error, fetching } = await useGetPostQuery({
  variables: { id: route.params.id },
  requestPolicy: "network-only",
})

watch(error, () => {
  if (error.value) throw error.value
})

const post = computed(() => data.value?.node as Post | undefined)
const update = useUpdatePostMutation()

function updatePost(val: UpdatePostInput, labelIDs: string[]) {
  const addedLabels = labelIDs.filter(
    (id) => !post.value.labels?.edges?.map(({ node }) => node.id).includes(id)
  )
  const removedLabels = post.value.labels?.edges
    ?.map(({ node }) => node.id)
    .filter((id) => !labelIDs.includes(id))

  update
    .executeMutation({
      id: route.params.id,
      input: {
        title: val.title,
        content: val.content,
        slug: val.slug,
        addLabelIDs: addedLabels,
        removeLabelIDs: removedLabels,
        publishedAt: val.publishedAt,
        public: val.public,
      },
    })
    .then((result) => {
      if (!result.error) {
        message.success("Post updated")
        router.push({ name: "/admin/posts" })
      } else {
        message.error("Error updating post: " + result.error.toString())
      }
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

    <n-spin :show="fetching">
      <n-alert v-if="error" title="Error fetching post" type="error">
        {{ error }}
      </n-alert>

      <PostCreateEdit v-if="post" :post="post" @update:post="updatePost" />
    </n-spin>
  </div>
</template>
