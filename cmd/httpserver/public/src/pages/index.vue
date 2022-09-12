<template>
  <!-- TODO: use less flex -->
  <div class="flex items-stretch justify-center flex-auto w-full h-full lg:items-center">
    <div
      class="flex h-full w-full shrink grow-0 basis-auto flex-col items-stretch pt-[15px] md:items-center lg:max-h-[28rem] lg:max-w-3xl"
    >
      <CoreNavigation />
      <CoreTerminal
        class="mb-4 flex flex-auto justify-center text-[38px] md:text-[45px]"
        path="~"
        prefix="#"
        value="Hello!"
      />
      <n-card
        content-style="padding: 0;display: flex;flex-direction: column;"
        class="flex flex-auto w-full h-full p-0"
      >
        <EventsRender
          class="w-full h-full overflow-x-auto shrink grow basis-0"
          @event-count="(e) => (eventCount = e)"
        />

        <n-layout-footer bordered class="bottom-bar">
          <span v-motion-fade class="flex flex-auto">
            <span class="bar-item misc">
              <n-icon class="align-middle">
                <i-logos-visual-studio-code />
              </n-icon>
            </span>

            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="bar-item">
                  <n-icon class="align-middle">
                    <i-mdi-source-branch />
                  </n-icon>
                  {{
                    state.base.version?.commit == "unknown"
                      ? "master"
                      : state.base.version?.commit?.substring(0, 8)
                  }}
                </span>
              </template>
              build date: {{ state.base.version.date }}
            </n-tooltip>
            <span class="bar-item misc">
              <n-icon class="mr-1 align-middle">
                <i-logos-gopher />
              </n-icon>
              {{ state.base.version.goVersion.replace(/^go/, "") }}
            </span>
            <span class="bar-item misc">
              <n-icon class="mr-1 align-middle">
                <i-logos-vue />
              </n-icon>
              {{ vueVersion }}
            </span>

            <n-popover placement="top" :width="250">
              <template #trigger>
                <span class="bar-item misc">
                  <n-icon class="mr-1 align-middle text-violet-400">
                    <i-mdi-history />
                  </n-icon>
                  {{ state.base.codingStats.totalDuration }}
                </span>
              </template>

              <p class="text-violet-400">{{ state.base.codingStats.calculatedDays }} day coding stats</p>

              <div
                v-for="stat in codingStats"
                :key="stat.language"
                class="flex flex-row items-center flex-auto"
              >
                <div class="text-right shrink-0 mr-[1ch]" :style="{ width: stat.titleLength + 'ch' }">
                  {{ stat.language }}
                </div>

                <div class="w-full rounded bg-zinc-900">
                  <div
                    class="h-2 rounded bg-gradient-to-r from-fuchsia-600 to-pink-600"
                    :style="{ width: stat.percentage + '%' }"
                  />
                </div>
                <div class="shrink-0 ml-[1ch] w-[3ch]">{{ stat.percentage }}%</div>
              </div>
            </n-popover>

            <span class="ml-auto" />

            <span v-if="eventCount > 0" class="bar-item misc"> ln:{{ eventCount }} </span>
            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="bar-item misc"> spaces:4 </span>
              </template>
              ... or just gofmt
            </n-tooltip>
            <n-tooltip trigger="hover">
              <template #trigger>
                <a
                  class="px-2 transition bg-blue-600 rounded-br-sm hover:bg-blue-800 hover:text-current"
                  style="color: white !important"
                  :href="state.base.githubUser.htmlurl"
                >
                  <n-icon class="align-middle mt-[-3px] mr-[-7px]">
                    <i-mdi-github />
                  </n-icon>
                  {{ state.base.githubUser.login }}
                </a>
              </template>
              <p>
                <n-icon class="align-middle"><i-mdi-github /></n-icon>
                {{ state.base.githubUser.name }} &middot; {{ state.base.githubUser.bio }}
              </p>
            </n-tooltip>
          </span>
        </n-layout-footer>
      </n-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { version as vueVersion } from "vue"

const eventCount = ref<number>()
const state = useState()

interface LanguageBucket {
  language: string
  totalSeconds: number
  percentage?: number
  titleLength?: number
}

const codingStats = computed(() => {
  const out: LanguageBucket[] = []
  let maxTitleLength = 5

  // Bucket by language with a cap.
  for (const stat of state.base.codingStats.languages) {
    if (out.length === 6) {
      out[5].language = "Other"
      out[5].totalSeconds += stat.totalSeconds
      continue
    }

    maxTitleLength = Math.max(maxTitleLength, stat.language.length)
    out.push(stat)
  }

  // Calculate percentages.
  for (const stat of out) {
    stat.titleLength = maxTitleLength
    stat.percentage = Math.round((stat.totalSeconds / state.base.codingStats.totalSeconds) * 100)
  }

  return out
})
</script>

<style scoped>
.bottom-bar {
  line-height: 1.5;
  @apply text-[1.2em] md:text-[1em];
}

.bar-item {
  @apply pl-1.5 pr-2 rounded-br-sm inline-flex text-zinc-400 align-middle cursor-pointer transition hover:bg-zinc-800;
}

.bar-item.misc {
  @apply hidden sm:inline-flex;
}
</style>
