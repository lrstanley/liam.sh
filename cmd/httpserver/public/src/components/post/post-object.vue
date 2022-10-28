<template>
  <component
    :is="props.linkable ? 'router-link' : 'div'"
    :to="{ name: '/p/[slug]', params: { slug: post.slug } }"
  >
    <n-thing class="mb-7" content-indented v-bind="$attrs">
      <template #avatar>
        <n-avatar class="hidden lg:inline-block" round size="medium" :src="post.author.avatarURL" />
      </template>
      <template #header>
        <span class="post-title text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500">
          {{ post.title }}
        </span>
      </template>
      <template #description>
        <span>
          <i>
            Published {{ useTimeAgo(post.publishedAt).value }} by
            <a :href="post.author.htmlURL" target="_blank">{{ post.author.name }}</a>
          </i>
        </span>
      </template>

      <span v-html="post.summary" />

      <template v-if="post.labels" #action>
        <div class="flex justify-between flex-auto">
          <div class="inline-flex flex-wrap flex-auto gap-1">
            <LabelObject
              v-for="label in post.labels.edges.map(({ node }) => node)"
              :key="label.id"
              :value="label"
              linkable
            />
          </div>

          <PostViewCount :value="post.viewCount" />
          <n-tag v-if="!post.public" class="ml-2" type="warning">draft</n-tag>
        </div>
      </template>
    </n-thing>
  </component>
</template>

<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import type { Post } from "@/lib/api"

const props = defineProps<{
  value: Post
  linkable?: boolean
}>()

const post = ref(props.value)
</script>

<style scoped>
.post-title {
  @apply text-[1.5em];
}

.n-thing :deep(.n-thing-avatar) {
  @apply hidden md:inline-flex;
}
</style>
