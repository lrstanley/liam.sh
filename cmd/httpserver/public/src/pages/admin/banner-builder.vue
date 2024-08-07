<script setup lang="ts">
import { listGithubRepositories } from "@/lib/http/services.gen"
import type { GetGithubSvgData, GetGithubRepoSvgData, GithubRepository } from "@/lib/http/types.gen"
import { client } from "@hey-api/client-fetch"

const { data: repositories } = useQuery({
  queryKey: ["repositories", "banners"],
  queryFn: () =>
    unwrapErrors(
      listGithubRepositories({
        query: {
          page: 1,
          per_page: 1000,
          "public.eq": true,
          "archived.eq": false,
        },
      })
    ),
})

definePage({
  meta: {
    layout: "admin",
  },
})

type Banner = GetGithubSvgData["query"] &
  GetGithubRepoSvgData["query"] & {
    repo?: GithubRepository
  }

const input = ref<Banner>({} as Banner)
const repo = ref<string>("")

const repos = computed(
  () =>
    repositories.value?.content.map((repo) => ({
      label: repo.full_name + ": " + repo.description,
      value: repo.full_name,
    })) ?? []
)

const url = computed(() => {
  const params = new URLSearchParams()
  for (const key in input.value) {
    if (input.value[key]) params.set(key, input.value[key])
  }
  return `${client.getConfig().baseUrl}/gh/svg${repo.value ? "/" + repo.value : ""}?${params.toString()}`
})
const urlDebounced = refDebounced(url, 250)
</script>

<template>
  <div class="p-4 sm:container sm:mx-auto">
    <n-page-header
      subtitle="Build beautiful Github repository banners"
      class="hidden mt-4 mb-8 lg:block"
    >
      <template #avatar>
        <n-icon :size="40"><i-mdi-image-edit-outline /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">Banner Builder</a>
      </template>
    </n-page-header>

    <n-card class="mb-10">
      <div class="flex justify-center flex-auto">
        <a :href="urlDebounced" target="_blank">
          <img :src="urlDebounced" />
        </a>
      </div>

      <n-divider class="!mt-3 !mb-2">Main Configuration</n-divider>

      <div class="grid grid-cols-3 gap-4">
        <n-form-item label="Repository">
          <n-select
            v-model:value="repo"
            placeholder="Select repository"
            filterable
            clearable
            :options="repos"
          />
        </n-form-item>
        <n-form-item label="Title">
          <n-input v-model:value="input.title" placeholder="Banner title" :disabled="!!repo" />
        </n-form-item>
        <n-form-item label="Description">
          <n-input
            v-model:value="input.description"
            placeholder="Banner description"
            :disabled="!!repo"
          />
        </n-form-item>
      </div>

      <n-divider class="!mt-3 !mb-2">Dimensions & Layout</n-divider>

      <div class="grid grid-cols-4 gap-4">
        <n-form-item label="Layout">
          <n-radio-group v-model:value="input.layout" name="layout">
            <n-space>
              <n-radio value="">default</n-radio>
              <n-radio value="all">all (both)</n-radio>
              <n-radio value="left">left</n-radio>
              <n-radio value="right">right</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="Height">
          <n-input-number v-model:value="input.h" />
        </n-form-item>
        <n-form-item label="Width">
          <n-input-number v-model:value="input.w" />
        </n-form-item>
        <n-form-item label="Font Scale">
          <n-input-number v-model:value="input.font" />
        </n-form-item>
      </div>

      <n-divider class="!mt-3 !mb-2">Background</n-divider>

      <div class="grid grid-cols-3 gap-4">
        <n-form-item label="Background">
          <n-radio-group v-model:value="input.bg" name="background">
            <n-space>
              <n-radio value="">default</n-radio>
              <n-radio value="geometric">geometric</n-radio>
              <n-radio value="topography">topography</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="Background Color">
          <n-color-picker v-model:value="input.bgcolor" />
        </n-form-item>
      </div>

      <n-divider class="!mt-3 !mb-2">
        Icon Configuration --
        <a href="https://icones.js.org/" class="ml-[1ch]" target="_blank">icon ref</a>
      </n-divider>

      <div class="grid grid-cols-3 gap-4">
        <n-form-item label="Icon">
          <n-input v-model:value="input.icon" placeholder="<category>:<icon>" />
        </n-form-item>

        <n-form-item label="Icon Height">
          <n-input-number v-model:value="input['icon.height']" />
        </n-form-item>
        <n-form-item label="Icon Width">
          <n-input-number v-model:value="input['icon.width']" />
        </n-form-item>

        <n-form-item label="Icon Flip">
          <n-radio-group v-model:value="input['icon.flip']" name="icon.flip">
            <n-space>
              <n-radio value="">default</n-radio>
              <n-radio value="horizontal">horizontal</n-radio>
              <n-radio value="vertical">vertical</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="Icon Rotation">
          <n-input-number v-model:value="input['icon.rotate']" />
        </n-form-item>

        <n-form-item label="Icon Color">
          <n-color-picker v-model:value="input['icon.color']" />
        </n-form-item>
      </div>
    </n-card>
  </div>
</template>
