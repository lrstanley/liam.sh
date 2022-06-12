<template>
  <div id="main" v-infinite-scroll="[fetchEvents, { distance: '40' }]">
    <TransitionGroup appear>
      <div
        v-for="(e, i) in fetched"
        :key="e.id"
        :style="{ '--i': fetched.length - i, '--total': fetched.length }"
        class="flex flex-auto flex-row gap-x-1 px-1 hover:bg-dark-100/20 text-gray-400 transition duration-50 ease-out border-b-1px border-b-dark-300/50"
      >
        <a :href="e.actor.login" target="_blank">
          <n-avatar square :size="15" :src="e.actor.avatarURL" class="align-middle mr-1" />
        </a>

        <component :is="eventMap[e.eventType]" :event="e" class="truncate flex flex-grow gap-2" />
        <div class="flex-none">
          <EventHoverItem placement="right">
            <template #value>
              <i-mdi-clock-time-two-outline
                :style="{ '--i': fetched.length - i, '--total': fetched.length }"
                class="timestamp"
              />
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
    count: 25,
  },
  pause: true,
  requestPolicy: "cache-first",
})

function fetchEvents() {
  if (!hasNextPage.value) return

  events.executeQuery().then((result) => {
    const data = result.data.value.githubevents

    if (data.pageInfo.hasNextPage) {
      cursor.value = data.pageInfo.endCursor
    } else {
      hasNextPage.value = false
    }

    fetched.value = [...fetched.value, ...data.edges.map(({ node }) => toRaw(node))]
  })
}

onMounted(() => {
  fetchEvents()
})
</script>

<style scoped>
#main {
  font-size: 1.15em;
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