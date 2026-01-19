<script setup lang="ts">
import { EditorView } from "codemirror"
import { Codemirror } from "vue-codemirror"
import { markdown } from "@codemirror/lang-markdown"
import { oneDark } from "@codemirror/theme-one-dark"
import type { SchemaPostRead, SchemaPostCreate } from '#open-fetch-schemas/api'

const codeExtensions = [markdown(), oneDark, EditorView.lineWrapping]

const props = defineProps<{
  create: boolean
  loading?: boolean
  error?: string
}>()

const emit = defineEmits<{
  save: []
}>()

const post = defineModel<SchemaPostRead | SchemaPostCreate>({ required: true })
const labels = defineModel<number[]>("labels", { required: true })

const titleMinLength = 5
const titleMaxLength = 100
const slugMinLength = 5
const slugMaxLength = 50

function save() {
  emit("save")
}
</script>

<template>
  <div class="flex flex-row gap-2">
    <UButton color="secondary" variant="subtle" icon="lucide:arrow-left" :to="{ name: 'admin-posts' }"
      class="cursor-pointer">
      Back
    </UButton>
  </div>
  <ContainerStickySidebar>
    <UCard variant="subtle" class="p-1 xl:p-3 flex flex-col">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
        <UFormField label="Post title" required class="flex flex-col">
          <UInput v-model="post.title" placeholder="Post title" type="text" class="w-full" :minlength="titleMinLength"
            :maxlength="titleMaxLength">
            <template #trailing>
              <div class="text-xs text-muted tabular-nums" :class="{
                'text-error-400/50': (post.title?.length || 0) < titleMinLength,
              }" role="status">
                {{ post.title?.length || 0 }}/{{ titleMaxLength }}
              </div>
            </template>
          </UInput>
        </UFormField>

        <UFormField label="Post slug" required class="flex flex-col">
          <UInput v-model="post.slug" placeholder="Post slug" type="text" class="w-full" :minlength="slugMinLength"
            :maxlength="slugMaxLength">
            <template #trailing>
              <div class="text-xs text-muted tabular-nums" :class="{
                'text-error-400/50': (post.slug?.length || 0) < slugMinLength,
              }" role="status">
                {{ post.slug?.length || 0 }}/{{ slugMaxLength }}
              </div>
            </template>
          </UInput>
        </UFormField>
      </div>

      <div class="flex w-full shrink grow-0 mt-4 rounded">
        <codemirror v-model="post.content" placeholder="Post content" :autofocus="true"
          :style="{ 'min-height': '400px', width: '100%' }" class="rounded" :indent-with-tab="true" :tab-size="4"
          :extensions="codeExtensions" />
      </div>
    </UCard>

    <template #sidebar>
      <UCard variant="subtle" class="p-1 xl:p-3">
        <div class="flex flex-col gap-4">
          <UFormField label="Published at" class="flex flex-col" required>
            <UInput v-model="post.published_at" type="text" class="w-full" />
          </UFormField>

          <UFormField label="Public" class="flex flex-col" required>
            <USwitch v-model="post.public" class="w-full" />
            <div class="text-xs text-muted">
              <span class="font-bold">Public</span>
              posts are visible to all users
            </div>
          </UFormField>

          <UFormField label="Labels" class="flex flex-col" required>
            <LabelSelect v-model="labels" field="id" :suggest-text="post.content" allow-create allow-suggest />
          </UFormField>

          <UButton block color="primary" variant="subtle" @click="save()" :loading="loading" icon="mdi:content-save">
            {{ create ? "Create post" : "Save post" }}
          </UButton>

          <UAlert v-if="error && !loading" color="error">
            <template #description>
              <div v-html="error" />
            </template>
          </UAlert>
        </div>
      </UCard>
    </template>
  </ContainerStickySidebar>
</template>
