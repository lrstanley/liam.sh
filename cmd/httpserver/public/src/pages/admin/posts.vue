<route lang="yaml">
meta:
  layout: admin
</route>

<template>
  <div class="p-4 sm:container sm:mx-auto">
    <n-table v-motion-slide-top bordered single-line striped size="small">
      <thead>
        <tr>
          <th>Title</th>
          <th class="hidden md:table-cell">Slug</th>
          <th class="hidden md:table-cell">Labels</th>
          <th class="hidden md:table-cell">Views</th>
          <th class="hidden md:table-cell">Published</th>
          <th>
            Actions
            <n-button class="ml-10" type="error" size="small" @click="regenerate">
              <n-icon><i-mdi-reload /></n-icon>
            </n-button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="post in posts" :key="post.id">
          <td>
            <router-link :to="{ name: '/p/[slug]', params: { slug: post.slug } }">
              {{ post.title }}
              <span v-if="!post.public" class="text-yellow-500">[draft]</span>
            </router-link>
          </td>
          <td class="hidden md:table-cell">
            <router-link :to="{ name: '/p/[slug]', params: { slug: post.slug } }">
              {{ post.slug }}
            </router-link>
          </td>
          <td class="hidden md:flex">
            <n-popover
              style="max-height: 240px"
              trigger="hover"
              content-style="padding: 0;"
              scrollable
              placement="left"
            >
              <template #trigger>
                <n-button size="small">{{ post.labels.edges.length }} tags</n-button>
              </template>

              <CoreObjectRender :value="post.labels" linkable class="grid mx-1 my-[2px]" />
            </n-popover>
          </td>
          <td class="hidden md:table-cell">
            <PostViewCount :value="post.viewCount" />
          </td>
          <td class="hidden md:table-cell">{{ useTimeAgo(Date.parse(post.publishedAt)).value }}</td>
          <td>
            <router-link :to="{ name: '/admin/edit-post/[id]', params: { id: post.id } }">
              <n-button size="small" type="primary" tertiary>
                <n-icon class="mr-1"><i-mdi-pencil-outline /></n-icon>
                Edit
              </n-button>
            </router-link>
            <n-button size="small" type="error" tertiary class="ml-2" @click="deletePost(post)">
              <n-icon class="mr-1"><i-mdi-trash-can-outline /></n-icon>
              Delete
            </n-button>
          </td>
        </tr>
      </tbody>
    </n-table>

    <router-link :to="{ name: '/admin/new-post' }" class="absolute no-underline bottom-5 right-5">
      <n-button tertiary circle size="large" type="primary" class="h-[13] w-[13]">
        <n-icon class="text-[2em]"><i-mdi-pencil-outline /></n-icon>
      </n-button>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { message, dialog } from "@/lib/core/status"
import { useTimeAgo } from "@vueuse/core"
import { useGetPostsQuery, useDeletePostMutation, useRegeneratePostsMutation } from "@/lib/api"

const props = defineProps({
  cursor: {
    type: String,
    default: null,
  },
})

const {
  data,
  error,
  executeQuery: refetch,
} = await useGetPostsQuery({
  variables: { first: 100, before: props.cursor },
})

watch(error, () => {
  if (error.value) throw error.value
})

const del = useDeletePostMutation()

const posts = computed(() => data?.value?.posts.edges.map((v) => v.node))

function deletePost(row: Record<string, any>) {
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
  regeneratePosts({}).then(({ error }) => {
    if (!error) {
      message.success("Regenerated posts")
    } else {
      message.error(error.toString())
    }
  })
}
</script>
