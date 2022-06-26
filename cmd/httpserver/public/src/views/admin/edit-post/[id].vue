<template>
  <LayoutAdmin>
    <n-page-header :subtitle="data?.node.title" class="px-5 mt-4 mb-8 hidden lg:block">
      <template #avatar>
        <n-icon :size="40"><i-mdi-pencil-outline /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">
          Editing post #{{ data?.node.id }}
        </a>
      </template>
    </n-page-header>

    <div class="sm:container sm:mx-auto p-4 lg:p-0">
      <n-spin :show="fetching">
        <n-alert v-if="error" title="Error fetching post" type="error">
          {{ error }}
        </n-alert>

        <PostCreateEdit v-if="data?.node" :post="data.node" @update:post="updatePost" />
      </n-spin>
    </div>
  </LayoutAdmin>
</template>

<script setup>
import { message } from "@/lib/core/status"
import { useGetPostQuery, useUpdatePostMutation } from "@/lib/api"

const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})

const router = useRouter()
const { data, error, fetching } = useGetPostQuery({
  variables: { id: props.id },
  requestPolicy: "network-only",
})
const update = useUpdatePostMutation()

function updatePost(val) {
  let addedLabels = val.labelIDs.filter(
    (id) => !data.value?.node.labels?.edges?.map(({ node }) => node.id).includes(id)
  )
  let removedLabels = data.value?.node.labels?.edges
    ?.map(({ node }) => node.id)
    .filter((id) => !val.labelIDs.includes(id))

  update
    .executeMutation({
      id: props.id,
      input: {
        title: val.title,
        content: val.content,
        slug: val.slug,
        addLabelIDs: addedLabels,
        removeLabelIDs: removedLabels,
        publishedAt: val.publishedAt,
      },
    })
    .then((result) => {
      if (!result.error) {
        message.success("Post updated")
        router.push({ name: "admin-posts" })
      } else {
        message.error("Error updating post: " + result.error.toString())
      }
    })
}
</script>
