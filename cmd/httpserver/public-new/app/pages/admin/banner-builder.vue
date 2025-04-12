<script setup lang="ts">
import {} from "#hey-api/schemas.gen"

definePageMeta({
  title: "Banner Builder",
  layout: "admin",
})

const toast = useToast()

type Banner = Partial<
  GetGithubSvgData["query"] & {
    repo?: GithubRepository
  }
>

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
const input = ref<Banner>({ title: "", description: "", layout: undefined })
const repo = ref<string>("lrstanley/liam.sh")

const url = computed(() => {
  let base = `${getBackendURL()}/gh/svg`

  if (repo.value && !input.value.title) base += "/" + repo.value

  const params = new URLSearchParams()
  const raw: Record<string, any> = input.value

  for (const key in raw) {
    if (raw[key]) params.set(key, raw[key])
  }
  return params.size > 0 ? `${base}?${params.toString()}` : base
})
const urlDebounced = refDebounced(url, 250)

const [DefineSection, ReuseSection] = createReusableTemplate<{ title: string }>()

function copyURL() {
  navigator.clipboard.writeText(url.value)
  toast.add({
    color: "success",
    title: "Copied to clipboard",
    description: "The URL has been copied to your clipboard",
    duration: 2000,
  })
}
</script>

<template>
  <DefineSection v-slot="{ $slots, title }">
    <div class="flex flex-col gap-3">
      <UPageCard :title="title" variant="naked" />
      <UCard variant="subtle" class="p-3" :ui="{ body: 'flex flex-col gap-3' }">
        <component :is="$slots.default" />
      </UCard>
    </div>
  </DefineSection>

  <UContainer class="xl:w-[961px] sticky top-0 z-10">
    <UCard
      variant="outline"
      class="p-1 xl:p-3 flex xl:min-h-[200px] items-center justify-center drop-shadow-lg ring-primary-400/30"
    >
      <img :src="urlDebounced" />

      <UButton
        label="Copy URL"
        size="xs"
        color="primary"
        variant="outline"
        class="absolute bottom-4 right-4 opacity-50 hover:opacity-100 hover:cursor-pointer"
        @click="copyURL()"
      />
    </UCard>
  </UContainer>

  <UContainer class="flex flex-col gap-5 xl:w-[961px]">
    <ReuseSection title="Main Configuration">
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
          class="md:min-w-60"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Title"
        description="Banner title (vs repo name)"
        :required="!!input.title"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInput v-model="input.title" placeholder="Example Title" class="md:min-w-60" />
      </UFormField>

      <USeparator />

      <UFormField
        label="Description"
        description="Banner description"
        :required="!!input.title"
        :error="!!input.title && !input.description && 'required when title set'"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInput v-model="input.description" placeholder="The brown fox..." class="md:min-w-60" />
      </UFormField>
    </ReuseSection>

    <ReuseSection title="Dimensions & Layout">
      <UFormField
        label="Layout"
        description="Layout structure of the icons, title, description, etc"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <USelect
          v-model="input.layout"
          placeholder="Select layout"
          value-key="value"
          class="md:min-w-60"
          :items="[
            { label: 'default', value: undefined },
            { label: 'all', value: 'all' },
            { label: 'left', value: 'left' },
            { label: 'right', value: 'right' },
          ]"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Height"
        description="Height of the banner"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInputNumber
          @update:model-value="(v) => (input.h = v === 200 ? undefined : v)"
          class="md:min-w-60"
          :default-value="200"
          :min="200"
          :max="1000"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Width"
        description="Width of the banner"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInputNumber
          @update:model-value="(v) => (input.w = v === 961 ? undefined : v)"
          class="md:min-w-60"
          :default-value="961"
          :min="600"
          :max="2000"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Font Scale"
        description="Font scale of the banner"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInputNumber
          @update:model-value="(v) => (input.font = v === 1 ? undefined : v)"
          class="md:min-w-60"
          :default-value="1"
          :min="0.5"
          :max="3"
          :step="0.5"
          :format-options="{
            minimumFractionDigits: 1,
          }"
        />
      </UFormField>
    </ReuseSection>

    <ReuseSection title="Background">
      <UFormField
        label="Background Style"
        description="Background style of the banner"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <USelect
          v-model="input.bg"
          placeholder="Select background"
          value-key="value"
          class="md:min-w-60"
          :items="[
            { label: 'default', value: undefined },
            { label: 'geometric', value: 'geometric' },
            { label: 'topography', value: 'topography' },
          ]"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Background Color"
        description="For backgrounds which support it, the background color"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UPopover>
          <UButton label="Pick color" color="neutral" variant="outline">
            <template #leading>
              <span
                v-if="input.bgcolor"
                :style="{ 'background-color': input.bgcolor }"
                class="size-3 rounded-full"
              />
            </template>
          </UButton>

          <template #content>
            <UColorPicker @update:model-value="(v) => (input.bgcolor = v)" class="p-1" />

            <div class="flex flex-col items-center pb-1">
              <UButton
                label="Clear"
                color="neutral"
                variant="outline"
                @click="input.bgcolor = undefined"
              />
            </div>
          </template>
        </UPopover>
      </UFormField>
    </ReuseSection>

    <ReuseSection title="Icon Configuration">
      <UAlert color="primary" variant="subtle" orientation="horizontal">
        <template #description>
          This SVG generator uses icon references in the format of "category:icon-name". You can find the
          list of available icons on the
          <ULink href="https://icones.js.org/" target="_blank" external>icones.js.org</ULink>
          website.
        </template>
      </UAlert>

      <USeparator />

      <UFormField
        label="Icon"
        description="Icon to use for the banner"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInput
          @update:model-value="(v) => (input.icon = '' + v || undefined)"
          type="text"
          placeholder="category:icon-name"
          class="md:min-w-60"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Icon Height"
        description="Height of the icon (in px)"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInputNumber v-model="input['icon.height']" class="md:min-w-60" :min="0" :max="300" />
      </UFormField>

      <USeparator />

      <UFormField
        label="Icon Width"
        description="Width of the icon (in px)"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInputNumber v-model="input['icon.width']" class="md:min-w-60" :min="0" :max="300" />
      </UFormField>

      <USeparator />

      <UFormField
        label="Icon Flip"
        description="Flip the icon"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <USelect
          v-model="input['icon.flip']"
          placeholder="Select icon flip"
          value-key="value"
          class="md:min-w-60"
          :items="[
            { label: 'default', value: undefined },
            { label: 'horizontal', value: 'horizontal' },
            { label: 'vertical', value: 'vertical' },
          ]"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Icon Rotate"
        description="Rotate the icon (1=90deg, 2=180deg, 3=270deg)"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UInputNumber
          v-model="input['icon.rotate']"
          class="md:min-w-60"
          :min="0"
          :max="3"
          :step="1"
          :format-options="{
            minimumFractionDigits: 0,
          }"
        />
      </UFormField>

      <USeparator />

      <UFormField
        label="Icon Color"
        description="Color of the icon (hex, rgb, rgba, hsl)"
        class="flex max-sm:flex-col justify-between items-start gap-4"
      >
        <UPopover>
          <UButton label="Pick color" color="neutral" variant="outline">
            <template #leading>
              <span
                v-if="input['icon.color']"
                :style="{ 'background-color': input['icon.color'] }"
                class="size-3 rounded-full"
              />
            </template>
          </UButton>

          <template #content>
            <UColorPicker @update:model-value="(v) => (input['icon.color'] = v)" class="p-1" />

            <div class="flex flex-col items-center pb-1">
              <UButton
                label="Clear"
                color="neutral"
                variant="outline"
                @click="input['icon.color'] = undefined"
              />
            </div>
          </template>
        </UPopover>
      </UFormField>
    </ReuseSection>
  </UContainer>
</template>
