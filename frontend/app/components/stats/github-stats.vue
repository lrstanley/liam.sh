<script setup lang="ts">
const githubStats = await getGithubStats({ composable: "$fetch" })

const gh = useGithubUser()
const year = new Date().getFullYear()
const { width } = useWindowSize()

const small = computed(() => width.value <= 640)
</script>

<template>
  <UPopover
    v-if="githubStats && gh"
    :mode="small ? 'click' : 'hover'"
    :open-delay="100"
    :content="{ side: 'top', align: 'end', sideOffset: 0 }"
    v-bind="$attrs"
  >
    <div
      class="px-2 text-(--ui-text-highlighted) transition hover:text-current inline-flex items-center cursor-pointer"
    >
      <UIcon name="mdi:github" class="align-middle size-4 max-sm:text-lg mr-1" />
      <span class="max-sm:text-sm">{{ gh.login }}</span>
    </div>

    <template #content>
      <div class="p-1">
        <a class="flex flex-row items-center flex-auto mt-1" :href="gh.html_url" target="_blank">
          <UAvatar size="md" :src="gh.avatar_url + '&s=80'" class="mr-2 align-middle" :alt="gh.login" />

          <div>
            <h2 class="text-base leading-tight">{{ gh.name }}</h2>
            <h2 class="text-base leading-tight text-(--ui-text-muted)">{{ gh.login }}</h2>
          </div>
        </a>

        <div class="mt-4">
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:source-commit" class="text-(--ui-text-muted)" />
              Total Commits ({{ year }})
            </span>
            <UBadge class="ml-1" color="primary" variant="soft">
              {{ githubStats.commits_year.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:heart-multiple-outline" class="text-(--ui-text-muted)" />
              Contributed To ({{ year }})
            </span>
            <UBadge class="ml-1" color="primary" variant="soft">
              {{ githubStats.contributed_repositories.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:star-check-outline" class="text-(--ui-text-muted)" />
              Total Stars Earned
            </span>
            <UBadge class="ml-1" color="primary" variant="soft">
              {{ githubStats.stars.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:source-pull" class="text-(--ui-text-muted)" />
              Total PRs
            </span>
            <UBadge class="ml-1" color="primary" variant="soft">
              {{ githubStats.pull_requests.toLocaleString() }}
            </UBadge>
          </div>
          <div class="flex flex-row justify-between flex-auto mt-1">
            <span>
              <UIcon name="mdi:help-circle-outline" class="text-(--ui-text-muted)" />
              Total Issues
            </span>
            <UBadge class="ml-1" color="primary" variant="soft">
              {{ githubStats.all_issues.toLocaleString() }}
            </UBadge>
          </div>
        </div>
      </div>
    </template>
  </UPopover>
</template>
