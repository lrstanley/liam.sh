<script setup lang="ts">
definePageMeta({
  title: "Banner Builder",
  layout: "admin",
})

type Banner = GetGithubSvgData["query"] &
  GetGithubRepoSvgData["query"] & {
    repo?: GithubRepository
  }

const { data } = await listGithubRepositories({
  composable: "useFetch",
  query: {
    page: 1,
    per_page: 1000,
    "public.eq": true,
    "archived.eq": false,
  },
})

const repos = computed(() =>
  (data.value?.content ?? []).map((r) => r.full_name).sort((a, b) => a.localeCompare(b))
)
const input = ref<Banner>({} as Banner)
const repo = ref<string>("lrstanley/liam.sh")

const url = computed(() => {
  const params = new URLSearchParams()
  const raw: Record<string, any> = input.value

  for (const key in raw) {
    if (raw[key]) params.set(key, raw[key])
  }
  return `${getBackendURL()}/gh/svg${
    repo.value && !input.value.title ? "/" + repo.value : ""
  }?${params.toString()}`
})
const urlDebounced = refDebounced(url, 250)
</script>

<template>
  <UContainer class="flex flex-col gap-5 xl:w-[961px]">
    <UCard variant="subtle" class="p-3 flex xl:min-h-[200px] items-center justify-center">
      <img :src="urlDebounced" />
    </UCard>

    <div class="flex flex-col gap-3">
      <UPageCard title="Main Configuration" variant="naked" />
      <UCard
        variant="subtle"
        class="p-3"
        :ui="{
          body: 'flex flex-col gap-3',
        }"
      >
        <UFormField
          label="Repository"
          description="What repository to apply banner to"
          :required="!input.title"
          class="flex max-sm:flex-col justify-between items-start gap-4"
        >
          <USelectMenu
            v-model="repo"
            :search-input="{ placeholder: 'Filter...', icon: 'lucide:search' }"
            :items="repos"
            class="min-w-60"
          />
        </UFormField>

        <USeparator />

        <UFormField
          label="Title"
          description="Banner title (vs repo name)"
          :required="!!input.title"
          class="flex max-sm:flex-col justify-between items-start gap-4"
        >
          <UInput v-model="input.title" placeholder="Example Title" class="min-w-60" />
        </UFormField>

        <USeparator />

        <UFormField
          label="Description"
          description="Banner description"
          :required="!!input.title"
          :error="!!input.title && !input.description && 'required when title set'"
          class="flex max-sm:flex-col justify-between items-start gap-4"
        >
          <UInput v-model="input.description" placeholder="The brown fox..." class="min-w-60" />
        </UFormField>
      </UCard>
    </div>

    <div class="flex flex-col gap-3">
      <UPageCard title="Dimensions & Layout" variant="naked" />
      <UCard variant="subtle" class="p-3 flex">
        <div>TODO</div>
      </UCard>
    </div>

    <div class="flex flex-col gap-3">
      <UPageCard title="Background" variant="naked" />
      <UCard variant="subtle" class="p-3 flex">
        <div>TODO</div>
      </UCard>
    </div>

    <div class="flex flex-col gap-3">
      <UPageCard title="Icon Configuration" variant="naked" />
      <UCard variant="subtle" class="p-3 flex">
        <div>TODO</div>
      </UCard>
    </div>
  </UContainer>
</template>
