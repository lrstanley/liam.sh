<script setup lang="ts">
import { vInfiniteScroll } from "@vueuse/components"

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

const emit = defineEmits<{
  eventCount: [value: number]
}>()

const page = ref(1)

const {
  data: githubEvents,
  error,
  status,
} = await listGithubEvents({
  composable: "useFetch",
  query: {
    page: page,
    per_page: 50,
    "public.eq": true,
    sort: "created_at",
    order: "desc",
  },
})

const events = ref<GithubEvent[]>([])

watch(
  githubEvents,
  (data) => {
    if (status.value != "success" || data?.content == null) return
    events.value.push(...data.content.filter((e) => !!eventMap[e.event_type]))
    emit("eventCount", events.value.length)
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
  <div
    id="main"
    ref="scrollContainer"
    v-infinite-scroll="[fetchMoreEvents, { distance: 100 }]"
    v-auto-animate
  >
    <div
      v-for="(e, i) in events"
      :key="e.id"
      :style="{ '--i': events.length - i, '--total': events.length }"
      class="flex flex-row items-center flex-auto px-1 text-sm transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-(--ui-text-muted) border-b-DEFAULT border-b-gray-100"
    >
      <a :href="'https://github.com/' + e.actor.login" target="_blank">
        <UAvatar
          size="3xs"
          :src="e.actor.avatar_url + '&s=40'"
          class="mr-1 align-middle"
          :alt="e.actor.login"
        />
      </a>

      <component :is="eventMap[e.event_type]" :event="e" class="flex items-center gap-2 truncate grow" />
      <div class="flex-none">
        <EventHoverItem placement="left">
          <template #value>
            <UIcon name="mdi:clock-time-two-outline" class="timestamp" />
          </template>

          {{ useTimeAgo(e.created_at).value }}
        </EventHoverItem>
      </div>
    </div>
    <div
      v-show="status != 'success' || error"
      id="status"
      key="status"
      class="flex flex-row items-center flex-auto px-1 text-sm transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-(--ui-text-muted) border-b-DEFAULT border-b-gray-100"
    >
      <UIcon name="heroicons:arrow-path-16-solid" class="mr-1 align-middle animate-spin" />
      <div class="flex items-center gap-2 truncate grow">
        {{ error ? "error loading events: " + error.error : "loading..." }}
      </div>
    </div>
  </div>
</template>

<style scoped>
@reference "@/assets/css/main.css";

#main {
  font-size: 1.05em;

  @md {
    #main {
      font-size: 1.15em;
    }
  }
}

.timestamp {
  opacity: max(0.2, calc(var(--i) / var(--total)));
  @apply transition duration-100;
}

.timestamp:hover {
  opacity: 1;
}

div :deep(.icon) {
  font-size: 0.8em;
}
</style>
