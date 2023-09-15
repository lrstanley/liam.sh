<script setup lang="ts">
import { version as vueVersion } from "vue"
import { branchMenuOptions } from "@/lib/core/navigation"

const state = useState()

const props = defineProps<{
  prefix?: string
  command?: string
}>()
</script>

<template>
  <n-card
    content-style="padding: 0;display: flex;flex-direction: column;"
    class="flex flex-auto w-full h-full p-0"
  >
    <slot />

    <n-layout-footer bordered class="bottom-bar">
      <span v-motion-fade class="flex flex-auto">
        <n-tooltip trigger="hover">
          <template #trigger>
            <span class="bar-item misc">
              <n-icon class="align-middle">
                <i-logos-visual-studio-code />
              </n-icon>
            </span>
          </template>
          build date: {{ state.base.version.date }}
        </n-tooltip>

        <n-popover
          placement="top-start"
          raw
          :show-arrow="false"
          class="rounded border border-solid border-zinc-700 shadow-none !m-0"
        >
          <template #trigger>
            <span class="bar-item">
              <n-icon class="align-middle">
                <i-mdi-source-branch />
              </n-icon>
              {{ branchMenuOptions.find((o) => o.to === $route.name).name }}
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
                  <i-mdi-source-branch />
                  <span>{{ link.name }}</span>
                </a>
              </router-link>
            </li>
          </ul>
        </n-popover>

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
