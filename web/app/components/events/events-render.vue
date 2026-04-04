<script setup lang="ts">
import { vInfiniteScroll } from "@vueuse/components"
import type { SchemaGithubEvent } from '#open-fetch-schemas/api'

import {
  EventCreate,
  EventDelete,
  EventFork,
  EventIssueComment,
  EventIssues,
  EventMember,
  EventPublic,
  EventPullRequestReviewComment,
  EventPullRequestReviewThread,
  EventPullRequestReview,
  EventPullRequest,
  EventPush,
  EventRelease,
  EventWatch,
} from "#components"

const eventMap: Record<string, any> = {
  CreateEvent: EventCreate,
  DeleteEvent: EventDelete,
  ForkEvent: EventFork,
  IssueCommentEvent: EventIssueComment,
  IssuesEvent: EventIssues,
  MemberEvent: EventMember,
  PublicEvent: EventPublic,
  PullRequestReviewCommentEvent: EventPullRequestReviewComment,
  PullRequestReviewThreadEvent: EventPullRequestReviewThread,
  PullRequestReviewEvent: EventPullRequestReview,
  PullRequestEvent: EventPullRequest,
  PushEvent: EventPush,
  ReleaseEvent: EventRelease,
  WatchEvent: EventWatch,
}

const page = ref(1)

const {
  data: githubEvents,
  error,
  status,
} = await useApi('/github-events', {
  query: {
    page: page,
    per_page: 50,
    "public.eq": true,
    sort: "created_at",
    order: "desc",
  }
})

const events = ref<SchemaGithubEvent[]>([])

watch(
  githubEvents,
  (data) => {
    if (status.value != "success" || data?.content == null) return
    data.content.filter((e) => !eventMap[e.event_type]).forEach((e) => console.warn(`No component for event type: ${e.event_type}`))
    events.value.push(...data.content.filter((e) => !!eventMap[e.event_type]))
  },
  { immediate: true }
)

const scrollContainer = ref(null)

function fetchMoreEvents() {
  if (import.meta.server) return

  if (
    status.value != "success" ||
    githubEvents.value?.is_last_page ||
    githubEvents.value?.page != page.value
  )
    return
  document.getElementById("status")?.scrollIntoView()
  page.value += 1
}
</script>

<template>
  <div id="main" ref="scrollContainer" v-infinite-scroll="[fetchMoreEvents, { distance: 100 }]">
    <motion as="div" :initial="{ opacity: 0 }" :animate="{ opacity: 1 }" :exit="{ opacity: 0 }"
      :transition="{ delay: ((i % 50) + 1) * 0.015 }" v-for="(e, i) in events" :key="e.id"
      class="flex flex-row items-center flex-auto px-1 text-sm gap-x-1 hover:bg-zinc-500/10 text-muted border-b-DEFAULT border-b-gray-100">
      <NuxtLink :href="'https://github.com/' + e.actor.login" target="_blank">
        <UAvatar size="3xs" :src="e.actor.avatar_url + '&s=40'" class="mr-1 align-middle" :alt="e.actor.login" />
      </NuxtLink>

      <component :is="eventMap[e.event_type]" :event="e" class="flex items-center gap-2 truncate grow" />
      <div class="flex-none">
        <EventHoverItem placement="left">
          <template #value>
            <UIcon name="mdi:clock-time-two-outline"
              class="opacity-[max(0.2,calc(var(--i)/var(--total)))] hover:opacity-100 transition duration-100;" />
          </template>

          {{ useTimeAgo(e.created_at).value }}
        </EventHoverItem>
      </div>
    </motion>
    <div v-show="status != 'success' || error" id="status" key="status"
      class="flex flex-row items-center flex-auto px-1 text-sm transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-muted border-b-DEFAULT border-b-gray-100">
      <UIcon name="heroicons:arrow-path-16-solid" class="mr-1 align-middle animate-spin" />
      <div class="flex items-center gap-2 truncate grow">
        {{ error ? "error loading events: " + error.message : "loading..." }}
      </div>
    </div>
  </div>
</template>

<style scoped>
#main {
  font-size: 1.05em;

  @md {
    #main {
      font-size: 1.15em;
    }
  }
}

div :deep(.icon) {
  font-size: 0.8em;
}
</style>
