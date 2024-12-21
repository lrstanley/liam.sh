<script setup lang="ts">
import { getCodingStats } from "@/utils/http/services.gen"
import type { LanguageStat } from "@/utils/http/types.gen"

const { data: codingStats, suspense } = useQuery({
  queryKey: ["stats", "coding"],
  queryFn: () => unwrapErrors(getCodingStats()),
})
onServerPrefetch(async () => {
  await suspense()
})

type LanguageBucket = LanguageStat & {
  percentage?: number
  title_length?: number
}

const computedCodingStats = computed(() => {
  if (!codingStats.value) return []

  const out: LanguageBucket[] = []
  let maxTitleLength = 5

  // Bucket by language with a cap.
  for (const stat of codingStats.value.languages) {
    if (out.length === 6) {
      out[5].key = "Other"
      out[5].hex_color = stat.hex_color
      out[5].total += stat.total
      continue
    }

    maxTitleLength = Math.max(maxTitleLength, stat.key.length)
    out.push(toRaw(stat))
  }

  // Calculate percentages.
  for (const stat of out) {
    stat.title_length = maxTitleLength
    stat.percentage = Math.round((stat.total / codingStats.value.total_seconds) * 100)
  }

  return out
})
</script>

<template>
  <n-popover
    :width="250"
    raw
    :show-arrow="false"
    v-bind="$attrs"
    class="px-2 py-1 rounded border border-solid border-zinc-700 shadow-none !m-0"
  >
    <template #trigger>
      <span class="bar-item misc">
        <n-icon class="mr-1 align-middle text-violet-400">
          <Icon name="mdi:history" />
        </n-icon>
        {{ codingStats?.total_duration }}
      </span>
    </template>

    <p class="text-center text-violet-400">last {{ codingStats?.calculated_days }} day coding stats</p>

    <div
      v-for="stat in computedCodingStats"
      :key="stat.key"
      class="flex flex-row items-center flex-auto"
    >
      <div class="text-right shrink-0 mr-[1ch]" :style="{ width: stat.title_length + 'ch' }">
        {{ stat.key }}
      </div>

      <div class="w-full rounded o bg-zinc-900">
        <div
          class="h-2 rounded"
          :style="{ width: stat.percentage + '%', 'background-color': stat.hex_color }"
        />
      </div>
      <div class="shrink-0 ml-[1ch] w-[3ch]">{{ stat.percentage }}%</div>
    </div>
  </n-popover>
</template>
