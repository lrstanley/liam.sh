<template>
  <LayoutAdmin :loading="fetching" :error="error">
    <div class="sm:container sm:mx-auto p-4">
      <n-table v-if="!fetching" v-motion-slide-top bordered single-line striped size="small">
        <thead>
          <tr>
            <th>Title</th>
            <th class="hidden md:table-cell">Slug</th>
            <th class="hidden md:table-cell">Labels</th>
            <th class="hidden md:table-cell">Published</th>
            <th>
              Actions
              <n-button class="ml-10" type="error" @click="regenerate">
                <n-icon><i-mdi-reload /></n-icon>
              </n-button>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="post in posts" :key="post.id">
            <td>
              <router-link :to="{ name: 'posts-slug', params: { slug: post.slug } }">
                {{ post.title }}
              </router-link>
            </td>
            <td class="hidden md:table-cell">
              <router-link :to="{ name: 'posts-slug', params: { slug: post.slug } }">
                {{ post.slug }}
              </router-link>
            </td>
            <td class="hidden md:flex">
              <ObjectRender :value="post.labels" linkable class="mr-1" />
            </td>
            <td class="hidden md:table-cell">{{ useTimeAgo(Date.parse(post.publishedAt)).value }}</td>
            <td>
              <router-link :to="{ name: 'admin-edit-post-id', params: { id: post.id } }">
                <n-button size="small" type="primary" tertiary>
                  <n-icon class="mr-1"><i-mdi-pencil-outline /></n-icon> Edit
                </n-button>
              </router-link>
              <n-button size="small" type="error" tertiary class="ml-2" @click="deletePost(post)">
                <n-icon class="mr-1"><i-mdi-trash-can-outline /></n-icon> Delete
              </n-button>
            </td>
          </tr>
        </tbody>
      </n-table>
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
import { useGetPostsQuery, useDeletePostMutation, useRegeneratePostsMutation } from "@/lib/api"

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
  error,
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

const { executeMutation: regeneratePosts } = useRegeneratePostsMutation()
function regenerate() {
  message.info("Regenerating posts...")
  regeneratePosts().then(({ error }) => {
    if (!error) {
      message.success("Regenerated posts")
    } else {
      message.error(error.toString())
    }
  })
}
</script>
