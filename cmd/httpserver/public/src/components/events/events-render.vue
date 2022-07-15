<template>
  <div
    id="main"
    ref="scrollContainer"
    v-infinite-scroll="[fetchEvents, { distance: 40 }]"
    class="relative overflow-x-hidden"
  >
    <TransitionGroup appear>
      <div
        v-for="(e, i) in fetched"
        :key="e.id"
        :style="{ '--i': fetched.length - i, '--total': fetched.length }"
        class="flex flex-auto flex-row items-center gap-x-1 px-1 hover:bg-zinc-500/10 text-zinc-400 transition duration-75 ease-out border-b-[1px] border-b-gray-100"
      >
        <a :href="e.actor.login" target="_blank">
          <n-avatar square :size="15" :src="e.actor.avatarURL" class="align-middle mr-1" />
        </a>

        <component
          :is="eventMap[e.eventType]"
          :event="e"
          class="truncate flex items-center grow gap-2"
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

<script setup>
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

const fetched = ref([])
const hasNextPage = ref(true)
const cursor = ref(null)

const events = useGetEventsQuery({
  variables: {
    cursor,
    count: 20,
  },
  pause: true,
  requestPolicy: "cache-first",
})

const scrollContainer = ref(null)

function fetchEvents(wasScrollEvent) {
  if (!hasNextPage.value) return

  events.executeQuery().then((result) => {
    const data = result.data.value.githubevents

    if (data.pageInfo.hasNextPage) {
      cursor.value = data.pageInfo.endCursor
    } else {
      hasNextPage.value = false
    }

    fetched.value = [...fetched.value, ...data.edges.map(({ node }) => toRaw(node))]

    // if this fetch was triggered by a scroll event, don't trigger it a second time.
    if (wasScrollEvent) return
    if (scrollContainer.value.scrollHeight <= scrollContainer.value.clientHeight) {
      fetchEvents(true)
    }
  })
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

.v-enter-active,
.v-leave-active {
  @apply transform transition-opacity duration-300;
  transition-delay: min(0.05s, calc(0.005s * (var(--total) - var(--i))));
}

.v-enter-from,
.v-leave-to {
  @apply opacity-0;
}

div :deep(.icon) {
  font-size: 0.8em;
}
</style>
