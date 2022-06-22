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
        <div class="inline-flex flex-auto flex-wrap gap-1">
          <LabelObject
            v-for="label in repo.labels.edges.map(({ node }) => node)"
            :key="label.id"
            :value="label"
            route="repos"
            linkable
            @click="drawerActive = false"
          />
        </div>
      </n-drawer-content>
    </n-drawer>

    <n-thing class="mb-7" content-indented v-bind="$attrs">
      <template #avatar>
        <n-avatar :src="repo.owner.avatarURL" />
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
          <i class="mr-auto"> Updated {{ useTimeAgo(repo.pushedAt).value }} </i>
          <RepoStatus :value="repo" class="inline-flex gap-1" />
        </span>
      </template>

      <span v-html="repo.description || 'No description available'" />

      <template v-if="repo.labels" #action>
        <div class="flex flex-auto justify-between">
          <div class="hidden md:inline-flex flex-auto flex-wrap gap-1">
            <LabelObject
              v-for="label in repo.labels.edges.map(({ node }) => node)"
              :key="label.id"
              :value="label"
              route="repos"
              linkable
            />
          </div>
          <div class="inline-flex md:hidden flex-auto flex-wrap gap-1">
            <n-button href="#" @click.prevent="drawerActive = true"> repo labels </n-button>
          </div>

          <n-tag class="text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500">
            {{ repo.starCount.toLocaleString() }}
            {{ repo.starCount === 1 ? "star" : "stars" }}
          </n-tag>
        </div>
      </template>
    </n-thing>
  </a>
</template>

<script setup>
import { useTimeAgo } from "@vueuse/core"

const props = defineProps({
  linkable: {
    type: Boolean,
    default: false,
  },
  value: {
    type: Object,
    required: true,
  },
})

const state = useState()
const repo = ref(props.value)

const drawerActive = ref(false)
</script>

<style scoped>
.repo-name {
  @apply text-size-1.4em md:text-size-1.5em truncate;
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

/* div:deep(.n-thing-header),
div:deep(.n-thing-main__description) {
  @apply <md:(flex flex-col flex-wrap content-center items-center gap-2);
} */
</style>
