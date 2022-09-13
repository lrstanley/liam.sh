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
        :prefix="props.prefix ?? '#'"
        :value="props.command ?? 'Hello!'"
      />
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
                  {{ branchMenuOptions.find((o) => o.to.name === $route.name).name }}
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
                      <span> {{ link.name }} </span>
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
import { branchMenuOptions } from "@/lib/core/navigation"

const props = defineProps<{
  prefix?: string
  command?: string
}>()

const state = useState()
</script>

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
