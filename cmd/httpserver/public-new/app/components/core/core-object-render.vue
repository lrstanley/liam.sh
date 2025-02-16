<script setup lang="ts">
import { h } from "vue"
import PostObject from "@/components/post/post-object.vue"
import LabelObject from "@/components/label/label-object.vue"
import RepoObject from "@/components/repo/repo-object.vue"
import type { PostRead, LabelRead, GithubRepositoryRead } from "@/utils/http/types.gen"

type MaybeArray<T> = T | T[]
type Value = PostRead | LabelRead | GithubRepositoryRead

// TODO: use {} | {} syntax, but this isn't working here for some reason.
const props = defineProps<{
  value: MaybeArray<Value>
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
    case "label":
      return { component: h(LabelObject, { value: o as LabelRead }), object: o }
    case "repo":
      return { component: h(RepoObject, { value: o as GithubRepositoryRead }), object: o }
    default:
      throw new Error("unknown type")
  }
}
</script>

<template>
  <div v-if="loading" class="flex flex-row gap-2 mb-4">
    <!-- avatar -->
    <n-skeleton height="40px" circle />

    <div class="flex flex-col gap-3 grow">
      <div class="flex">
        <!-- name -->
        <n-skeleton height="20px" round width="20%" />
        <!-- homepage or similar -->
        <n-skeleton height="20px" :sharp="false" width="15%" class="ml-auto" />
      </div>

      <div class="flex">
        <!-- timestamp -->
        <n-skeleton height="15px" round width="33%" />
        <!-- license -->
        <n-skeleton height="20px" :sharp="false" width="10%" class="ml-auto" />
      </div>

      <!-- main text -->
      <n-skeleton height="30px" round width="75%" />

      <div class="flex">
        <!-- tags -->
        <n-skeleton height="30px" width="70%" :sharp="false" />
        <!-- views -->
        <n-skeleton height="30px" width="10%" class="ml-auto" />
      </div>
    </div>
  </div>
  <n-empty v-else-if="showEmpty && objects.length < 1" description="No items found matching filters" />
  <TransitionGroup v-else-if="objects.length > 0" appear name="fade">
    <div
      v-for="(object, i) in objects"
      :key="object.object.id"
      :style="{ '--i': i, '--total': objects.length }"
    >
      <component :is="object.component" v-bind="$attrs" />
      <n-divider v-if="divider && i != objects.length - 1" />
    </div>
  </TransitionGroup>
</template>

<style scoped>
.fade-move,
.fade-enter-active,
.fade-leave-active {
  transition: all 0.1s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
