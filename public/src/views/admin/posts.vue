<template>
  <LayoutAdmin>
    <div class="sm:container sm:mx-auto flex flex-auto flex-col flex-nowrap mt-7">
      <CoreDataTable :data="posts" :loading="loading" :headers="{ published_at: 'Published' }">
        <template #title="{ row }">{{ row.title }}</template>
        <template #slug="{ row }">{{ row.slug }}</template>
        <template #published="{ row }">
          {{ useTimeAgo(Date.parse(row.published_at)).value }}
        </template>
        <template #actions="{ row }">
          <router-link :to="{ name: 'admin-edit-post-id', params: { id: row.id } }">
            <n-button size="small" type="primary" tertiary> <CreateOutline /> Edit </n-button>
          </router-link>
          <n-button size="small" type="error" tertiary class="ml-2" @click="deletePost(row)">
            <Trash /> Delete
          </n-button>
        </template>
      </CoreDataTable>
    </div>

    <router-link :to="{ name: 'admin-new-post' }" class="no-underline absolute bottom-5 right-5">
      <n-button tertiary circle size="large" type="primary" class="h-13 w-13">
        <template #icon>
          <n-icon><Create /></n-icon>
        </template>
      </n-button>
    </router-link>
  </LayoutAdmin>
</template>

<script setup>
import { useDialog, useMessage } from "naive-ui"
import { useTimeAgo } from "@vueuse/core"
import { query } from "@/lib/http"

const dialog = useDialog()
const message = useMessage()

const posts = ref([])
const loading = ref(true)

function fetchPosts() {
  loading.value = true

  query.post
    .listPost(1, 1000)
    .then((res) => {
      posts.value = res
    })
    .finally(() => {
      loading.value = false
    })
}

function deletePost(row) {
  dialog.warning({
    title: `Delete post: "${row.title}"`,
    content: "Are you sure?",
    positiveText: "Delete",
    negativeText: "Cancel",
    onPositiveClick: () => {
      query.post
        .deletePost(row.id)
        .then(() => {
          message.success("Post deleted")
        })
        .catch((err) => {
          message.error(err)
        })
        .finally(() => {
          fetchPosts()
        })
    },
  })
}

onMounted(() => {
  fetchPosts()
})
</script>
