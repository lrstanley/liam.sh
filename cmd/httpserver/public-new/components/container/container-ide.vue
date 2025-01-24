<script setup lang="ts">
import { version as vueVersion } from "vue"
import { getServiceVersion } from "@/utils/http/sdk.gen"

const { data: version } = await getServiceVersion({ composable: "useFetch" })
</script>

<template>
  <div
    class="flex flex-col flex-auto max-h-full border rounded size-full max-w-fulkl bg-zinc-900 border-zinc-700/50"
  >
    <slot />

    <div class="box-border flex border-t bottom-bar border-zinc-700/50" v-auto-animate>
      <!-- <n-tooltip v-if="version.data.value" trigger="hover">
        <template #trigger>
          <span class="items-center bar-item misc">
            <UIcon name="logos:visual-studio-code" class="mr-1 align-middle" />
          </span>
        </template>
        build date: {{ version.data.value.build_date }}
      </n-tooltip> -->

      <!-- <n-popover
        placement="top-start"
        raw
        :show-arrow="false"
        class="rounded border border-solid border-zinc-700 shadow-none !m-0"
      >
        <template #trigger>
          <span class="items-center bar-item">
            <UIcon name="mdi:source-branch" />
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
                <UIcon name="mdi:source-branch" />
                <span>{{ link.name }}</span>
              </a>
            </router-link>
          </li>
        </ul>
      </n-popover> -->

      <div v-if="version" class="items-center bar-item misc" data-allow-mismatch>
        <UIcon name="logos:gopher" />
        <span>{{ version.go_version.replace(/^go/, "") }}</span>
      </div>
      <div class="items-center bar-item misc">
        <UIcon name="logos:vue" />
        {{ vueVersion }}
      </div>

      <slot name="footer">
        <div class="ml-auto" />
      </slot>

      <UTooltip
        text="... or just gofmt"
        :popper="{ placement: 'top' }"
        :ui="{ background: 'dark:bg-zinc-800' }"
      >
        <span class="bar-item misc">spaces:4</span>
      </UTooltip>

      <GithubStats placement="top-end" />
    </div>
  </div>
</template>

<style>
.bottom-bar {
  line-height: 1.4;
  @apply text-[1.15rem] md:text-[0.95rem];
}

.bar-item {
  @apply px-2 rounded-br-sm inline-flex text-zinc-400 align-middle cursor-pointer transition hover:bg-zinc-800;
}

.bar-item .iconify {
  @apply align-middle text-sm mr-1;
}

.bar-item.misc {
  @apply hidden sm:inline-flex;
}
</style>
