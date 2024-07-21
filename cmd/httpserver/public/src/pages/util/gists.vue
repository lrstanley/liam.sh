<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import { listGithubGists } from "@/lib/http/services.gen"

definePage({
  meta: {
    title: "Gists",
    layout: "terminal",
  },
})

const {
  data: gists,
  error,
  suspense,
} = useQuery({
  queryKey: ["gists"],
  queryFn: () =>
    unwrapErrors(
      listGithubGists({
        query: {
          page: 1,
          per_page: 250,
          "public.eq": true,
          sort: "name",
          foo: "bar",
        },
      })
    ),
})
await suspense()

watch(error, () => {
  if (error.value) throw error.value
})
</script>

<template>
  <ContainerIde>
    <div class="relative overflow-x-hidden size-full grow basis-0">
      <TransitionGroup name="stepped" appear>
        <div
          v-for="(gist, i) in gists?.content"
          :key="gist.id"
          :style="{ '--i': gists.content.length - i, '--total': gists.content.length }"
          class="flex flex-row items-center flex-auto px-1 transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-zinc-400 border-b-DEFAULT border-b-gray-100"
        >
          <a :href="gist.owner.html_url as string" target="_blank">
            <n-avatar
              round
              :size="15"
              :src="gist.owner.avatar_url + '&s=40'"
              class="mr-1 align-middle"
            />
          </a>

          <div class="flex items-center gap-2 truncate grow">
            <span class="text-zinc-300">{{ gist.name }}</span>
            <span class="text-violet-500">{{ gist.language || "Unknown" }}</span>
            <a target="_blank" class="text-emerald-500" :href="'/-/gh/g/' + gist.name">[open]</a>
            <a target="_blank" class="text-emerald-500" :href="gist.html_url">[gh]</a>
            <a target="_blank" class="text-amber-500 hover:text-amber-600" :href="gist.raw_url">[raw]</a>

            <span class="flex-1 lowercase truncate select-none text-zinc-600" :title="gist.description">
              {{ gist.description || "No description set" }}
            </span>
          </div>

          <div class="flex-none">
            <n-popover trigger="hover" style="padding: 2px 6px" :to="false" placement="left">
              <template #trigger>
                <i-mdi-clock-time-two-outline class="timestamp" />
              </template>

              {{ useTimeAgo(gist.updated_at).value }}
            </n-popover>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </ContainerIde>
</template>

<style scoped>
.timestamp {
  opacity: max(0.2, calc(var(--i) / var(--total)));
  @apply transition duration-100;
}

.timestamp:hover {
  opacity: 1;
}
</style>
