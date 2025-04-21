<script setup lang="ts">
import { ConfirmModal, UBadge } from "#components"
import type { TableColumn } from "#ui/types"

definePageMeta({
  title: "Posts",
  layout: "admin",
})

const toast = useToast()

const pagination = usePagination<PostSortableFields>({ sort: "published_at" })
const search = useRouteQuery<string>("q", "")
const debounceSearch = refDebounced<string>(search, 300)

// TODO: make this an argument to usePagination?
resetPagination(pagination, [debounceSearch])

const {
  data: posts,
  error,
  status,
  refresh: refreshPosts,
} = await listPosts({
  composable: "useFetch",
  query: computed(() => ({
    page: pagination.page.value,
    per_page: pagination.perPage.value,
    sort: "published_at",
    order: "desc",
    "search.ihas": debounceSearch.value,
  })),
})
if (error.value) throw error.value

const regenerating = ref(false)

async function triggerRegenerate() {
  regenerating.value = true
  regeneratePosts({ composable: "$fetch" })
    .then(() => {
      toast.add({
        title: "Regenerated posts",
        description: "The posts have been regenerated.",
        color: "success",
        duration: 2000,
      })
    })
    .catch((e) => {
      toast.add({
        title: "Error regenerating posts",
        description: e.message,
        color: "error",
        duration: 10000,
      })
    })
    .finally(() => {
      regenerating.value = false
    })
}

const columns: TableColumn<PostRead>[] = [
  { id: "title", header: "Title" },
  { header: "Slug", accessorKey: "slug" },
  { id: "labels", header: "Labels" },
  { id: "views", header: "Views" },
  {
    id: "published",
    header: "Published",
    accessorKey: "published_at",
    accessorFn: ({ published_at }) => useTimeAgo(published_at).value,
  },
  { id: "actions", header: "Actions" },
]

const overlay = useOverlay()
const postToDelete = ref<PostRead | null>(null)
const deleteLoading = ref(false)
const modalDelete = overlay.create(ConfirmModal)

async function promptDeletePost(post: PostRead) {
  deleteLoading.value = true
  postToDelete.value = post
  modalDelete.patch({
    title: `Delete Post #${post.id}?`,
    description: `Are you sure you want to delete "${post.title}"?`,
    confirmText: "Delete",
    color: "error",
  })
  const result = await modalDelete.open()

  if (!result) return

  await deletePost({
    composable: "$fetch",
    path: {
      postID: post.id,
    },
  })

  deleteLoading.value = false
  postToDelete.value = null
  toast.add({
    title: "Post deleted",
    description: `"${post.title}" has been deleted.`,
    color: "success",
    duration: 2000,
  })
  refreshPosts()
}
</script>

<template>
  <div class="flex flex-col">
    <div class="flex flex-auto gap-2 mt-1 mb-4">
      <UInput
        v-model="search"
        :loading="status == 'pending'"
        type="search"
        placeholder="Search for a post"
        icon="mdi:search"
        class="md:min-w-36"
      >
        <template v-if="search.length" #trailing>
          <UButton
            color="neutral"
            variant="link"
            size="sm"
            icon="lucide:circle-x"
            aria-label="Clear input"
            @click="search = ''"
          />
        </template>
      </UInput>

      <div class="ml-auto"></div>

      <UButton
        color="success"
        variant="subtle"
        size="sm"
        icon="mdi:plus-circle"
        :to="{ name: 'admin-posts-new' }"
      >
        New Post
      </UButton>

      <UButton
        color="warning"
        variant="subtle"
        size="sm"
        @click="triggerRegenerate()"
        icon="mdi:reload"
        :loading="regenerating"
      >
        Regenerate posts
      </UButton>

      <UPagination
        v-show="posts?.last_page === undefined || posts?.last_page > 1"
        v-model:page="pagination.page.value"
        :items-per-page="pagination.perPage.value"
        :sibling-count="1"
        active-color="primary"
        active-variant="subtle"
        :show-controls="false"
        show-edges
        :total="posts?.total_count"
        :disabled="status === 'pending'"
      />
    </div>
    <UCard variant="subtle">
      <UTable
        v-if="posts"
        :data="posts.content"
        :columns="columns"
        :loading="status === 'pending'"
        loading-color="primary"
        loading-animation="carousel"
        class="shrink-0"
        :ui="{ td: 'p-2' }"
      >
        <template #title-cell="{ row }">
          <NuxtLink :to="{ name: 'p-slug', params: { slug: row.original.slug } }">
            <span>{{ row.original.title }}</span>
            <UBadge
              variant="outline"
              :color="row.original.public ? 'success' : 'warning'"
              size="sm"
              class="ml-2 text-xs"
            >
              {{ row.original.public ? "public" : "draft" }}
            </UBadge>
          </NuxtLink>
        </template>

        <template #labels-cell="{ row }">
          <UPopover
            v-if="row.original.edges.labels"
            mode="hover"
            scrollable
            placement="left"
            :content="{
              align: 'end',
              side: 'left',
              sideOffset: 8,
            }"
          >
            <UButton size="xs" variant="subtle">{{ row.original.edges.labels.length }} labels</UButton>

            <template #content>
              <div class="flex flex-wrap flex-auto gap-1 p-1 max-w-40">
                <LabelObject
                  v-for="label in row.original.edges.labels"
                  :key="label.id"
                  :value="label"
                  route="/posts"
                  linkable
                  class="hidden md:flex"
                />
              </div>
            </template>
          </UPopover>
        </template>

        <template #views-cell="{ row }">
          <StatsButton>
            {{ row.original.view_count.toLocaleString() }}
            {{ row.original.view_count === 1 ? "view" : "views" }}
          </StatsButton>
        </template>

        <template #actions-cell="{ row }">
          <div class="flex flex-row gap-2">
            <UButton
              size="xs"
              variant="subtle"
              color="warning"
              icon="mdi:pencil"
              :to="{ name: 'admin-posts-id-edit', params: { id: row.original.id } }"
            >
              Edit
            </UButton>
            <UButton
              size="xs"
              variant="subtle"
              color="error"
              icon="mdi:delete"
              @click="promptDeletePost(row.original)"
            >
              Delete
            </UButton>
          </div>
        </template>
      </UTable>
    </UCard>
  </div>
</template>
