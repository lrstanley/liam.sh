<script setup lang="ts">
definePageMeta({
  layout: "default",
})

const { slug } = useRoute("p-slug").params

if (!slug || typeof slug !== "string") throw new Error("Invalid slug")

const { data: posts, error } = await useApi('/posts', {
  query: {
    "slug.eq": slug,
  },
})
if (error.value) throw error.value

const postRef = useTemplateRef("postRef")
const post = computed(() => (posts.value?.content ? posts.value.content[0] : null))
const toc = computed(() => (postRef.value ? getNodeTree(postRef.value) : []))
const self = useSelf()

useHead({
  title: post.value ? post.value?.title + " · Liam Stanley" : null,
})
</script>

<template>
  <ContainerStickySidebar v-if="post" class="mt-3 md:mt-8">
    <CoreTerminal class="hidden md:inline-flex" path="posts" prefix=""
      :value="'cat &quot;' + post.slug + '.md&quot;'" />

    <motion as="div" :initial="{ opacity: 0 }" :animate="{ opacity: 1 }" class="flex-col flex-auto mt-3">
      <div class="text-[30px]/12 md:text-[45px] text-gradient bg-linear-to-br from-pink-500 via-red-500 to-yellow-500">
        {{ post.title }}
      </div>

      <div class="flex flex-col flex-auto my-3 lg:flex-row lg:items-center">
        <div class="inline-flex items-center">
          <UAvatar class="mr-3" round size="lg" :src="post.edges.author.avatar_url + '&s=40'" />
          <p>
            <NuxtLink :href="post.edges.author.html_url" target="_blank">{{ post.edges.author.name }}</NuxtLink>
            <br />
            <i>Published {{ useTimeAgo(post.published_at).value }}</i>
          </p>
        </div>
      </div>

      <USeparator />
      <div id="post-content" ref="postRef" class="lg:mb-[100px]" v-html="post.content_html" />
    </motion>

    <template #sidebar>
      <div v-if="self">
        <motion as="div" :initial="{ opacity: 0, x: 20 }" :animate="{ opacity: 1, x: 0 }"
          class="mb-1 font-bold text-primary">
          Manage Post
        </motion>
        <motion as="div" :initial="{ opacity: 0, x: 20 }" :animate="{ opacity: 1, x: 0 }" class="mb-4">
          <UButton color="warning" variant="subtle" size="sm" icon="mdi:pencil"
            :to="{ name: 'admin-posts-id-edit', params: { id: post.id } }">
            Edit Post
          </UButton>
        </motion>
      </div>

      <CoreTableOfContents v-if="postRef" :links="toc" :element="postRef" />

      <div v-if="post.edges.labels && post.edges.labels.length > 0">
        <motion as="div" :initial="{ opacity: 0, x: 20 }" :animate="{ opacity: 1, x: 0 }"
          class="mb-1 font-bold text-primary">
          Post Labels
        </motion>
        <div class="flex flex-wrap gap-1">
          <motion as="div" :initial="{ opacity: 0, x: 20 }" :animate="{ opacity: 1, x: 0 }"
            v-for="(label, i) in post.edges.labels" :transition="{ delay: (i + 1) * 0.03 }" :key="label.id">
            <LabelObject :value="label" route="/posts" linkable />
          </motion>
        </div>
      </div>
    </template>
  </ContainerStickySidebar>
</template>

<style scoped>
@reference "../../assets/main.css";

#post-content {
  line-height: 1.55;
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
  @apply border border-(--ui-color-secondary-600)/40 border-solid;
  height: auto;
}

#post-content :deep(p) {
  margin-top: 1em;
  margin-bottom: 1em;
  @apply text-center md:text-left;
}

#post-content :deep(a) {
  @apply !whitespace-pre-wrap wrap-break-word;
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
  @apply border-l-primary-600 border-l-4 border-solid bg-(--ui-color-primary-600)/30 rounded;
}

#post-content :deep(blockquote)>p {
  margin-top: 0;
  margin-bottom: 0;
  text-indent: 0;
}

#post-content> :deep(p):not(blockquote)>strong {
  @apply text-success;
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
  @apply text-transparent bg-linear-to-tr bg-clip-text font-bold from-10% via-50% to-90%;
  @apply bg-linear-to-r from-info-400 to-info-600 ml-[10px];
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
  @apply text-primary;
}

#post-content :deep(:not(pre))>code {
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
