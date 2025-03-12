<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import type { PostRead } from "@/utils/http/types.gen"

const props = defineProps<{
  value: PostRead
  linkable?: boolean
}>()

const post = ref(props.value)
</script>

<template>
  <component
    :is="props.linkable ? 'router-link' : 'div'"
    :to="{ name: '/p/[slug]', params: { slug: post.slug } }"
  >
    <n-thing class="mb-7" content-indented v-bind="$attrs">
      <template #avatar>
        <n-avatar
          class="hidden lg:inline-block"
          round
          size="medium"
          :src="post.edges.author.avatar_url + '&s=40'"
        />
      </template>
      <template #header>
        <span class="post-title text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500">
          {{ post.title }}
        </span>
      </template>
      <template #description>
        <span>
          <i>
            Published {{ useTimeAgo(post.published_at).value }} by
            <a :href="post.edges.author.html_url" target="_blank">{{ post.edges.author.name }}</a>
          </i>
        </span>
      </template>

      <span v-html="post.summary" />

      <template v-if="post.edges.labels" #action>
        <div class="flex justify-between flex-auto">
          <div class="inline-flex flex-wrap flex-auto gap-1">
            <LabelObject v-for="label in post.edges.labels" :key="label.id" :value="label" linkable />
          </div>

          <PostViewCount :value="post.view_count" />
          <n-tag v-if="!post.public" class="ml-2" type="warning">draft</n-tag>
        </div>
      </template>
    </n-thing>
  </component>
</template>

<style scoped>
@reference "@/assets/css/main.css";

.post-title {
  @apply text-[1.5em];
}

.n-thing :deep(.n-thing-avatar) {
  @apply hidden md:inline-flex;
}
</style>
