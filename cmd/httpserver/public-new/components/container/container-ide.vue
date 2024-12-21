<script setup lang="ts">
import { version as vueVersion } from "vue"
import { getServiceVersion } from "@/utils/http/services.gen"

const { data: version, suspense } = useQuery({
  queryKey: ["version"],
  queryFn: () => unwrapErrors(getServiceVersion()),
})
await suspense()
</script>

<template>
  <n-card
    content-style="padding: 0;display: flex;flex-direction: column;"
    class="flex flex-auto p-0 size-full"
  >
    <slot />

    <n-layout-footer bordered class="bottom-bar" v-auto-animate>
      <span class="flex flex-auto">
        <n-tooltip trigger="hover">
          <template #trigger>
            <span class="items-center bar-item misc">
              <Icon name="logos:visual-studio-code" />
            </span>
          </template>
          build date: {{ version?.build_date }}
        </n-tooltip>

        <n-popover
          placement="top-start"
          raw
          :show-arrow="false"
          class="rounded border border-solid border-zinc-700 shadow-none !m-0"
        >
          <template #trigger>
            <span class="items-center bar-item">
              <Icon name="mdi:source-branch" />
              {{ branchMenuOptions.find((o) => o.to === $route.name)?.name }}
            </span>
          </template>

          <ul class="flex flex-col flex-nowrap">
            <li v-for="link in branchMenuOptions" :key="link.name">
              <router-link
                v-slot="{ isActive, href, navigate }"
                :to="link.to"
                active-class="bg:zinc-900/50"
                custom
              >
                <a
                  :href="href"
                  class="flex items-center px-1 rounded hover:bg-zinc-900/50 hover:text-zinc-300 text-zinc-400"
                  :class="{ 'bg-zinc-900/50 text-zinc-300': isActive }"
                  @click="navigate"
                >
                  <Icon name="mdi:source-branch" />
                  <span>{{ link.name }}</span>
                </a>
              </router-link>
            </li>
          </ul>
        </n-popover>

        <span class="items-center bar-item misc">
          <Icon name="logos:gopher" class="mr-1" />
          {{ version?.go_version.replace(/^go/, "") }}
        </span>
        <span class="items-center bar-item misc">
          <Icon name="logos:vue" class="mr-1" />
          {{ vueVersion }}
        </span>

        <slot name="footer">
          <span class="ml-auto" />
        </slot>

        <n-tooltip trigger="hover">
          <template #trigger>
            <span class="bar-item misc">spaces:4</span>
          </template>
          ... or just gofmt
        </n-tooltip>

        <GithubStats placement="top-end" />
      </span>
    </n-layout-footer>
  </n-card>
</template>

<style>
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
