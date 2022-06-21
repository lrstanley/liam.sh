<template>
  <a :href="props.linkable ? repo.htmlURL : ''" target="_blank">
    <n-thing class="mb-7" content-indented v-bind="$attrs">
      <template #avatar>
        <n-avatar :src="repo.owner.avatarURL" class="hidden md:inline-flex" />
      </template>
      <template #header>
        <span class="repo-name text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500">
          {{ repo.owner.login == state.base.githubUser.login ? repo.name : repo.fullName }}
        </span>
      </template>
      <template #header-extra>
        <span class="inline-flex gap-1">
          <n-tag v-if="repo.owner.login != state.base.githubUser.login" type="info" size="small">
            maintainer
          </n-tag>
          <n-tag v-if="repo.fork" type="error" size="small">
            <template #icon>
              <n-icon><i-mdi-source-fork /></n-icon>
            </template>
            fork
          </n-tag>
          <n-tag v-if="repo.archived" type="warning" size="small">
            <template #icon>
              <n-icon><i-mdi-archive-outline /></n-icon>
            </template>
            archived
          </n-tag>
          <a v-if="repo.homepage" :href="repo.homepage" target="_blank">
            <n-tag type="success" size="small" class="cursor-pointer">
              <template #icon>
                <n-icon><i-mdi-link /></n-icon>
              </template>
              homepage
            </n-tag>
          </a>
          <n-tag v-if="repo.license" size="small" :title="repo.license.name">
            <template #icon>
              <n-icon><i-mdi-scale-balance /></n-icon>
            </template>
            {{ repo.license.key }}
          </n-tag>
        </span>
      </template>
      <template #description>
        <span>
          <i> Updated {{ useTimeAgo(repo.pushedAt).value }} </i>
        </span>
      </template>

      <span v-html="repo.description || 'No description available'" />

      <template v-if="repo.labels" #action>
        <div class="flex flex-auto justify-between">
          <div class="inline-flex flex-auto flex-wrap gap-1">
            <LabelObject
              v-for="label in repo.labels.edges.map(({ node }) => node)"
              :key="label.id"
              :value="label"
              route="repos"
              linkable
            />
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
</script>

<style scoped>
.repo-name {
  @apply text-size-1.4em md:text-size-1.5em;
}

div:deep(.n-thing-header),
div:deep(.n-thing-main__description) {
  @apply <md:(flex flex-col flex-wrap content-center items-center gap-2);
  /* display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  align-content: center;
  align-items: center;
  gap: 6px; */
}
</style>
