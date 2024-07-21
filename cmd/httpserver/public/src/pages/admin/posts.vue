<script setup lang="ts">
import { message, dialog } from "@/lib/core/status"
import { useTimeAgo } from "@vueuse/core"
import { listPosts, deletePost, regeneratePosts } from "@/lib/http/services.gen"
import type { PostSortableFields, PostRead } from "@/lib/http/types.gen"
import { RouterLink } from "vue-router"

definePage({
  meta: {
    layout: "admin",
  },
})

const pagination = usePagination<PostSortableFields>({ sort: "published_at" })
const queryClient = useQueryClient()

const {
  data: posts,
  error,
  isFetching,
} = useQuery({
  queryKey: ["posts", "admin", pagination],
  queryFn: () =>
    unwrapErrors(
      listPosts({
        query: {
          page: pagination.page.value,
          per_page: pagination.perPage.value,
          sort: pagination.sort.value,
          order: pagination.order.value,
        },
      })
    ),
  placeholderData: keepPreviousData,
  retry: false,
})

const { mutate: deletePostMutation } = useMutation({
  mutationFn: (id: number) => unwrapErrors(deletePost({ path: { postID: id + 10000 } })),
  onError: (error) => {
    message.error("error deleting post: " + error.message)
  },
  onSuccess: () => {
    queryClient.invalidateQueries({ queryKey: ["posts"] })
    message.success("post deleted")
  },
})

const { mutate: regeneratePostsMutation } = useMutation({
  mutationFn: () => unwrapErrors(regeneratePosts()),
  onError: (error) => {
    message.error("error regenerating posts: " + error.message)
  },
  onSuccess: () => {
    message.success("posts regenerated")
    return queryClient.invalidateQueries({ queryKey: ["posts"] })
  },
})

function invokeDeletePost(post: PostRead) {
  dialog.warning({
    title: `Delete post: "${post.title}"`,
    content: "Are you sure?",
    positiveText: "Delete",
    negativeText: "Cancel",
    onPositiveClick: () => deletePostMutation(post.id),
  })
}
</script>

<template>
  <div class="p-4 sm:container sm:mx-auto">
    <n-page-header class="hidden mt-4 mb-8 lg:block" subtitle="Manage blog posts">
      <template #avatar>
        <n-icon :size="40"><i-mdi-book-open-outline /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">Blog Posts</a>
      </template>
    </n-page-header>

    <CoreError v-if="error" :error="error" />
    <n-spin v-else :show="isFetching" :delay="250">
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
              <n-button class="ml-10" type="error" size="small" @click="() => regeneratePostsMutation()">
                <n-icon><i-mdi-reload /></n-icon>
              </n-button>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="post in posts?.content" :key="post.id">
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
                  <n-button size="small">{{ post.edges.labels.length }} tags</n-button>
                </template>

                <CoreObjectRender
                  :value="post.edges.labels"
                  type="label"
                  linkable
                  class="grid mx-1 my-[2px]"
                />
              </n-popover>
            </td>
            <td class="hidden md:table-cell">
              <PostViewCount :value="post.view_count" />
            </td>
            <td class="hidden md:table-cell">{{ useTimeAgo(Date.parse(post.published_at)).value }}</td>
            <td>
              <router-link :to="{ name: '/admin/edit-post/[id]', params: { id: post.id } }">
                <n-button size="small" type="primary" tertiary>
                  <n-icon class="mr-1"><i-mdi-pencil-outline /></n-icon>
                  Edit
                </n-button>
              </router-link>
              <n-button size="small" type="error" tertiary class="ml-2" @click="invokeDeletePost(post)">
                <n-icon class="mr-1"><i-mdi-trash-can-outline /></n-icon>
                Delete
              </n-button>
            </td>
          </tr>
        </tbody>
      </n-table>
      <div class="flex">
        <n-pagination
          v-model:page="pagination.page.value"
          v-model:page-size="pagination.perPage.value"
          :item-count="posts?.total_count"
          :page-sizes="pagination.sizes"
          show-size-picker
          class="mt-2 ml-auto"
        />
      </div>
    </n-spin>

    <router-link :to="{ name: '/admin/new-post' }" class="absolute no-underline bottom-5 right-5">
      <n-button tertiary circle size="large" type="primary" class="size-[13]">
        <n-icon class="text-[2em]"><i-mdi-pencil-outline /></n-icon>
      </n-button>
    </router-link>
  </div>
</template>
