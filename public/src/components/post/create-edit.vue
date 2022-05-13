<template>
  <div class="h-full flex flex-auto flex-col md:flex-row">
    <div class="h-full flex-auto mb-5 ml-7 md:ml-0 mr-7">
      <n-card>
        <n-space vertical>
          <div class="flex flex-auto flex-col md:flex-row">
            <n-form-item label="Post title" :required="true" class="flex-auto mr-2">
              <n-input
                v-model:value="post.title"
                :maxlength="100"
                :status="post.title?.length > 5 ? 'success' : 'error'"
                show-count
                type="text"
                placeholder="Post title"
              />
            </n-form-item>

            <n-form-item label="Post slug" :required="true" class="flex-auto">
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

          <n-form-item label="Post content" :required="true">
            <n-input
              v-model:value="post.content"
              :status="post.content?.length > 5 ? 'success' : 'error'"
              type="textarea"
              placeholder="Post content"
              :autosize="{ minRows: 5, maxRows: 16 }"
            />
          </n-form-item>
        </n-space>
      </n-card>
    </div>
    <div class="h-full mb-5 mx-7 md:mx-0 md:w-70">
      <n-card>
        <n-form-item
          label="Post published date"
          :required="true"
          class="flex flex-col flex-auto items-center"
        >
          <n-date-picker v-model:value="datetime" type="datetime" />
        </n-form-item>

        <n-button block strong secondary type="primary" @click="emit('update:post', post)">
          <n-icon class="mr-1"><i-mdi-content-save /></n-icon>
          Save post
        </n-button>
      </n-card>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  post: {
    type: Object,
    default: () => ({}),
  },
})
const emit = defineEmits(["update:post"])

const post = reactive({
  publishedAt: new Date().toISOString(),
  ...props.post,
})

const datetime = computed({
  get: () => Date.parse(post.publishedAt),
  set: (val) => {
    post.publishedAt = new Date(val).toISOString()
  },
})
</script>
