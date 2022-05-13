<template>
  <LayoutAdmin>
    <div class="sm:container sm:mx-auto flex flex-auto flex-col flex-nowrap mt-7">
      <CoreDataTable :data="posts" :loading="fetching.value" :headers="{ published_at: 'Published' }">
        <template #title="{ row }">{{ row.title }}</template>
        <template #slug="{ row }">{{ row.slug }}</template>
        <template #published="{ row }">
          {{ useTimeAgo(Date.parse(row.publishedAt)).value }}
        </template>
        <template #actions="{ row }">
          <router-link :to="{ name: 'admin-edit-post-id', params: { id: row.id } }">
            <n-button size="small" type="primary" tertiary>
              <n-icon class="mr-1"><i-mdi-pencil-outline /></n-icon> Edit
            </n-button>
          </router-link>
          <n-button size="small" type="error" tertiary class="ml-2" @click="deletePost(row)">
            <n-icon class="mr-1"><i-mdi-trash-can-outline /></n-icon> Delete
          </n-button>
        </template>
      </CoreDataTable>
    </div>
    <router-link :to="{ name: 'admin-new-post' }" class="no-underline absolute bottom-5 right-5">
      <n-button tertiary circle size="large" type="primary" class="h-13 w-13">
        <n-icon class="text-2em"><i-mdi-pencil-outline /></n-icon>
      </n-button>
    </router-link>
  </LayoutAdmin>
</template>

<script setup>
import { useDialog, useMessage } from "naive-ui"
import { useTimeAgo } from "@vueuse/core"
import { useGetPostsQuery, useDeletePostMutation } from "@/lib/api"

const props = defineProps({
  cursor: {
    type: String,
    default: null,
  },
})

const dialog = useDialog()
const message = useMessage()
const {
  data,
  fetching,
  executeQuery: refetch,
} = useGetPostsQuery({
  variables: { count: 100, cursor: props.cursor },
})
const del = useDeletePostMutation()

const posts = computed({
  get: () => data?.value?.posts.edges.map((v) => v.node),
})

function deletePost(row) {
  dialog.warning({
    title: `Delete post: "${row.title}"`,
    content: "Are you sure?",
    positiveText: "Delete",
    negativeText: "Cancel",
    onPositiveClick: () => {
      del.executeMutation({ id: row.id }).then((result) => {
        if (!result.error) {
          message.success("Post deleted")
        } else {
          message.error(result.error.toString())
        }

        refetch()
      })
    },
  })
}
</script>
