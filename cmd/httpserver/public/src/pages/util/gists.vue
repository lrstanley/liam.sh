<template>
  <LayoutTerminal>
    <n-alert v-if="error" type="error">{{ error.message }}</n-alert>
    <div v-else class="relative w-full h-full overflow-x-hidden grow basis-0">
      <TransitionGroup name="stepped" appear>
        <div
          v-for="(gist, i) in gists"
          :key="gist.id"
          :style="{ '--i': gists.length - i, '--total': gists.length }"
          class="flex flex-auto flex-row items-center gap-x-1 px-1 hover:bg-zinc-500/10 text-zinc-400 transition duration-75 ease-out border-b-[1px] border-b-gray-100"
        >
          <a :href="gist.owner.htmlurl" target="_blank">
            <n-avatar round :size="15" :src="gist.owner.avatarURL" class="mr-1 align-middle" />
          </a>

          <div class="flex items-center gap-2 truncate grow">
            <span class="text-zinc-300">{{ gist.name }}</span>
            <span class="text-violet-500">{{ gist.language || "Unknown" }}</span>
            <a target="_blank" class="text-emerald-500" :href="'/-/gh/g/' + gist.name">[open]</a>
            <a target="_blank" class="text-emerald-500" :href="gist.htmlURL">[gh]</a>
            <a target="_blank" class="text-amber-500 hover:text-amber-600" :href="gist.rawURL">[raw]</a>

            <span class="flex-1 lowercase truncate select-none text-zinc-600" :title="gist.description">
              {{ gist.description || "No description set" }}
            </span>
          </div>

          <div class="flex-none">
            <n-popover trigger="hover" style="padding: 2px 6px" :to="false" placement="left">
              <template #trigger>
                <i-mdi-clock-time-two-outline class="timestamp" />
              </template>

              {{ useTimeAgo(gist.updatedAt).value }}
            </n-popover>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </LayoutTerminal>
</template>

<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import { useGetGistsQuery } from "@/lib/api"

const { data, error } = useGetGistsQuery()

const gists = computed(() => data.value?.githubgists.edges.map((e) => e.node) ?? [])
</script>

<style scoped>
.timestamp {
  opacity: max(0.2, calc(var(--i) / var(--total)));
  @apply transition duration-100;
}

.timestamp:hover {
  opacity: 1;
}
</style>
