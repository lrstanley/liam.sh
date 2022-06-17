<template>
  <component
    :is="props.linkable ? 'router-link' : 'div'"
    :to="{ name: 'p-slug', params: { slug: post.slug } }"
  >
    <n-thing class="mb-7" content-indented v-bind="$attrs">
      <template #avatar>
        <n-avatar>
          <n-icon>
            <i-mdi-post-outline />
          </n-icon>
        </n-avatar>
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
        <div class="flex flex-auto justify-between">
          <div class="inline-flex flex-auto flex-wrap gap-1">
            <ObjectLabel
              v-for="label in post.labels.edges.map(({ node }) => node)"
              :key="label.id"
              :value="label"
              linkable
            />
          </div>

          <n-tag class="text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500">
            {{ post.viewCount.toLocaleString() }}
            {{ post.viewCount === 1 ? "view" : "views" }}
          </n-tag>
        </div>
      </template>
    </n-thing>
  </component>
</template>

<script setup>
import { useTimeAgo } from "@vueuse/core"

const props = defineProps({
  linkable: {
    type: Boolean,
    default: false,
  },
  value: {
    type: Object,
    required: true,
  },
})

const post = ref(props.value)
</script>

<style scoped>
.post-title {
  @apply text-size-1.5em;
}
</style>
