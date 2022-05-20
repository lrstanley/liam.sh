<template>
  <LayoutDefault :loading="fetching" :error="error">
    <CoreTableOfContents :element="postRef" />
    <n-page-header class="container mx-auto hidden md:inline-flex xl:px-200px mt-4 mb-8">
      <template #title>
        <CoreTerminal
          class="mb-4 text-size-38px md:text-size-34px"
          path="posts"
          prefix=""
          :value="'cat &quot;' + post.slug + '.md&quot;'"
        />
      </template>
    </n-page-header>

    <div class="container mx-auto px-15px xl:px-200px">
      <n-alert v-if="error" title="Error fetching post" type="error">
        {{ error }}
      </n-alert>

      <div v-motion-fade class="h-full flex-auto flex-col">
        <div
          class="text-size-50px text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500"
        >
          {{ post.title }}
        </div>

        <div class="flex flex-auto flex-row flex-wrap items-center mt-3 mb-8">
          <n-avatar class="mr-3" round size="medium" :src="post.author.avatarURL" />
          <p>
            <a :href="post.author.htmlURL" target="_blank">{{ post.author.name }}</a>
            <br />
            <i>Published {{ useTimeAgo(post.publishedAt).value }}</i>
          </p>
          <span class="ml-auto">
            <router-link
              v-if="state.base?.self"
              :to="{ name: 'admin-edit-post-id', params: { id: post.id } }"
            >
              <n-button class="mr-3" type="secondary"> Edit post </n-button>
            </router-link>
          </span>
        </div>

        <div id="post-content" ref="postRef" v-html="post.contentHTML" />
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup>
import { useTimeAgo } from "@vueuse/core"
import { useGetPostContentQuery } from "@/lib/api"

const props = defineProps({
  slug: {
    type: String,
    required: true,
  },
})

const state = useState()
const { data, error, fetching } = useGetPostContentQuery({ variables: { slug: props.slug } })
const post = computed(() => data?.value?.posts?.edges[0].node)
const postRef = ref(null)
</script>

<style scoped>
#post-content {
  @apply text-size-1.2em leading-relaxed;
}

#post-content :deep(img) {
  max-width: 100%;
  height: auto;
}

#post-content :deep(p) {
  margin-top: 25px;
}

#post-content :deep(blockquote) {
  line-height: 1.5rem;
  padding: 0.6rem 1.2rem;
  opacity: 0.8;
  margin-left: 0;
  color: inherit;
  border-left-width: 0.25rem;
  /* border-color: #e5e7eb; */
  margin-top: 1.6em;
  margin-bottom: 1.6em;
  @apply border-green-600;
}

#post-content :deep(blockquote) > p {
  margin-top: 0;
  margin-bottom: 0;
}

#post-content :deep(h2),
#post-content :deep(h3),
#post-content :deep(h4) {
  margin-top: 30px;
  @apply text-transparent bg-gradient-to-tr bg-clip-text font-bold;
  @apply bg-gradient-to-r from-sky-400 to-blue-500;
}

#post-content :deep(pre) {
  border-radius: 7px;
  padding: 0.5rem 0.8rem;
  white-space: pre-wrap;
  @apply !bg-dark-400;
}
</style>
