<script setup lang="ts">
import { getGithubStats } from "@/utils/http/sdk.gen"

const githubStats = await getGithubStats({ composable: "$fetch" })

const gh = useGithubUser()
const year = new Date().getFullYear()
</script>

<template>
  <UPopover
    v-if="githubStats && gh"
    mode="hover"
    :popper="{ placement: 'top' }"
    v-bind="$attrs"
    :ui="{
      background: 'dark:bg-zinc-800',
      rounded: 'rounded',
      ring: 'ring-zinc-700',
      shadow: 'shadow-none',
      trigger: 'flex',
    }"
  >
    <a
      class="px-2 transition bg-blue-600 rounded-br-sm hover:bg-blue-800 hover:text-current"
      style="color: white !important"
      :href="gh.html_url"
    >
      <UIcon name="mdi:github" class="align-middle mt-[-3px] mr-[-7px]" />
      {{ gh.login }}
    </a>

    <template #panel>
      <div class="p-1">
        <div class="flex flex-row items-center flex-auto mt-1">
          <UAvatar size="md" :src="gh.avatar_url + '&s=80'" class="mr-2 align-middle" :alt="gh.login" />

          <div>
            <h2 class="text-base leading-tight">{{ gh.name }}</h2>
            <h2 class="text-base leading-tight text-zinc-400">{{ gh.login }}</h2>
          </div>
        </div>

        <div class="mt-4">
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:source-commit" class="text-zinc-400" />
              Total Commits ({{ year }})
            </span>
            <UBadge class="ml-1" color="emerald" variant="soft">
              {{ githubStats.commits_year.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:heart-multiple-outline" class="text-zinc-400" />
              Contributed To ({{ year }})
            </span>
            <UBadge class="ml-1" color="emerald" variant="soft">
              {{ githubStats.contributed_repositories.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:star-check-outline" class="text-zinc-400" />
              Total Stars Earned
            </span>
            <UBadge class="ml-1" color="emerald" variant="soft">
              {{ githubStats.stars.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:source-pull" class="text-zinc-400" />
              Total PRs
            </span>
            <UBadge class="ml-1" color="emerald" variant="soft">
              {{ githubStats.pull_requests.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:help-circle-outline" class="text-zinc-400" />
              Total Issues
            </span>
            <UBadge class="ml-1" color="emerald" variant="soft">
              {{ githubStats.all_issues.toLocaleString() }}
            </UBadge>
          </div>
        </div>
      </div>
    </template>
  </UPopover>
</template>
