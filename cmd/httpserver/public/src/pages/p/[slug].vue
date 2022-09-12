<template>
  <LayoutSidebar
    :loading="fetching"
    :error="error || (!post && `Post with ID '${props.slug}' not found`) || undefined"
    affix
    class="order-last md:order-first"
  >
    <n-back-top :visibility-height="600" class="z-[9999]" />

    <n-page-header class="container hidden mb-2 md:inline-flex">
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

        <div class="flex flex-col flex-auto mt-3 lg:flex-row lg:items-center mb-7 md:mb-12">
          <div class="inline-flex items-center">
            <n-avatar class="mr-3" round size="medium" :src="post.author.avatarURL" />
            <p>
              <a :href="post.author.htmlURL" target="_blank">{{ post.author.name }}</a>
              <br />
              <i>Published {{ useTimeAgo(post.publishedAt).value }}</i>
            </p>
          </div>
        </div>

        <div id="post-content" ref="postRef" class="lg:mb-[100px]" v-html="post.contentHTML" />
      </div>
    </div>

    <template #sidebar>
      <div v-if="state.base?.self" class="flex flex-col gap-1">
        <div class="text-emerald-500">Admin Options</div>

        <div>
          <PostViewCount :value="post.viewCount" />
          <n-tag v-if="!post.public" class="ml-2" type="warning">draft</n-tag>
        </div>

        <router-link :to="{ name: 'admin-edit-post-id', params: { id: post.id } }">
          <n-button type="success" tertiary size="small"> Edit post </n-button>
        </router-link>
      </div>

      <CoreTableOfContents :element="postRef" />

      <div>
        <div class="text-emerald-500">Post Labels</div>
        <div class="flex flex-wrap gap-1">
          <CoreObjectRender :value="post.labels" linkable />
        </div>
      </div>
    </template>
  </LayoutSidebar>
</template>

<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import { useGetPostContentQuery } from "@/lib/api"

const props = defineProps<{
  slug: string
}>()

const state = useState()
const { data, error, fetching } = useGetPostContentQuery({ variables: { slug: props.slug } })
const post = computed(() => data?.value?.posts?.edges[0]?.node)
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
  @apply max-w-[calc(100%)] lg:max-w-[calc(80%)] px-2 lg:px-0 !my-0 rounded-lg mx-auto;
  @apply border border-indigo-600/40 border-solid;
  height: auto;
}

#post-content :deep(p) {
  margin-top: 1em;
  margin-bottom: 1em;
  @apply text-center md:text-left;
}

#post-content :deep(a) {
  @apply !whitespace-pre-wrap break-words;
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
  margin-top: 1.5em;
}
#post-content :deep(h2) {
  font-size: 1.65em;
  margin-top: 1.3em;
}
#post-content :deep(h3) {
  font-size: 1.5em;
  margin-top: 1.2em;
}
#post-content :deep(h4) {
  font-size: 1.4em;
  margin-top: 0.75em;
}
#post-content :deep(h5) {
  font-size: 1.3em;
  margin-top: 0.1em;
}

#post-content :deep(h1),
#post-content :deep(h2),
#post-content :deep(h3),
#post-content :deep(h4),
#post-content :deep(h5) {
  @apply text-transparent bg-gradient-to-tr bg-clip-text font-bold;
  @apply bg-gradient-to-r from-sky-400 to-blue-500 ml-[10px];
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
