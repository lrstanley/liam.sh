<script setup lang="ts">
import { getGithubStats } from "@/lib/http/services.gen"

const { data: githubStats } = useQuery({
  queryKey: ["stats", "github"],
  queryFn: () => unwrapErrors(getGithubStats()),
})

const state = useState()
const year = new Date().getFullYear()
</script>

<template>
  <n-popover
    v-if="githubStats && state.githubUser"
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
        :href="state.githubUser.html_url"
      >
        <n-icon class="align-middle mt-[-3px] mr-[-7px]">
          <i-mdi-github />
        </n-icon>
        {{ state.githubUser.login }}
      </a>
    </template>

    <div class="flex flex-row items-center flex-auto mt-1">
      <n-avatar
        round
        :size="45"
        :src="state.githubUser.avatar_url + '&s=80'"
        class="mr-2 align-middle"
      />

      <div>
        <h2 class="text-base leading-tight">{{ state.githubUser.name }}</h2>
        <h2 class="text-base leading-tight text-zinc-400">{{ state.githubUser.login }}</h2>
      </div>
    </div>

    <!-- <div class="mt-4 text-sm">
      {{ state.githubUser.bio }}
    </div> -->

    <div class="mt-4">
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <i-mdi-source-commit class="text-zinc-400" />
          Total Commits ({{ year }})
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.commits_year.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <i-mdi-heart-multiple-outline class="text-zinc-400" />
          Contributed To ({{ year }})
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.contributed_repositories.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <i-mdi-star-check-outline class="text-zinc-400" />
          Total Stars Earned
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.stars.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <i-mdi-source-pull class="text-zinc-400" />
          Total PRs
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.pull_requests.toLocaleString() }}
        </n-tag>
      </div>
      <div class="flex flex-row justify-between flex-auto mt-1">
        <span>
          <i-mdi-help-circle-outline class="text-zinc-400" />
          Total Issues
        </span>
        <n-tag type="success" :bordered="false" size="small">
          {{ githubStats.all_issues.toLocaleString() }}
        </n-tag>
      </div>
    </div>
  </n-popover>
</template>
