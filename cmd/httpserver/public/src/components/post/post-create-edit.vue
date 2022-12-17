<template>
  <div v-motion-fade class="grid gap-4 mb-20 grid-sidebar">
    <n-card size="small">
      <div class="grid grid-cols-1 gap-2 lg:grid-cols-2">
        <n-form-item label="Post title" :required="true">
          <n-input
            v-model:value="post.title"
            :maxlength="100"
            :status="post.title?.length > 5 ? 'success' : 'error'"
            show-count
            type="text"
            placeholder="Post title"
          />
        </n-form-item>

        <n-form-item label="Post slug" :required="true">
          <n-input
            v-model:value="post.slug"
            :maxlength="50"
            :status="post.slug?.length > 5 ? 'success' : 'error'"
            show-count
            type="text"
            placeholder="Post slug"
          />
        </n-form-item>
      </div>

      <div class="flex w-full shrink grow-0">
        <codemirror
          v-model="post.content"
          placeholder="Post content"
          :autofocus="true"
          :style="{ 'min-height': '400px', width: '100%' }"
          :indent-with-tab="true"
          :tab-size="4"
          :extensions="codeExtensions"
        />
      </div>
    </n-card>
    <div>
      <n-card size="small" class="md:sticky md:top-5 md:left-0">
        <div class="flex flex-col gap-3">
          <div>
            <span class="text-emerald-400"> Post published date </span>
            <n-date-picker v-model:value="datetime" type="datetime" />
          </div>

          <div>
            <div class="text-emerald-400">Post attributes</div>
            <n-checkbox v-model:checked="post.public"> Public </n-checkbox>
          </div>

          <LabelInput v-model="labelIDs" :suggest="post.content" />

          <n-button block type="primary" @click="emit('update:post', post, labelIDs)">
            <n-icon class="mr-1"><i-mdi-content-save /></n-icon>
            Save post
          </n-button>
        </div>
      </n-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { EditorView } from "codemirror"
import { Codemirror } from "vue-codemirror"
import { markdown } from "@codemirror/lang-markdown"
import { oneDark } from "@codemirror/theme-one-dark"
import type { Post } from "@/lib/api"

const codeExtensions = [markdown(), oneDark, EditorView.lineWrapping]

const props = defineProps<{
  post?: Post
  create?: boolean
}>()
const emit = defineEmits(["update:post"])

const post = ref<Post>(props.post ?? ({} as Post))
const labelIDs = ref<string[]>(props.post?.labels?.edges?.map(({ node }) => node.id) ?? [])

const datetime = computed({
  get: () => Date.parse(post.value.publishedAt ?? new Date()),
  set: (val) => {
    post.value.publishedAt = new Date(val).toISOString()
  },
})
</script>

<style scoped>
.grid-sidebar {
  @apply grid-cols-1 lg:grid-cols-[1fr,280px];
}
</style>
