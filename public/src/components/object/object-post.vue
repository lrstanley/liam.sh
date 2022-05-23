<template>
  <component
    :is="props.linkable ? 'router-link' : 'div'"
    :to="{ name: 'posts-slug', params: { slug: props.value.slug } }"
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
          {{ props.value.title }}
        </span>
      </template>
      <template #description>
        <span class="">
          <i>
            Published {{ useTimeAgo(props.value.publishedAt).value }} by
            <a :href="props.value.author.htmlURL" target="_blank">{{ props.value.author.name }}</a>
          </i>
        </span>
      </template>

      <span v-html="props.value.summary" />

      <template #action>
        <ObjectRender :value="props.value.labels" linkable class="mr-1" />
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
</script>

<style scoped>
.post-title {
  @apply text-size-1.5em;
}
</style>
