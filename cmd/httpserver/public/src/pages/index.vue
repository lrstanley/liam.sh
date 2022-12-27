<route lang="yaml">
meta:
  title: Home
  layout: terminal
</route>

<template>
  <ContainerIde>
    <EventsRender
      class="relative w-full h-full overflow-x-hidden grow basis-0"
      @event-count="(e) => (eventCount = e)"
    />

    <template #footer>
      <n-popover
        placement="top-start"
        :width="250"
        raw
        :show-arrow="false"
        class="px-2 py-1 rounded border border-solid border-zinc-700 shadow-none !m-0"
      >
        <template #trigger>
          <span class="bar-item misc">
            <n-icon class="mr-1 align-middle text-violet-400">
              <i-mdi-history />
            </n-icon>
            {{ state.base.codingStats.totalDuration }}
          </span>
        </template>

        <p class="text-center text-violet-400">
          last {{ state.base.codingStats.calculatedDays }} day coding stats
        </p>

        <div
          v-for="stat in codingStats"
          :key="stat.language"
          class="flex flex-row items-center flex-auto"
        >
          <div class="text-right shrink-0 mr-[1ch]" :style="{ width: stat.titleLength + 'ch' }">
            {{ stat.language }}
          </div>

          <div class="w-full rounded o bg-zinc-900">
            <div
              class="h-2 rounded"
              :style="{ width: stat.percentage + '%', 'background-color': stat.hexColor }"
            />
          </div>
          <div class="shrink-0 ml-[1ch] w-[3ch]">{{ stat.percentage }}%</div>
        </div>
      </n-popover>

      <span class="ml-auto" />

      <span v-if="eventCount > 0" class="bar-item misc"> ln:{{ eventCount }} </span>
    </template>
  </ContainerIde>
</template>

<script setup lang="ts">
const eventCount = ref<number>()
const state = useState()

interface LanguageBucket {
  language: string
  hexColor: string
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
      out[5].hexColor = stat.hexColor
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
