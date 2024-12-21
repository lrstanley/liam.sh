<script setup lang="ts">
import { getGithubStats } from "@/utils/http/services.gen"

const { data: githubStats, suspense } = useQuery({
  queryKey: ["stats", "github"],
  queryFn: () => unwrapErrors(getGithubStats()),
})
onServerPrefetch(async () => {
  await suspense()
})

const gh = useGithubUser()
const year = new Date().getFullYear()
</script>

<template>
  <n-popover
    v-if="githubStats && gh"
    :width="275"
    raw
    :show-arrow="false"
    v-bind="$attrs"
    class="px-2 py-1 rounded border border-solid border-zinc-700 shadow-none !m-0"
  >
    <template #trigger>
      <a
        class="px-2 transition bg-blue-600 rounded-br-sm hover:bg-blue-800 hover:text-current"
        style="color: white !important"
        :href="gh.html_url"
      >
        <n-icon class="align-middle mt-[-3px] mr-[-7px]">
          <Icon name="mdi:github" />
        </n-icon>
        {{ gh.login }}
      </a>
    </template>

    <div class="flex flex-row items-center flex-auto mt-1">
      <n-avatar round :size="45" :src="gh.avatar_url + '&s=80'" class="mr-2 align-middle" />

      <div>
        <h2 class="text-base leading-tight">{{ gh.name }}</h2>
        <h2 class="text-base leading-tight text-zinc-400">{{ gh.login }}</h2>
      </div>
    </div>

    <!-- <div class="mt-4 text-sm">
      {{ gh.bio }}
    </div> -->

    <div class="mt-4">
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <Icon name="mdi:source-commit" class="text-zinc-400" />
          Total Commits ({{ year }})
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.commits_year.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <Icon name="mdi:heart-multiple-outline" class="text-zinc-400" />
          Contributed To ({{ year }})
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.contributed_repositories.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <Icon name="mdi:star-check-outline" class="text-zinc-400" />
          Total Stars Earned
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.stars.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <Icon name="mdi:source-pull" class="text-zinc-400" />
          Total PRs
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.pull_requests.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <Icon name="mdi:help-circle-outline" class="text-zinc-400" />
          Total Issues
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.all_issues.toLocaleString() }}
        </n-tag>
      </div>
    </div>
  </n-popover>
</template>
