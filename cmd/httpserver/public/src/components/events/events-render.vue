<script setup lang="ts">
import { useTimeAgo } from "@vueuse/core"
import { vInfiniteScroll } from "@vueuse/components"
import { useInfiniteQuery } from "@tanstack/vue-query"
import { listGithubEvents } from "@/lib/http/services.gen"

import EventCreate from "@/components/events/objects/event-create.vue"
import EventDelete from "@/components/events/objects/event-delete.vue"
import EventFork from "@/components/events/objects/event-fork.vue"
import EventIssueComment from "@/components/events/objects/event-issue-comment.vue"
import EventIssues from "@/components/events/objects/event-issues.vue"
import EventMember from "@/components/events/objects/event-member.vue"
import EventPublic from "@/components/events/objects/event-public.vue"
import EventPullRequestReviewComment from "@/components/events/objects/event-pull-request-review-comment.vue"
import EventPullRequestReviewThread from "@/components/events/objects/event-pull-request-review-thread.vue"
import EventPullRequestReview from "@/components/events/objects/event-pull-request-review.vue"
import EventPullRequest from "@/components/events/objects/event-pull-request.vue"
import EventPush from "@/components/events/objects/event-push.vue"
import EventRelease from "@/components/events/objects/event-release.vue"
import EventWatch from "@/components/events/objects/event-watch.vue"

const eventMap = {
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

const {
  data: events,
  hasNextPage,
  fetchNextPage,
  isFetchingNextPage,
  error,
} = useInfiniteQuery({
  queryKey: ["events", "infinite"],
  queryFn: ({ pageParam }) =>
    listGithubEvents({
      query: {
        page: pageParam,
        per_page: 20,
        "public.eq": true,
        sort: "created_at",
        order: "desc",
      },
    }),
  initialPageParam: 1,
  getNextPageParam: (lastPage) => (lastPage.data.is_last_page ? undefined : lastPage.data.page + 1),
})

const allEvents = computed(() => {
  const results = events.value?.pages.flatMap((page) => page.data.content) ?? []
  emit("eventCount", results.length)
  return results
})

const scrollContainer = ref(null)

function fetchMoreEvents() {
  if (!hasNextPage.value || isFetchingNextPage.value) return
  document.getElementById("status")?.scrollIntoView({ behavior: "smooth" })
  fetchNextPage()
}
</script>

<template>
  <div id="main" ref="scrollContainer" v-infinite-scroll="[fetchMoreEvents, { distance: 40 }]">
    <TransitionGroup name="stepped" appear>
      <div
        v-for="(e, i) in allEvents"
        :key="e.id"
        :style="{ '--i': allEvents.length - i, '--total': allEvents.length }"
        class="flex flex-row items-center flex-auto px-1 transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-zinc-400 border-b-DEFAULT border-b-gray-100"
      >
        <a :href="e.actor.login as string" target="_blank">
          <n-avatar round :size="15" :src="e.actor.avatar_url + '&s=40'" class="mr-1 align-middle" />
        </a>

        <component
          :is="eventMap[e.event_type]"
          :event="e"
          class="flex items-center gap-2 truncate grow"
        />
        <div class="flex-none">
          <EventHoverItem placement="left">
            <template #value>
              <i-mdi-clock-time-two-outline class="timestamp" />
            </template>

            {{ useTimeAgo(e.created_at).value }}
          </EventHoverItem>
        </div>
      </div>
      <div
        v-show="hasNextPage || error"
        id="status"
        key="status"
        class="flex flex-row items-center flex-auto px-1 transition duration-75 ease-out gap-x-1 hover:bg-zinc-500/10 text-zinc-400 border-b-DEFAULT border-b-gray-100"
      >
        <i-mdi-loading class="mr-1 align-middle animate-spin" />
        <div class="flex items-center gap-2 truncate grow">
          {{ error ? "error loading events: " + error.message : "loading..." }}
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
#main {
  font-size: 1.05em;
}

@screen md {
  #main {
    font-size: 1.15em;
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
