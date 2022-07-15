<template>
  <div v-motion-fade class="grid grid-sidebar gap-4 mb-20">
    <n-card size="small">
      <div class="grid gap-2 grid-cols-1 lg:grid-cols-2">
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

      <component
        :is="code"
        v-if="code && codeExtensions.length == 3"
        v-model="post.content"
        placeholder="Post content"
        :autofocus="true"
        :style="{ 'min-height': '400px' }"
        :indent-with-tab="true"
        :tab-size="4"
        :extensions="codeExtensions"
      />
    </n-card>
    <div>
      <n-card size="small" class="md:sticky md:top-5 md:left-0">
        <n-form-item
          label="Post published date"
          :required="true"
          class="flex flex-col flex-auto items-center"
        >
          <n-date-picker v-model:value="datetime" type="datetime" />
        </n-form-item>

        <LabelInput v-model="post.labelIDs" class="pb-5" :suggest="post.content" />

        <n-button block type="primary" @click="emit('update:post', post)">
          <n-icon class="mr-1"><i-mdi-content-save /></n-icon>
          Save post
        </n-button>
      </n-card>
    </div>
  </div>
</template>

<script setup>
const code = shallowRef(null)
import("vue-codemirror").then(({ Codemirror }) => {
  code.value = Codemirror
})

const codeExtensions = ref([])
import("@codemirror/theme-one-dark").then(({ oneDark }) => {
  codeExtensions.value.push(oneDark)
})
import("@codemirror/lang-markdown").then(({ markdown }) => {
  codeExtensions.value.push(markdown())
})
import("@codemirror/view").then(({ EditorView }) => {
  codeExtensions.value.push(EditorView.lineWrapping)
})

const props = defineProps({
  post: {
    type: Object,
    default: () => ({}),
  },
  create: {
    type: Boolean,
    default: false,
  },
})
const emit = defineEmits(["update:post"])

const post = reactive({
  publishedAt: new Date().toISOString(),
  ...props.post,
  labelIDs: props.post.labels?.edges?.map(({ node }) => node.id) ?? [],
})

const datetime = computed({
  get: () => Date.parse(post.publishedAt),
  set: (val) => {
    post.publishedAt = new Date(val).toISOString()
  },
})
</script>

<style scoped>
.grid-sidebar {
  @apply grid-cols-1 lg:grid-cols-[1fr,280px];
}
</style>
