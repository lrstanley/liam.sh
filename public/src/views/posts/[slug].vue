<template>
  <LayoutDefault :loading="fetching" :error="error">
    <n-page-header class="mt-4 mb-8">
      <!-- <template #avatar>
        <n-icon :size="40"><i-mdi-pencil-outline /></n-icon>
      </template> -->
      <template #title>
        <!-- <a href="#" class="no-underline capitalize" style="color: inherit">
          Editing post #{{ post.id }}
        </a> -->
        <CoreTerminal
          class="mb-4 text-size-38px md:text-size-45px"
          path="posts"
          prefix=""
          :value="'cat &quot;' + post.slug + '.md&quot;'"
        />
      </template>
    </n-page-header>

    <div class="sm:container sm:mx-auto flex flex-auto flex-col flex-nowrap">
      <n-alert v-if="error" title="Error fetching post" type="error">
        {{ error }}
      </n-alert>

      <div v-motion-fade class="h-full flex flex-auto flex-col md:flex-row">
        <div class="h-full flex-auto mb-5 ml-7 md:ml-0 mr-7">
          <!-- <n-card> -->
          <n-space vertical>
            <div
              class="text-size-50px text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500"
            >
              {{ post.title }}
            </div>

            <div class="flex flex-auto flex-row flex-wrap items-center mt-3">
              <n-avatar class="mr-3" round size="medium" :src="post.author.avatarURL" />
              <p>
                <a :href="post.author.htmlURL" target="_blank">{{ post.author.name }}</a>
                <br />
                <i>Published {{ useTimeAgo(post.publishedAt).value }}</i>
              </p>
              <!-- <p class="ml-auto">
                <i>Published</i>
                <br />
                {{ useTimeAgo(post.publishedAt).value }}
              </p> -->
            </div>

            <div v-html="post.contentHTML" />
          </n-space>
          <!-- </n-card> -->
        </div>
        <div class="h-full mb-5 mx-7 md:mx-0 md:w-70">
          <!-- <n-card> -->
          <n-button block strong secondary type="primary">
            <n-icon class="mr-1"><i-mdi-content-save /></n-icon>
            Save post
          </n-button>
          <!-- </n-card> -->
        </div>
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

const { data, error, fetching } = useGetPostContentQuery({ variables: { slug: props.slug } })
const post = computed(() => data?.value?.posts?.edges[0].node)
</script>
