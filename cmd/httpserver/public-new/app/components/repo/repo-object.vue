<script setup lang="ts">
const { value: repo } = defineProps<{
  value: GithubRepositoryRead
}>()

const drawerActive = ref(false)
const gh = useGithubUser()
</script>

<template>
  <div class="flex flex-row gap-4 py-4">
    <UAvatar :src="repo.owner.avatar_url + '&s=40'" class="mt-1" />
    <div class="flex flex-col w-full gap-1">
      <div class="flex flex-row">
        <a
          class="text-[1.4em] md:text-[1.5em] truncate text-gradient bg-gradient-to-br from-pink-500 via-red-500 to-yellow-500"
          :href="repo.html_url"
          target="_blank"
        >
          {{ repo.owner.login == gh?.login ? repo.name : repo.full_name }}
        </a>

        <a v-if="repo.homepage" :href="repo.homepage" target="_blank" class="ml-auto">
          <UButton color="primary" variant="outline" size="xs" class="cursor-pointer">
            <template #icon>
              <UIcon name="mdi:link" />
            </template>
            homepage
          </UButton>
        </a>
      </div>

      <div class="flex flex-row gap-2 text-sm">
        <UTooltip
          v-if="repo.pushed_at"
          :delay-duration="50"
          :content="{ side: 'top' }"
          text="last updated"
        >
          <div class="inline-flex flex-row items-center">
            <UIcon name="mdi:update" class="mr-1 text-(--ui-color-secondary-400)" />
            <span class="italic text-(--ui-text-muted)">
              {{ useTimeAgo(repo.pushed_at).value }}
            </span>
          </div>
        </UTooltip>

        <UTooltip :delay-duration="50" :content="{ side: 'top' }" text="created">
          <div class="flex-row items-center hidden mr-auto md:inline-flex">
            <UIcon name="mdi:rocket-launch-outline" class="mr-1 text-(--ui-color-success-400) ml-1ch" />
            <span class="italic text-(--ui-text-muted)">
              {{ useTimeAgo(repo.created_at).value }}
            </span>
          </div>
        </UTooltip>

        <RepoStatus :value="repo" class="inline-flex gap-1 ml-auto" />
      </div>

      <div v-html="repo.description || 'No description available'" class="my-2" />

      <div class="flex flex-row gap-2">
        <div v-if="repo.edges.labels" class="flex flex-row flex-wrap gap-1">
          <LabelObject
            v-for="label in repo.edges.labels"
            :key="label.id"
            :value="label"
            route="/repos"
            linkable
            class="hidden md:flex"
          />

          <UDrawer v-model:open="drawerActive" direction="bottom">
            <UButton
              label="Repo labels"
              color="neutral"
              variant="subtle"
              size="xs"
              trailing-icon="lucide:chevron-up"
              class="flex md:hidden"
            />

            <template #content>
              <div class="flex flex-wrap flex-auto gap-1 p-4">
                <LabelObject
                  v-for="label in repo.edges.labels"
                  :key="label.id"
                  :value="label"
                  route="/repos"
                  linkable
                  @click="drawerActive = false"
                />
              </div>
            </template>
          </UDrawer>
        </div>

        <!-- views -->
        <div class="flex flex-col ml-auto shrink-0">
          <StatsButton>
            {{ repo.star_count.toLocaleString() }}
            {{ repo.star_count === 1 ? "star" : "stars" }}
          </StatsButton>
        </div>
      </div>
    </div>
  </div>
</template>
