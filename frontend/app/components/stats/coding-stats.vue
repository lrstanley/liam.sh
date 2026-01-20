<script setup lang="ts">
const { data: codingStats, error } = await useApi('/stats/coding')
if (error.value) throw error.value

const maxTitleLength = computed(
  () => codingStats.value?.summary.reduce((max, stat) => Math.max(max, stat.title_length), 0) ?? 0
)
</script>

<template>
  <UPopover v-if="codingStats" mode="hover" :open-delay="150" :content="{ side: 'top', align: 'start', sideOffset: 0 }"
    v-bind="$attrs">
    <div class="items-center bar-item misc">
      <UIcon name="heroicons:clock" class="text-secondary-400" />
      <span>{{ codingStats.total_duration }}</span>
    </div>

    <template #content>
      <div class="px-2 py-1">
        <p class="text-center text-secondary-400">
          last {{ codingStats.calculated_days }} day coding stats
        </p>

        <div v-for="stat in codingStats.summary" :key="stat.key" class="flex flex-row items-center flex-auto">
          <div class="text-right shrink-0 mr-[1ch]" :style="{ width: maxTitleLength + 'ch' }">
            {{ stat.key }}
          </div>

          <div class="w-full rounded-sm bg-default">
            <div class="h-2 rounded-sm" :style="{ width: stat.percentage + '%', 'background-color': stat.hex_color }" />
          </div>
          <div class="shrink-0 ml-[1ch] w-[3ch]">{{ stat.percentage }}%</div>
        </div>
      </div>
    </template>
  </UPopover>
</template>
