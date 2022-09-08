<template>
  <LayoutDefault :loading="fetching" :error="error">
    <CoreTableOfContents :element="postRef" />
    <n-page-header class="container hidden mb-2 md:inline-flex mt-14">
      <template #title>
        <CoreTerminal
          class="text-[20px]"
          path="posts"
          prefix=""
          :value="'cat &quot;' + post.slug + '.md&quot;'"
        />
      </template>
    </n-page-header>

    <div class="container">
      <div v-motion-fade class="flex-col flex-auto h-full mt-7 md:mt-0">
        <div
          class="text-[30px] md:text-[45px] text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500"
        >
          {{ post.title }}
        </div>

        <div
          class="flex flex-col flex-wrap items-start flex-auto mt-3 lg:flex-row lg:items-center mb-7 md:mb-20"
        >
          <span class="inline-flex">
            <n-avatar class="mr-3" round size="medium" :src="post.author.avatarURL" />
            <p>
              <a :href="post.author.htmlURL" target="_blank">{{ post.author.name }}</a>
              <br />
              <i>Published {{ useTimeAgo(post.publishedAt).value }}</i>
            </p>
          </span>
          <span class="inline-flex flex-wrap items-center mt-4 ml-0 mr-auto md:mt-0 md:mr-0 md:ml-auto">
            <CoreObjectRender :value="post.labels" linkable class="mr-1" />
            <router-link
              v-if="state.base?.self"
              class="ml-1"
              :to="{ name: 'admin-edit-post-id', params: { id: post.id } }"
            >
              <n-button class="mr-3" type="success" tertiary> Edit post </n-button>
            </router-link>
          </span>
        </div>

        <div id="post-content" ref="postRef" class="lg:mb-[100px]" v-html="post.contentHTML" />
      </div>
    </div>
  </LayoutDefault>
</template>

<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import { useGetPostContentQuery } from "@/lib/api"

const props = defineProps<{
  slug: string
}>()

const state = useState()
const { data, error, fetching } = useGetPostContentQuery({ variables: { slug: props.slug } })
const post = computed(() => data?.value?.posts?.edges[0].node)
const postRef = ref(null)
</script>

<style scoped>
#post-content {
  font-size: 1.1em;
  line-height: 1.55;
}

@screen lg {
  #post-content {
    font-size: 1.2em;
  }
}

#post-content :deep(center) {
  margin: 1em;
}

#post-content :deep(center) p {
  margin-top: 0.5em;
  @apply text-center;
}

#post-content :deep(img) {
  @apply max-w-[calc(100%)] lg:max-w-[calc(80%)] px-2 lg:px-0 !m-0;
  height: auto;
}

#post-content :deep(p) {
  margin-top: 1em;
  margin-bottom: 1em;
  @apply text-center md:text-left;
}

#post-content :deep(blockquote) {
  line-height: 1.5rem;
  padding: 0.6rem 1.2rem;
  opacity: 0.8;
  margin-left: 0;
  margin-right: 0;
  color: inherit;
  border-left-width: 0.25rem;
  margin-top: 1.6em;
  margin-bottom: 1.6em;
  @apply border-l-emerald-600 border-l-4 border-solid bg-emerald-600/30 rounded;
}

#post-content :deep(blockquote) > p {
  margin-top: 0;
  margin-bottom: 0;
  text-indent: 0;
}

#post-content > :deep(p):not(blockquote) > strong {
  @apply text-lime-400;
}

#post-content :deep(h1) {
  font-size: 1.8em;
  margin-top: 1.6em;
}
#post-content :deep(h2) {
  font-size: 1.65em;
  margin-top: 1.5em;
}
#post-content :deep(h3) {
  font-size: 1.5em;
  margin-top: 1.4em;
}
#post-content :deep(h4) {
  font-size: 1.4em;
  margin-top: 1.3em;
}
#post-content :deep(h5) {
  font-size: 1.3em;
  margin-top: 1.2em;
}

#post-content :deep(h1),
#post-content :deep(h2),
#post-content :deep(h3),
#post-content :deep(h4),
#post-content :deep(h5) {
  @apply text-transparent bg-gradient-to-tr bg-clip-text font-bold;
  @apply bg-gradient-to-r from-sky-400 to-blue-500;
}

#post-content :deep(h1)::before,
#post-content :deep(h2)::before,
#post-content :deep(h3)::before,
#post-content :deep(h4)::before,
#post-content :deep(h5)::before,
#post-content :deep(h6)::before {
  content: "#";
  position: relative;
  right: 10px;
  @apply text-emerald-500;
}

#post-content :deep(:not(pre)) > code {
  border-radius: 3px;
  padding: 0.18em 0.35em;
  @apply !bg-zinc-700;
}

#post-content :deep(pre) {
  margin-top: 1.5rem;
  margin-bottom: 2rem;
  border-radius: 7px;
  padding: 0.5rem 0.8rem;
  @apply !bg-[#1c1c1e] overflow-auto whitespace-pre lg:overflow-hidden lg:whitespace-pre-wrap;
}

#post-content :deep(ul) {
  @apply ml-1 md:ml-6 list-disc list-inside;
}
</style>
