<script setup lang="ts">
import { RepoObject, PostObject } from "#components"

type MaybeArray<T> = T | T[]
type Value = PostRead | GithubRepositoryRead

// TODO: use {} | {} syntax, but this isn't working here for some reason.
const props = defineProps<{
  value: MaybeArray<Value> | undefined
  type: "post" | "label" | "repo"
  showEmpty?: boolean
  divider?: boolean
  loading?: boolean
}>()

const objects = computed(() => {
  if (!props.value) {
    return []
  }
  const results: MappedObject[] = []

  if (Array.isArray(props.value)) {
    for (const item of props.value) {
      results.push(typeMapper(item))
    }
  } else {
    results.push(typeMapper(props.value))
  }

  return results
})

type MappedObject = {
  component: VNode
  object: Value
}

function typeMapper(o: Value): MappedObject {
  switch (props.type) {
    case "post":
      return { component: h(PostObject, { value: o as PostRead }), object: o }
    case "repo":
      return { component: h(RepoObject, { value: o as GithubRepositoryRead }), object: o }
    default:
      throw new Error("unknown type")
  }
}
</script>

<template>
  <div v-if="loading && objects.length < 1" class="flex flex-row gap-4">
    <!-- avatar -->
    <USkeleton class="rounded-full size-8" />

    <div class="flex flex-col w-full gap-2 my-1">
      <div class="flex flex-row">
        <!-- name -->
        <USkeleton class="h-6 w-[250px]" />
        <!-- homepage or similar link -->
        <USkeleton class="ml-auto h-6 w-[100px]" />
      </div>

      <!-- timestamp -->
      <USkeleton class="h-4 w-[350px] rounded-full" />

      <!-- main description -->
      <div class="flex flex-col gap-1 my-2">
        <USkeleton class="w-full h-4 rounded-full" />
        <USkeleton class="w-[80%] h-4 rounded-full" />
      </div>

      <div class="flex flex-row">
        <!-- tags -->
        <div class="flex flex-row flex-wrap gap-1">
          <USkeleton class="h-7 w-[100px]" />
          <USkeleton class="h-7 w-[100px]" />
          <USkeleton class="h-7 w-[100px]" />
          <USkeleton class="h-7 w-[100px]" />
          <USkeleton class="h-7 w-[100px]" />
          <USkeleton class="h-7 w-[100px]" />
          <USkeleton class="h-7 w-[100px]" />
          <USkeleton class="h-7 w-[100px]" />
        </div>

        <!-- views -->
        <USkeleton class="h-7 w-[100px]" />
      </div>
    </div>
  </div>
  <div v-else-if="showEmpty && objects.length < 1" class="flex flex-col items-center gap-4 mx-auto mb-4">
    <UIcon name="mdi:folder-remove-outline" class="text-5xl text-(--ui-primary)" />
    <span class="text-(--ui-text-muted)">No results found matching filters</span>
  </div>
  <!-- TODO: make sure v-auto-animate is working here -->
  <div v-auto-animate v-else-if="objects.length > 0" class="flex flex-col divide-y divide-zinc-500/20">
    <div v-for="object in objects" :key="object.object.id">
      <component :is="object.component" v-bind="$attrs" />
    </div>
  </div>
</template>
