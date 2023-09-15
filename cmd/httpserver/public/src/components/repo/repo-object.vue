<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import type { GithubRepository } from "@/lib/api"

const props = defineProps<{
  value: GithubRepository
  linkable?: boolean
}>()

const state = useState()
const repo = ref(props.value)

const drawerActive = ref(false)
</script>

<template>
  <a :href="props.linkable ? repo.htmlURL : ''" target="_blank">
    <n-drawer
      v-model:show="drawerActive"
      :auto-focus="false"
      :height="200"
      placement="bottom"
      class="bg-dark-400"
    >
      <n-drawer-content title="repo labels">
        <div class="inline-flex flex-wrap flex-auto gap-1">
          <LabelObject
            v-for="label in repo.labels.edges.map(({ node }) => node)"
            :key="label.id"
            :value="label"
            route="/repos"
            linkable
            @click="drawerActive = false"
          />
        </div>
      </n-drawer-content>
    </n-drawer>

    <n-thing class="mb-7" content-indented v-bind="$attrs">
      <template #avatar>
        <n-avatar :src="repo.owner.avatarURL + '&s=40'" round />
      </template>
      <template #header>
        <div class="repo-name text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500">
          {{ repo.owner.login == state.base.githubUser.login ? repo.name : repo.fullName }}
        </div>
      </template>
      <template #header-extra>
        <a v-if="repo.homepage" :href="repo.homepage" target="_blank">
          <n-tag type="success" size="small" class="cursor-pointer">
            <template #icon>
              <n-icon><i-mdi-link /></n-icon>
            </template>
            homepage
          </n-tag>
        </a>
      </template>
      <template #description>
        <span class="flex gap-1">
          <n-popover>
            <template #trigger>
              <span class="inline-flex">
                <n-icon class="mr-1 text-purple-400">
                  <i-mdi-update />
                </n-icon>
                <i class="text-zinc-400">
                  {{ useTimeAgo(repo.pushedAt).value }}
                </i>
              </span>
            </template>
            last updated
          </n-popover>

          <n-popover>
            <template #trigger>
              <span class="hidden mr-auto md:inline-flex">
                <n-icon class="mr-1 text-lime-400 ml-1ch">
                  <i-mdi-rocket-launch-outline />
                </n-icon>
                <i class="text-zinc-400">
                  {{ useTimeAgo(repo.createdAt).value }}
                </i>
              </span>
            </template>
            created
          </n-popover>
          <RepoStatus :value="repo" class="inline-flex gap-1 ml-auto" />
        </span>
      </template>

      <span v-html="repo.description || 'No description available'" />

      <template v-if="repo.labels" #action>
        <div class="flex justify-between flex-auto">
          <div class="flex-wrap flex-auto hidden gap-1 md:inline-flex">
            <LabelObject
              v-for="label in repo.labels.edges.map(({ node }) => node)"
              :key="label.id"
              :value="label"
              route="/repos"
              linkable
            />
          </div>
          <div class="inline-flex flex-wrap flex-auto gap-1 md:hidden">
            <n-button href="#" @click.prevent="drawerActive = true">repo labels</n-button>
          </div>

          <n-tag class="ml-3 text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500">
            {{ repo.starCount.toLocaleString() }}
            {{ repo.starCount === 1 ? "star" : "stars" }}
          </n-tag>
        </div>
      </template>
    </n-thing>
  </a>
</template>

<style scoped>
.repo-name {
  @apply text-[1.4em] md:text-[1.5em] truncate;
}

.n-thing,
.n-thing :deep(.n-thing-main),
.n-thing :deep(.n-thing-header),
.n-thing :deep(.n-thing-header__title) {
  display: flex;
  flex-grow: 1;
  flex-shrink: 1;
  flex-basis: 0;
  width: 100%;
}

.n-thing :deep(.n-thing-avatar) {
  @apply hidden md:inline-flex;
}

.n-thing :deep(.n-thing-main) {
  flex-direction: column;
}
</style>
