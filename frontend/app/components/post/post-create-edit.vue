<script setup lang="ts">
import type { EditorToolbarItem, EditorSuggestionMenuItem } from "@nuxt/ui"
import { CodeBlockShiki } from "tiptap-extension-code-block-shiki"
import type { SchemaPostRead, SchemaPostCreate } from "#open-fetch-schemas/api"

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

const content = computed({
  get: () => post.value.content ?? "",
  set: (v) => {
    post.value.content = v
  },
})

const titleMinLength = 5
const titleMaxLength = 100
const slugMinLength = 5
const slugMaxLength = 50

function save() {
  emit("save")
}

const editorExtensions = [
  CodeBlockShiki.configure({
    defaultTheme: "material-theme",
    themes: {
      light: "material-theme-lighter",
      dark: "material-theme-palenight",
    },
  }),
]

// SSR-safe function to append menus to body (avoids z-index issues in docs)
const appendToBody = import.meta.client ? () => document.body : undefined

const toolbarItems: EditorToolbarItem[][] = [[
  {
    icon: "i-lucide-heading",
    tooltip: { text: "Headings" },
    content: { align: "start" },
    items: [
      { kind: "heading", level: 1, icon: "i-lucide-heading-1", label: "Heading 1" },
      { kind: "heading", level: 2, icon: "i-lucide-heading-2", label: "Heading 2" },
      { kind: "heading", level: 3, icon: "i-lucide-heading-3", label: "Heading 3" },
      { kind: "heading", level: 4, icon: "i-lucide-heading-4", label: "Heading 4" },
    ],
  },
  {
    icon: "i-lucide-list",
    tooltip: { text: "Lists" },
    content: { align: "start" },
    items: [
      { kind: "bulletList", icon: "i-lucide-list", label: "Bullet List" },
      { kind: "orderedList", icon: "i-lucide-list-ordered", label: "Ordered List" },
    ],
  },
  { kind: "blockquote", icon: "i-lucide-text-quote", tooltip: { text: "Blockquote" } },
  { kind: "codeBlock", icon: "i-lucide-square-code", tooltip: { text: "Code Block" } },
], [
  { kind: "mark", mark: "bold", icon: "i-lucide-bold", tooltip: { text: "Bold" } },
  { kind: "mark", mark: "italic", icon: "i-lucide-italic", tooltip: { text: "Italic" } },
  { kind: "mark", mark: "underline", icon: "i-lucide-underline", tooltip: { text: "Underline" } },
  { kind: "mark", mark: "strike", icon: "i-lucide-strikethrough", tooltip: { text: "Strikethrough" } },
  { kind: "mark", mark: "code", icon: "i-lucide-code", tooltip: { text: "Code" } },
  { kind: "link", icon: "i-lucide-link", tooltip: { text: "Link" } },
]]

const plainTextMode = ref(false)

const suggestionItems: EditorSuggestionMenuItem[][] = [[
  { type: "label", label: "Style" },
  { kind: "paragraph", label: "Paragraph", icon: "i-lucide-type" },
  { kind: "heading", level: 1, label: "Heading 1", icon: "i-lucide-heading-1" },
  { kind: "heading", level: 2, label: "Heading 2", icon: "i-lucide-heading-2" },
  { kind: "heading", level: 3, label: "Heading 3", icon: "i-lucide-heading-3" },
  { kind: "bulletList", label: "Bullet List", icon: "i-lucide-list" },
  { kind: "orderedList", label: "Numbered List", icon: "i-lucide-list-ordered" },
  { kind: "blockquote", label: "Blockquote", icon: "i-lucide-text-quote" },
  { kind: "codeBlock", label: "Code Block", icon: "i-lucide-square-code" },
], [
  { type: "label", label: "Insert" },
  { kind: "horizontalRule", label: "Horizontal Rule", icon: "i-lucide-separator-horizontal" },
]]

</script>

<template>
  <div class="flex flex-row gap-2 py-0.5">
    <UButton color="secondary" variant="subtle" size="sm" icon="lucide:arrow-left" :to="{ name: 'admin-posts' }"
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

      <div class="flex w-full mt-4 rounded bg-zinc-900 min-h-[400px] max-h-[70vh] overflow-y-auto">
        <UEditor v-if="!plainTextMode" v-slot="{ editor }" v-model="content" content-type="markdown"
          :starter-kit="{ codeBlock: false }" :extensions="editorExtensions"
          placeholder="Post content (type / for commands)" autofocus
          class="post-editor w-full min-h-[400px] rounded" :ui="{ base: 'p-4' }">
          <UEditorToolbar :editor="editor" :items="toolbarItems" layout="bubble" :append-to="appendToBody"
            :ui="{ base: 'z-9999' }" />
          <UEditorSuggestionMenu :editor="editor" :items="suggestionItems" :append-to="appendToBody" />
          <UEditorDragHandle :editor="editor" />
        </UEditor>
        <div v-else class="relative flex-1 flex flex-col min-h-[400px] w-full">
          <textarea v-model="content" placeholder="Post content (markdown)" autofocus
            class="absolute inset-0 w-full min-h-[400px] rounded p-4 font-mono text-sm resize-none border-0 bg-transparent placeholder:text-dimmed focus:outline-none focus:ring-0" />
        </div>
      </div>
    </UCard>

    <template #sidebar>
      <UCard variant="subtle" class="p-1 xl:p-3">
        <div class="flex flex-col gap-4">
          <UFormField label="Published at" class="flex flex-col" required>
            <UInput v-model="post.published_at" type="text" class="w-full" />
          </UFormField>

          <UFormField label="Plain text" class="flex flex-col">
            <USwitch v-model="plainTextMode" class="w-full" />
            <div class="text-xs text-muted">
              edit raw markdown instead of rich editor
            </div>
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

<style>
html.dark .tiptap .shiki,
html.dark .tiptap .shiki span {
  color: var(--shiki-dark) !important;
  background-color: var(--ui-bg-muted) !important;
}
</style>
