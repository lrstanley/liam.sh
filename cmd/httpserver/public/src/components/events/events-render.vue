<template>
  <div id="main" ref="scrollContainer" v-infinite-scroll="[fetchEvents, { distance: 40 }]">
    <TransitionGroup name="stepped" appear>
      <div
        v-for="(e, i) in fetched"
        :key="e.id"
        :style="{ '--i': fetched.length - i, '--total': fetched.length }"
        class="flex flex-auto flex-row items-center gap-x-1 px-1 hover:bg-zinc-500/10 text-zinc-400 transition duration-75 ease-out border-b-[1px] border-b-gray-100"
      >
        <a :href="e.actor.login" target="_blank">
          <n-avatar round :size="15" :src="e.actor.avatarURL + '&s=40'" class="mr-1 align-middle" />
        </a>

        <component
          :is="eventMap[e.eventType]"
          :event="e"
          class="flex items-center gap-2 truncate grow"
        />
        <div class="flex-none">
          <EventHoverItem placement="left">
            <template #value>
              <i-mdi-clock-time-two-outline class="timestamp" />
            </template>

            {{ useTimeAgo(e.createdAt).value }}
          </EventHoverItem>
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup lang="ts">
import { useGetEventsQuery } from "@/lib/api"
import { useTimeAgo } from "@vueuse/core"
import { vInfiniteScroll } from "@vueuse/components"

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

import type { GithubEvent } from "@/lib/api"

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
  (e: "eventCount", value: number): void
}>()

const fetched = ref<GithubEvent[]>([])
const hasNextPage = ref(true)
const cursor = ref<string>(null)
const nextCursor = ref<string>(null)

const events = useGetEventsQuery({
  variables: {
    cursor,
    count: 20,
  },
  requestPolicy: "cache-first",
  pause: false,
})

const scrollContainer = ref(null)

watch(
  events.data,
  (data) => {
    if (!data) return

    hasNextPage.value = data.githubEvents.pageInfo.hasNextPage
    nextCursor.value = data.githubEvents.pageInfo.endCursor
    fetched.value = [...fetched.value, ...data.githubEvents.edges.map(({ node }) => node as GithubEvent)]
    emit("eventCount", fetched.value.length)

    setTimeout(() => {
      if (
        hasNextPage.value &&
        scrollContainer.value.scrollHeight <=
          scrollContainer.value.scrollTop + scrollContainer.value.clientHeight
      ) {
        fetchEvents()
      }
    }, 300)
  },
  { immediate: true }
)

function fetchEvents() {
  if (!hasNextPage.value) return
  cursor.value = nextCursor.value
}

onMounted(() => {
  fetchEvents()
})
</script>

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
