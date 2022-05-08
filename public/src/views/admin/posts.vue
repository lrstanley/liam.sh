<template>
  <LayoutAdmin>
    <div class="sm:container sm:mx-auto flex flex-auto flex-col flex-nowrap mt-7">
      <n-data-table :columns="columns" :data="posts" :loading="loading" />
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
import { NButton, NTime, useDialog, useMessage } from "naive-ui"
import { query } from "@/lib/http"
import { h } from "vue"

const router = useRouter()
const dialog = useDialog()
const message = useMessage()

const posts = ref([])
const loading = ref(true)

const columns = [
  {
    title: "Title",
    key: "title",
  },
  {
    title: "Slug",
    key: "slug",
  },
  {
    title: "Published",
    key: "published_at",
    render(row) {
      return h(NTime, {
        time: Date.parse(row.published_at),
        type: "relative",
      })
    },
  },
  {
    title: "Action",
    key: "actions",
    render(row) {
      return [
        h(
          NButton,
          {
            size: "small",
            onClick: () => router.push({ name: "admin-edit-post-id", params: { id: row.id } }),
          },
          { default: () => "Edit" }
        ),
        h(
          NButton,
          {
            size: "small",
            style: "margin-left: 5px",
            onClick: () => deletePost(row),
          },
          { default: () => "Delete" }
        ),
      ]
    },
  },
]

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
